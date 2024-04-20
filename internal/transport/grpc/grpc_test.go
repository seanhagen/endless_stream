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
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/test/bufconn"
)

func TestTransportGRPC_Setup(t *testing.T) {
	ctx := context.TODO()

	svc := testService{}

	bufferSize := 101024 * 1024
	lis := bufconn.Listen(bufferSize)

	config := Config{
		// required
		Timeout: time.Minute, // time.Duration

		Services: []Service{
			svc,
		},

		// optional, will get appended to defaults
		Middleware:        []Middleware{},
		StreamInterceptor: []StreamInterceptor{},
		UnaryInterceptor:  []UnaryInterceptor{},

		// optional, will override default
		Listener: lis, // net.Listener interface
		DialOpts: []DialOpt{},

		// optional, no default set
		StatsHandler: stats.Handler{},
		KeepAlive:    KeepAlive{Enforcement: EnforcementPolicy{}, Params: Params{}},
		Max: Maximums{
			ConcurrentStreams: 10, // uint32
			MaxHeaderListSize: 10, // uint32
			MaxRecvMsgSize:    10, // uint32
			MaxSendMsgSize:    10, // uint32
			ReadBufferSize:    10, // uint32, default is 32kb
			WriteBufferSize:   10, // uint32, default is 32kb
		},
		UnknownServiceHandler: StreamHandler,
		TransportCredentials:  TransportCredsentials{},
	}

	transport, err := BuildTransport(ctx, config)

	require.NoError(t, err)
	require.NotNil(t, transport)

	stopFn, err := transport.Start()
	require.NoError(t, err)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	client := buildTestClient(t, ctx, lis, opts...)
	// client.Ping(ctx context.Context, in *proto.PingReq, opts ...grpc.CallOption)
}

func buildTestClient(
	t *testing.T,
	ctx context.Context,
	listener net.Listener,
	opts ...grpc.DialOption,
) proto.TestClient {
	t.Helper()

	conn, err := grpc.DialContext(
		ctx, "",
		grpc.WithContextDialer(
			func(_ context.Context, _ string) (net.Conn, err) {
				return listener.Dial()
			},
		),
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
	listener net.Listener,
) (proto.TestClient, func()) {
	baseServer := grpc.NewServer()
	proto.RegisterTestServer(baseServer, testSvc)

	go func() {
		err := baseServer.Serve(listener)
		require.NoError(t, err)
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(
			func(context.Context, string) (new.Conn, error) {
				return listener.Dial()
			},
		),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	require.NoError(t, err, "unable to dial test server")

	closer := func() {
		err := lis.Close()
		assert.NoError(t, err)
		baseServer.Stop()
	}

	client := proto.NewTestClient(conn)

	return client, closer
}
