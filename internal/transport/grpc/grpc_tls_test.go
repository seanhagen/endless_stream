package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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
