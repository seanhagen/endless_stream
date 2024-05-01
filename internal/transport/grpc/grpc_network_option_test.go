package grpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/motemen/go-loghttp"
	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func buildPortListener(t *testing.T, ctx context.Context) (NetworkConfig, proto.TestClient) {
	conf := WithGrpcOnly(DefaultGRPCPort)

	uri := fmt.Sprintf("localhost:%d", DefaultGRPCPort)
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	require.NoError(t, err, "unable to dial %q", uri)

	client := proto.NewTestClient(conn)

	return conf, client
}

func buildBufferListener(t *testing.T, ctx context.Context) (NetworkConfig, proto.TestClient) {
	bufferSize := 101024 * 1024
	lis := bufconn.Listen(bufferSize)

	conf := WithCustomListener(lis)
	require.NotNil(t, conf, "expected non-nil NetworkConfig from WithCustomListener")

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}

	client := buildTestClient(t, ctx, lis, opts...)
	return conf, client
}

func TestTransportGRPC_NetworkOptions(t *testing.T) {
	ctx := context.TODO()

	tests := []struct {
		name             string
		enableGateway    bool
		setup            func(*testing.T, context.Context) (NetworkConfig, proto.TestClient)
		buildHTTPRequest func(*testing.T, context.Context) *http.Request
	}{
		{
			name: "grpc-only listener",
			setup: func(t *testing.T, _ context.Context) (NetworkConfig, proto.TestClient) {
				t.Helper()

				conf := WithGrpcOnly(DefaultGRPCPort)
				require.NotNil(t, conf, "expected non-nil NetworkConfig from WithGrpcOnly")

				serviceURI := fmt.Sprintf("localhost:%d", DefaultGRPCPort)
				conn, err := grpc.Dial(serviceURI, grpc.WithInsecure())
				require.NoError(t, err, "unable to dial %s", serviceURI)

				client := proto.NewTestClient(conn)

				return conf, client
			},
		},

		{
			name:          "grpc with gateway on same port as http",
			enableGateway: true,
			setup: func(t *testing.T, ctx context.Context) (NetworkConfig, proto.TestClient) {
				t.Helper()

				conf := WithSharedGrpcGatewayPort(DefaultGRPCPort)
				require.NotNil(
					t,
					conf,
					"expected non-nil NetworkConfig from WithSharedGrpcGatewayPort",
				)

				serviceURI := fmt.Sprintf("127.0.0.1:%d", DefaultGRPCPort)
				conn, err := grpc.DialContext(
					ctx,
					serviceURI,
					grpc.WithInsecure(),
				)
				require.NoError(t, err, "unable to dial %s", serviceURI)

				client := proto.NewTestClient(conn)

				return conf, client
			},
			buildHTTPRequest: func(t *testing.T, ctx context.Context) *http.Request {
				t.Helper()

				buf := bytes.NewBuffer(nil)
				req := proto.PingReq{
					Msg: "hello world",
				}

				err := json.NewEncoder(buf).Encode(req)
				require.NoError(t, err, "unable to JSON encode proto.PingReq into buffer")

				hr, err := http.NewRequestWithContext(
					ctx,
					http.MethodPost,
					fmt.Sprintf("http://127.0.0.1:%d/v1/ping", DefaultGRPCPort),
					buf,
				)
				require.NoError(t, err, "unable to build HTTP request")
				// spew.Dump(hr)

				return hr
			},
		},

		{
			name:          "grpc gateway on separate port from grpc",
			enableGateway: true,
			setup: func(t *testing.T, ctx context.Context) (NetworkConfig, proto.TestClient) {
				t.Helper()

				conf := WithSeparateGrpcGatewayPort(8888, 9999)
				require.NotNil(
					t,
					conf,
					"expected non-nil NetworkConfig from WithSeparateGrpcGatewayPort",
				)

				serviceURI := "127.0.0.1:8888"
				conn, err := grpc.DialContext(ctx, serviceURI, grpc.WithInsecure())
				require.NoError(t, err, "unable to dial %s", serviceURI)

				client := proto.NewTestClient(conn)

				return conf, client
			},
			buildHTTPRequest: func(t *testing.T, ctx context.Context) *http.Request {
				t.Helper()

				buf := bytes.NewBuffer(nil)
				req := proto.PingReq{Msg: "hello world"}

				err := json.NewEncoder(buf).Encode(req)
				require.NoError(t, err, "uanble to JSON encode proto.PingReq as JSON into buffer")

				hr, err := http.NewRequestWithContext(
					ctx,
					http.MethodPost,
					"http://127.0.0.1:9999/v1/ping",
					buf,
				)
				require.NoError(t, err, "unable to build HTTP request")

				return hr
			},
		},

		{
			name:  "custom listener works",
			setup: buildBufferListener,
		},
	}

	for i, x := range tests {
		tt := x

		t.Run(
			fmt.Sprintf(
				"test %d %s",
				i+1, tt.name,
			),
			func(t *testing.T) {
				logger := logs.NewTestLog(
					t, &logs.Config{Out: io.Discard},
				)

				// configure the NetworkConfig & proto.TestClient
				netConf, client := tt.setup(t, ctx)

				testPing := &testPingHandler{}
				svc := &testService{
					pingHandler: testPing.PingHandler,
				}

				// create the transport
				t.Log("creating transport")
				config := Config{
					Logger:   logger,
					Network:  netConf,
					Timeout:  DefaultTimeout,
					Services: []Service{svc},
				}
				transport, err := BuildTransport(ctx, config)
				require.NoError(t, err, "unable to build transport from config")
				require.NotNil(t, transport, "expected non-nil transport")

				// start the transport
				t.Log("starting transport")
				ctxWithCancel, cancelFn := context.WithCancel(ctx)
				t.Cleanup(cancelFn)
				err = transport.Start(ctxWithCancel)
				require.NoError(t, err)
				assert.True(t, transport.Running())
				assert.NotZero(t, svc.registerCalls)

				// make a request
				t.Log("creating request")
				req := &proto.PingReq{
					Msg: "hello world",
				}
				ctxWithTimeout, cancelTimeoutFn := context.WithTimeout(ctx, time.Second*5)
				t.Log("sending ping request via GRPC client")
				resp, err := client.Ping(ctxWithTimeout, req)
				t.Cleanup(cancelTimeoutFn)

				// do some asserting
				assert.NoError(t, err)
				assert.Equal(t, "dlrow olleh", resp.GetGsm())
				assert.Equal(t, 1, svc.pingCalls)
				assert.Contains(t, testPing.msgs, "hello world")

				if tt.enableGateway {
					t.Log("testing GRPC gateway")
					require.NotZero(t, svc.gatewayRegisterCalls)

					// make an HTTP request
					t.Log("building HTTP request")
					ctxWithTimeout, cancelHTTPTimeoutFn := context.WithTimeout(ctx, time.Second*5)
					req := tt.buildHTTPRequest(t, ctxWithTimeout)
					t.Cleanup(cancelHTTPTimeoutFn)

					t.Log("sending HTTP request")
					client := &http.Client{Transport: &loghttp.Transport{}}
					httpResp, err := client.Do(req)
					// spew.Dump(httpResp.Body, err)

					require.NoError(
						t,
						err,
						"expected no error when making http request to grpc-gateway",
					)

					err = json.NewDecoder(httpResp.Body).Decode(resp)
					require.NoError(
						t,
						err,
						"expected no error when unmarshaling JSON response body to proto.PongResp",
					)

					assert.Equal(t, "dlrow olleh", resp.GetGsm())
				}

				// stop the server, wait a bit for it to finish
				cancelFn()
				time.Sleep(time.Millisecond * 200)

				// then check the server is no longer running
				assert.False(t, transport.Running())
			},
		)
	}
}
