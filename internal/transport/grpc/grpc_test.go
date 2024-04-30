package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
)

func TestTransportGRPC_TLSConfiguration(t *testing.T) {
	ctx := context.TODO()

	tests := []struct {
		name  string
		setup func(*testing.T, context.Context) (TLSConfig, grpc.DialOption)
	}{
		{
			name: "insecure TLS",
			setup: func(t *testing.T, ctx context.Context) (TLSConfig, grpc.DialOption) {
				return WithInsecureTLS(), grpc.WithInsecure()
			},
		},

		{
			name: "normal TLS configuration",
			setup: func(t *testing.T, ctx context.Context) (TLSConfig, grpc.DialOption) {
				t.Helper()

				certPath := "testdata/server.pem"
				keyPath := "testdata/server-key.pem"
				tlsConf := WithTLSConfig(certPath, keyPath)

				pemServerCA, err := os.ReadFile("testdata/root.pem")
				require.NoError(t, err, "uanble to load root.pem")

				certPool := x509.NewCertPool()
				ok := certPool.AppendCertsFromPEM(pemServerCA)
				require.True(t, ok, "unable to append certs from pool")

				tlsClientConfig := &tls.Config{
					RootCAs: certPool,
				}

				opt := grpc.WithTransportCredentials(credentials.NewTLS(tlsClientConfig))

				return tlsConf, opt
			},
		},

		{
			name: "mutual TLS configuration",
			setup: func(t *testing.T, ctx context.Context) (TLSConfig, grpc.DialOption) {
				t.Helper()

				certPath := "testdata/server.pem"
				keyPath := "testdata/server-key.pem"
				rootPath := "testdata/root.pem"
				tlsConf := WithMutualTLSConfig(certPath, keyPath, rootPath)

				pemServerCA, err := os.ReadFile(rootPath)
				require.NoError(t, err, "uanble to load root.pem")

				certPool := x509.NewCertPool()
				ok := certPool.AppendCertsFromPEM(pemServerCA)
				require.True(t, ok, "unable to append certs from pool")

				clientCertPath := "testdata/client.pem"
				clientKeyPath := "testdata/client-key.pem"

				clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
				require.NoError(
					t,
					err,
					"unable to load x509 key pair from %q and %q",
					clientCertPath,
					clientKeyPath,
				)

				tlsClientConfig := &tls.Config{
					Certificates: []tls.Certificate{clientCert},
					RootCAs:      certPool,
				}

				opt := grpc.WithTransportCredentials(credentials.NewTLS(tlsClientConfig))

				return tlsConf, opt
			},
		},
	}

	for i, x := range tests {
		tt := x
		t.Run(
			fmt.Sprintf("test %d %s", i+1, tt.name),
			func(t *testing.T) {
				logger := logs.NewTestLog(
					t,
					&logs.Config{Out: logs.NewTestLogOutput(t, true), Level: logs.LevelDebug},
				)

				tlsConf, clientOpt := tt.setup(t, ctx)

				expectResp := &proto.PongResp{Gsm: "boop"}

				handler := func(ctx context.Context, pr *proto.PingReq) (*proto.PongResp, error) {
					values := metadata.ValueFromIncomingContext(ctx, ":authority")
					require.Len(t, values, 1, "expect only one value for ':authority' in metadata")
					expect := []string{fmt.Sprintf("127.0.0.1:%d", DefaultGRPCPort)}
					assert.Equal(t, expect, values, "wrong authority value from metadata")

					return expectResp, nil
				}

				svc := &testService{pingHandler: handler}

				config := Config{
					Logger:   logger,
					Network:  WithGrpcOnly(DefaultGRPCPort),
					TLS:      tlsConf,
					Timeout:  DefaultTimeout,
					Services: []Service{svc},
				}

				transport, err := BuildTransport(ctx, config)
				require.NoError(t, err, "unable to build transport from config")
				require.NotNil(t, transport, "expected non-nil transport")

				t.Log("starting transport")
				ctxWithCancel, cancelFn := context.WithCancel(ctx)
				t.Cleanup(cancelFn)
				err = transport.Start(ctxWithCancel)
				require.NoError(t, err)
				assert.True(t, transport.Running())
				assert.NotZero(t, svc.registerCalls)

				conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", DefaultGRPCPort), clientOpt)
				require.NoError(t, err, "unable to dial test server")
				t.Cleanup(func() {
					assert.NoError(t, conn.Close())
				})

				client := proto.NewTestClient(conn)

				req := &proto.PingReq{
					Msg: "hello world",
				}
				ctxWithTimeout, cancelTimeoutFn := context.WithTimeout(ctx, time.Second*5)
				t.Log("sending ping request via GRPC client")
				resp, err := client.Ping(ctxWithTimeout, req)
				t.Cleanup(cancelTimeoutFn)
				require.NoError(t, err)
				assert.Equal(t, expectResp.Gsm, resp.Gsm)

				cancelFn()
				time.Sleep(time.Millisecond * 200)
				assert.False(t, transport.Running())
			},
		)
	}
}

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
