package grpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type testStreamInterceptor struct {
	t        *testing.T
	count    int
	assertFn func(*testing.T, context.Context, *grpc.StreamServerInfo)
}

// intercept ...
func (tui *testStreamInterceptor) intercept(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	tui.count++

	tui.assertFn(tui.t, ss.Context(), info)
	return handler(srv, ss)
}

type grpcStreamInterceptorTestCase struct {
	name                   string
	buildServiceTester     func(*testing.T) (*testService, func(*testing.T))
	buildInterceptorTester func(*testing.T, *bool) (*testStreamInterceptor, func(*testing.T))
	runClient              func(*testing.T, context.Context, proto.TestClient)
}

func TestTransportGRPC_StreamInterceptor(t *testing.T) {
	ctx := context.TODO()

	tests := []grpcStreamInterceptorTestCase{
		buildClientStreamTestCase(t),
		buildServerStreamTestCase(t),
		buildBiDiStreamTestCase(t),
	}

	for i, x := range tests {
		tt := x
		t.Run(
			fmt.Sprintf("test %d %s", i+1, tt.name),
			func(t *testing.T) {
				logger := logs.NewTestLog(
					t, &logs.Config{Out: io.Discard},
				)

				netConf, client := buildPortListener(t, ctx)
				svc, serviceValidation := tt.buildServiceTester(t)

				streamTesterCalled := false
				streamTester, validateStreamFn := tt.buildInterceptorTester(t, &streamTesterCalled)

				config := Config{
					Logger:   logger,
					Network:  netConf,
					Timeout:  DefaultTimeout,
					Services: []Service{svc},
					StreamInterceptors: []grpc.StreamServerInterceptor{
						streamTester.intercept,
					},
				}

				transport, err := BuildTransport(ctx, config)
				require.NoError(t, err, "unable to build transport from config")
				require.NotNil(t, transport, "expected non-nil transport")

				// start the transport
				ctxWithCancel, cancelFn := context.WithCancel(ctx)
				t.Cleanup(cancelFn)
				err = transport.Start(ctxWithCancel)
				require.NoError(t, err)
				assert.True(t, transport.Running())
				assert.NotZero(t, svc.registerCalls)

				// setup the stream
				// ctxWithTimeout, cancelTimeoutFn := context.WithTimeout(ctx, time.Second)
				tt.runClient(t, ctx, client)

				// time.Sleep(time.Second)

				// fmt.Printf("cancelling ctx!\n")
				// cancelTimeoutFn()

				// stop the server, wait a bit for it to finish
				cancelFn()
				time.Sleep(time.Millisecond * 200)
				assert.False(t, transport.Running())

				// test the stream interceptor was called
				assert.True(
					t, streamTesterCalled,
					"expected stream interceptor test struct method to be called",
				)

				// now do some final test validation
				validateStreamFn(t)
				serviceValidation(t)
			},
		)
	}
}

func buildBiDiStreamTestCase(t *testing.T) grpcStreamInterceptorTestCase {
	toSend := map[int]string{
		3: "third",
		1: "should be first",
		2: "middle",
	}

	expectResponses := map[int]string{}
	for k, v := range toSend {
		expectResponses[k] = reverseStr(v)
	}

	return grpcStreamInterceptorTestCase{
		name: "testing stream interceptor with bi-directional stream",
		buildServiceTester: func(*testing.T) (*testService, func(*testing.T)) {
			t.Helper()

			svc := &testService{
				biDiStreamHandler: func(srv proto.Test_BiDiStreamServer) error {
					for {
						msg, err := srv.Recv()
						if errors.Is(err, io.EOF) {
							break
						}
						if err != nil {
							return err
						}
						require.NotNil(
							t, msg,
							"expected non-nil message from bi-directional server Recv() call",
						)

						resp := &proto.TestStreamResponse{
							RespId: msg.GetChunkId(),
							Gsm:    reverseStr(msg.GetMsg()),
						}

						if err := srv.SendMsg(resp); err != nil {
							return err
						}
					}

					return nil
				},
			}

			validate := func(t *testing.T) {
				t.Helper()
				assert.Equal(t, 1, svc.biDiStreamCalls)
			}

			return svc, validate
		},
		buildInterceptorTester: func(t *testing.T, called *bool) (*testStreamInterceptor, func(*testing.T)) {
			t.Helper()

			tsi := &testStreamInterceptor{
				t:     t,
				count: 0,
				assertFn: func(t *testing.T, ctx context.Context, ssi *grpc.StreamServerInfo) {
					t.Helper()

					*called = true

					ssi.IsServerStream = true
					ssi.IsClientStream = true

					assert.Equal(t, "/endless.Test/BiDiStream", ssi.FullMethod)
				},
			}

			validate := func(t *testing.T) {
				t.Helper()
				assert.Equal(t, 1, tsi.count)
			}

			return tsi, validate
		},
		runClient: func(t *testing.T, ctx context.Context, client proto.TestClient) {
			strm, err := client.BiDiStream(ctx)
			require.NoError(t, err, "unable to open bi-directional stream")

			received := map[int]string{}
			lock := sync.Mutex{}
			waitc := make(chan struct{})

			for idx, val := range toSend {
				msg := &proto.TestStreamRequest{
					ChunkId: int32(idx),
					Msg:     val,
				}

				err := strm.Send(msg)
				assert.NoError(
					t, err,
					"unexpected error from bi-directional stream client send",
				)
			}

			go func() {
				msg := &proto.TestStreamRequest{}
				for {
					err := strm.RecvMsg(msg)
					if errors.Is(err, io.EOF) {
						close(waitc)
						break
					}
					assert.NoError(t, err, "unexpected error calling strm.RecvMsg")

					lock.Lock()
					received[int(msg.GetChunkId())] = msg.GetMsg()
					lock.Unlock()
				}
			}()

			resp, err := strm.CloseAndRecv()
			require.NoError(t, err)
			require.NotNil(t, resp)

			lock.Lock()
			received[int(resp.GetRespId())] = resp.GetGsm()
			lock.Unlock()

			assert.Equal(t, expectResponses, received)
		},
	}
}

