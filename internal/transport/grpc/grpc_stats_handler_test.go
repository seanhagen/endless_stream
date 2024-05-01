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
	"google.golang.org/grpc/stats"
)

func TestTransportGRPC_StatsHandler(t *testing.T) {
	ctx := context.TODO()

	logger := logs.NewTestLog(
		t, &logs.Config{Out: io.Discard},
	)

	// configure the NetworkConfig & proto.TestClient
	netConf, client := buildBufferListener(t, ctx)

	testPing := func(_ context.Context, _ *proto.PingReq) (*proto.PongResp, error) {
		return &proto.PongResp{}, nil
	}
	svc := &testService{
		pingHandler: testPing,
	}

	tagRPCCalled := false
	tagConnCalled := false

	handleRPCCalled := 0
	expectHandleRPC := 7
	handleConnCalled := 0
	expectHandleConn := 2

	testStatHandler := &testStatsHandler{
		t: t,
		tagRpc: func(ctx context.Context, tagInfo *stats.RPCTagInfo) context.Context {
			tagRPCCalled = true
			return ctx
		},

		handleRpc: func(ctx context.Context, data stats.RPCStats) {
			handleRPCCalled++
		},

		tagConn: func(ctx context.Context, tagInfo *stats.ConnTagInfo) context.Context {
			tagConnCalled = true
			return ctx
		},

		handleConn: func(ctx context.Context, data stats.ConnStats) {
			handleConnCalled++
		},
	}

	// create the transport
	config := Config{
		Logger:       logger,
		Network:      netConf,
		Timeout:      DefaultTimeout,
		Services:     []Service{svc},
		StatsHandler: testStatHandler,
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
	req := &proto.PingReq{Msg: "hello world"}

	ctxWithTimeout, cancelTimeoutFn := context.WithTimeout(ctx, time.Second*5)
	resp, err := client.Ping(ctxWithTimeout, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Cleanup(cancelTimeoutFn)

	// stop the server, wait a bit for it to finish
	cancelFn()
	time.Sleep(time.Millisecond * 200)

	assert.True(t, tagRPCCalled, "expected tag rpc method to be called")
	assert.True(t, tagConnCalled, "expected tag conn method to be called")
	assert.Equal(
		t, expectHandleConn, handleConnCalled,
		"expected 'handle conn' to be called certain number of times",
	)
	assert.Equal(
		t, expectHandleRPC, handleRPCCalled,
		"expected 'handle rpc' to be called certain number of times",
	)
}

func TestTransportGRPC_testStatsHandler(t *testing.T) {
	assert.Implements(t, (*stats.Handler)(nil), (*testStatsHandler)(nil))
}

type testStatsHandler struct {
	t *testing.T

	tagRpc    func(context.Context, *stats.RPCTagInfo) context.Context
	handleRpc func(context.Context, stats.RPCStats)

	tagConn    func(context.Context, *stats.ConnTagInfo) context.Context
	handleConn func(context.Context, stats.ConnStats)
}

// TagRPC ...
func (tsh testStatsHandler) TagRPC(ctx context.Context, tagInfo *stats.RPCTagInfo) context.Context {
	return tsh.tagRpc(ctx, tagInfo)
}

// HandleRPC ...
func (tsh testStatsHandler) HandleRPC(ctx context.Context, stats stats.RPCStats) {
	tsh.handleRpc(ctx, stats)
}

// TagConn ...
func (tsh testStatsHandler) TagConn(
	ctx context.Context,
	connInfo *stats.ConnTagInfo,
) context.Context {
	return tsh.tagConn(ctx, connInfo)
}

// HandleConn ...
func (tsh testStatsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	tsh.handleConn(ctx, connStats)
}
