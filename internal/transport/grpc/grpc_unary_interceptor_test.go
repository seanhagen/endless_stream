package grpc

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type testUnaryInterceptor struct {
	t        *testing.T
	count    int
	assertFn func(*testing.T, context.Context, any, *grpc.UnaryServerInfo)
}

// intercept ...
func (tui *testUnaryInterceptor) intercept(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	tui.count++
	tui.assertFn(tui.t, ctx, req, info)
	return handler(ctx, req)
}

func TestTransportGRPC_UnaryInterceptor(t *testing.T) {
	ctx := context.TODO()

	logger := logs.NewTestLog(
		t, &logs.Config{Out: io.Discard},
	)

	// configure the NetworkConfig & proto.TestClient
	netConf, client := buildBufferListener(t, ctx)

	testPing := &testPingHandler{}
	svc := &testService{
		pingHandler: testPing.PingHandler,
	}

	expectMetadata := map[string]string{"testing": "a value"}

	unaryTester := &testUnaryInterceptor{
		t: t,
		assertFn: func(t *testing.T, ctx context.Context, req any, info *grpc.UnaryServerInfo) {
			assert.IsType(t, (*proto.PingReq)(nil), req)
			require.NotNil(t, info)

			assert.Equal(t, "/endless.Test/Ping", info.FullMethod)

			md, ok := metadata.FromIncomingContext(ctx)
			assert.True(t, ok, "expected metadata in incoming context")

			value := md.Get("testing")
			expectValue := []string{"a value"}
			assert.Equal(t, expectValue, value)
		},
	}

	// create the transport
	config := Config{
		Logger:   logger,
		Network:  netConf,
		Timeout:  DefaultTimeout,
		Services: []Service{svc},
		UnaryInterceptors: []grpc.UnaryServerInterceptor{
			unaryTester.intercept,
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

	// make a request
	req := &proto.PingReq{
		Msg: "hello world",
	}

	ctxWithTimeout, cancelTimeoutFn := context.WithTimeout(ctx, time.Second*5)
	md := metadata.New(expectMetadata)
	outgoingCtx := metadata.NewOutgoingContext(ctxWithTimeout, md)

	resp, err := client.Ping(outgoingCtx, req)
	t.Cleanup(cancelTimeoutFn)

	// do some asserting
	assert.NoError(t, err)
	assert.Equal(t, "dlrow olleh", resp.GetGsm())
	assert.Equal(t, 1, svc.pingCalls)
	assert.Contains(t, testPing.msgs, "hello world")

	// stop the server, wait a bit for it to finish
	cancelFn()
	time.Sleep(time.Millisecond * 200)

	// then check the server is no longer running
	assert.False(t, transport.Running())

	// now check the unary interceptor
	assert.Equal(t, 1, unaryTester.count)
}