func buildServerStreamTestCase(t *testing.T) grpcStreamInterceptorTestCase {
	return grpcStreamInterceptorTestCase{
		name: "testing stream interceptor with server stream",
		buildServiceTester: func(t *testing.T) (*testService, func(*testing.T)) {
			t.Helper()

			svc := &testService{
				serverStreamHandler: func(req *proto.TestRequest, srv proto.Test_ServerStreamServer) error {
					require.NotNil(t, req, "expected non-nil request")
					assert.Equal(t, "testing", req.GetName())

					srv.Send(&proto.TestStreamResponse{
						RespId: 3,
						Gsm:    "three",
					})
					srv.Send(&proto.TestStreamResponse{
						RespId: 1,
						Gsm:    "one",
					})
					srv.Send(&proto.TestStreamResponse{
						RespId: 2,
						Gsm:    "two",
					})

					return io.EOF
				},
			}

			validate := func(t *testing.T) {
				t.Helper()
				assert.Equal(t, 1, svc.serverStreamCalls)
			}

			return svc, validate
		},
		buildInterceptorTester: func(t *testing.T, called *bool) (*testStreamInterceptor, func(*testing.T)) {
			t.Helper()

			tsi := &testStreamInterceptor{
				t:     t,
				count: 0,
				assertFn: func(t *testing.T, ctx context.Context, ssi *grpc.StreamServerInfo) {
					t.Helper()
					*called = true
					assert.True(t, ssi.IsServerStream, "expected IsServerStream to be true")
					assert.False(t, ssi.IsClientStream, "expected IsClientStream to be false")
					assert.Equal(t, "/endless.Test/ServerStream", ssi.FullMethod)
				},
			}

			validate := func(t *testing.T) {
				t.Helper()
				assert.Equal(t, 1, tsi.count)
			}

			return tsi, validate
		},
		runClient: func(t *testing.T, ctx context.Context, client proto.TestClient) {
			req := &proto.TestRequest{
				Name: "testing",
			}

			strm, err := client.ServerStream(ctx, req)
			require.NoError(t, err, "unable to start server stream request")

			responses := map[int]string{}
			expectResponses := map[int]string{
				1: "one",
				2: "two",
				3: "three",
			}

			for {
				msg, err := strm.Recv()
				if err == io.EOF {
					break
				}
				require.NoError(t, err, "expected no error from client.ServerStream().Recv()")
				require.NotNil(t, msg, "expected non-nil message from client.ServerStream().Recv()")

				responses[int(msg.GetRespId())] = msg.GetGsm()
			}

			assert.Equal(t, expectResponses, responses)
		},
	}
}

func buildClientStreamTestCase(t *testing.T) grpcStreamInterceptorTestCase {
	return grpcStreamInterceptorTestCase{
		name: "testing stream interceptor with client stream",
		buildServiceTester: func(t *testing.T) (*testService, func(*testing.T)) {
			t.Helper()

			clientStreamCalls := 0

			svc := &testService{
				clientStreamHandler: func(css proto.Test_ClientStreamServer) error {
					for {
						msg, err := css.Recv()
						if err == io.EOF {
							break
						}
						require.NoError(
							t,
							err,
							"unexpected error receiving msg in client stream test",
						)
						require.NotNil(t, msg, "expected non-nil request in client stream test")
						clientStreamCalls++
					}
					return css.SendAndClose(&proto.TestResponse{Resp: "goodbye"})
				},
			}

			validate := func(t *testing.T) {
				assert.Equal(t, 1, svc.clientStreamCalls)
				assert.Equal(t, 2, clientStreamCalls)
			}

			return svc, validate
		},
		buildInterceptorTester: func(t *testing.T, called *bool) (*testStreamInterceptor, func(*testing.T)) {
			t.Helper()

			tsi := &testStreamInterceptor{
				t: t,
				assertFn: func(t *testing.T, ctx context.Context, ssi *grpc.StreamServerInfo) {
					t.Helper()
					*called = true
					assert.NotNil(t, ssi)

					assert.True(t, ssi.IsClientStream, "expected IsClientStream to be true")
					assert.False(t, ssi.IsServerStream, "expected IsServerStream to be true")
					assert.Equal(t, "/endless.Test/ClientStream", ssi.FullMethod)
				},
			}

			validate := func(t *testing.T) {
				t.Helper()
				assert.Equal(t, 1, tsi.count)
			}

			return tsi, validate
		},
		runClient: func(t *testing.T, ctx context.Context, client proto.TestClient) {
			clientStream, err := client.ClientStream(ctx)
			require.NoError(t, err, "unable to open client stream")

			t.Log("sending client stream requests via GRPC client")
			test1 := &proto.TestStreamRequest{ChunkId: 1, Msg: "one"}
			err = clientStream.Send(test1)
			require.NoError(t, err, "unable to send test stream request")

			test2 := &proto.TestStreamRequest{ChunkId: 2, Msg: "two"}
			err = clientStream.Send(test2)

			// close the stream client
			resp, err := clientStream.CloseAndRecv()
			require.NoError(t, err, "unexpected error from clientStream.CloseAndRecv")
			require.NotNil(t, resp, "expected non-nil response from CloseAndRecv")

			// check we get the expected close message
			expectCloseMsg := &proto.TestResponse{
				Resp: "goodbye",
			}
			assert.Equal(t, expectCloseMsg.GetResp(), resp.GetResp())
		},
	}
}
