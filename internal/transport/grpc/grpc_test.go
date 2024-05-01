package grpc

import (
	"context"
	"fmt"
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

	netConf := WithCustomListener(lis)

	config := Config{
		Network:  netConf,
		Timeout:  DefaultTimeout, // time.Duration
		Services: []Service{svc},
	}

	transport, err := BuildTransport(ctx, config)

	require.NoError(t, err)
	require.NotNil(t, transport)

	ctxWithCancel, cancelFn := context.WithCancel(ctx)

	err = transport.Start(ctxWithCancel)
	require.NoError(t, err)

	assert.True(t, transport.Running())

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
