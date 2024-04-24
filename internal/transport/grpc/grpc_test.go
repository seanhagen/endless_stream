package grpc

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestTransportGRPC_Setup(t *testing.T) {
	ctx := context.TODO()

	testPing := &testPingHandler{}
	svc := &testService{
		pingHandler: testPing.PingHandler,
	}

	bufferSize := 101024 * 1024
	lis := bufconn.Listen(bufferSize)

	// // configure to only be GRPC on port 8000,
	// netConf := GrpcOnlyNetwork(8000)
	// // OR
	// // configure GRPC with GRPC Gateway on port 8000 ( via cmux )
	// netConf := GrpcGatewayNetwork(8000)
	// // OR
	// // configure GRPC on port 8000, GRPC Gateway HTTP on port 8080
	// netConf := GrpcGatewayWithSeparateHTTPNetwork(8000, 8080)
	// OR
	// configure custom net.Listerner
	netConf := WithCustomListener(lis)

	// // configure "normal" TLS
	// sslConf := TlsConfig("/path/to/certificates")
	// // OR
	// // configure mTLS
	// sslConf := MutualTlsConfig("/path/to/certs")

	config := Config{
		// optional
		Network: netConf,

		// // optional, applies to both GRPC & HTTP if GRPC Gateway is in use
		// SSL: sslConf,

		// required, defines connection timeout
		Timeout: DefaultTimeout, // time.Duration

		Services: []Service{
			// each svc that is going to be handled by the GRPC server
			// transport that's set up.
			svc,
		},

		// // optional, appended to default list of interfceptors
		// StreamInterceptors: []StreamInterceptor{},
		// UnaryInterceptors:  []UnaryInterceptor{},

		// // DialOpts is used as part of setting up grpc-gateway
		// DialOpts: []DialOpt{},

		// // optional, no default set
		// StatsHandler: stats.Handler{},
		// KeepAlive:    KeepAlive{Enforcement: EnforcementPolicy{}, Params: Params{}},
		// Max: Maximums{
		// 	ConcurrentStreams: 10, // uint32
		// 	MaxHeaderListSize: 10, // uint32
		// 	MaxRecvMsgSize:    10, // uint32
		// 	MaxSendMsgSize:    10, // uint32
		// 	ReadBufferSize:    10, // uint32, default is 32kb
		// 	WriteBufferSize:   10, // uint32, default is 32kb
		// },

		// EnableSharedWriteBuffer: true,

		// UnknownServiceHandler: StreamHandler,
	}

	transport, err := BuildTransport(ctx, config)

	require.NoError(t, err)
	require.NotNil(t, transport)

	ctxWithCancel, cancelFn := context.WithCancel(ctx)

	err = transport.Start(ctxWithCancel)
	require.NoError(t, err)

	assert.True(t, transport.Running())

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	client := buildTestClient(t, ctx, lis, opts...)

	req := &proto.PingReq{
		Msg: "hello world",
	}

	ctxWithTimeout, cancelTimeoutFn := context.WithTimeout(ctx, time.Second*5)

	resp, err := client.Ping(ctxWithTimeout, req)
	cancelTimeoutFn()

	assert.NoError(t, err)
	assert.Equal(t, "dlrow olleh", resp.GetGsm())
	assert.Equal(t, 1, svc.pingCalls)
	assert.Contains(t, testPing.msgs, "hello world")

	cancelFn()

	time.Sleep(time.Millisecond * 200)
	assert.False(t, transport.Running())
}

func TestTransportGRPC_Stop(t *testing.T) {
	ctx := context.TODO()

	testPing := &testPingHandler{}
	svc := &testService{
		pingHandler: testPing.PingHandler,
	}

	bufferSize := 101024 * 1024
	lis := bufconn.Listen(bufferSize)

	netConf := WithCustomListener(lis)

	config := Config{
		Network: netConf,
		Timeout: DefaultTimeout, // time.Duration
		Services: []Service{
			svc,
		},
	}

	transport, err := BuildTransport(ctx, config)

	require.NoError(t, err)
	require.NotNil(t, transport)

	ctx, timeoutCancelFn := context.WithTimeout(ctx, time.Second*5)
	t.Cleanup(func() { timeoutCancelFn() })

	ctxWithCancel, cancelFn := context.WithCancel(ctx)
	t.Cleanup(func() { cancelFn() })

	err = transport.Start(ctxWithCancel)
	require.NoError(t, err)
	assert.True(t, transport.Running())

	err = transport.Stop()
	assert.NoError(t, err)
	assert.NoError(t, ctx.Err())

	cancelFn()
	<-ctx.Done()

	assert.False(t, transport.Running())
}

func buildTestClient(
	t *testing.T,
	ctx context.Context,
	listener *bufconn.Listener,
	opts ...grpc.DialOption,
) proto.TestClient {
	t.Helper()

	opts = append(
		[]grpc.DialOption{
			grpc.WithContextDialer(
				func(_ context.Context, _ string) (net.Conn, error) {
					return listener.Dial()
				},
			),
		},
		opts...,
	)

	conn, err := grpc.DialContext(
		ctx, "",
		opts...,
	)
	require.NoError(t, err, "unable to dial listener for test client")

	return proto.NewTestClient(conn)
}

// this function should really be used for testing clients, not server stuff
//
// keeping it around so i don't have to re-write it later
func buildTestServer(
	t *testing.T,
	ctx context.Context,
	testSvc *testService,
	listener *bufconn.Listener,
) (proto.TestClient, func()) {
	baseServer := grpc.NewServer()
	proto.RegisterTestServer(baseServer, testSvc)

	go func() {
		err := baseServer.Serve(listener)
		require.NoError(t, err)
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			},
		),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	require.NoError(t, err, "unable to dial test server")

	closer := func() {
		err := listener.Close()
		assert.NoError(t, err)
		baseServer.Stop()
	}

	client := proto.NewTestClient(conn)

	return client, closer
}
