package grpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/endless_stream/internal/observability/logs"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/stats"
)

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

func TestTransportGRPC_testStatsHandler(t *testing.T) {
	assert.Implements(t, (*stats.Handler)(nil), (*testStatsHandler)(nil))
}

func TestTransportGRPC_StatsHandler(t *testing.T) {
	ctx := context.TODO()

	logger := logs.NewTestLog(
		t,
		&logs.Config{
			Out:   logs.NewTestLogOutput(t, true),
			Level: logs.LevelDebug,
		},
	)

	// configure the NetworkConfig & proto.TestClient
	t.Log("setting up network configuration & test client")
	netConf, client := buildBufferListener(t, ctx)

	testPing := func(ctx context.Context, req *proto.PingReq) (*proto.PongResp, error) {
		return nil, fmt.Errorf("not yet")
	}
	svc := &testService{
		pingHandler: testPing,
	}

	testStatHandler := &testStatsHandler{
		t: t,
		tagRpc: func(ctx context.Context, tagInfo *stats.RPCTagInfo) context.Context {
			spew.Dump(ctx, tagInfo)
			return ctx
		},
		handleRpc: func(ctx context.Context, stats stats.RPCStats) {
			spew.Dump(ctx, stats)
		},
		tagConn: func(ctx context.Context, tagInfo *stats.ConnTagInfo) context.Context {
			spew.Dump(ctx, tagInfo)
			return ctx
		},
		handleConn: func(ctx context.Context, stats stats.ConnStats) {
			spew.Dump(ctx, stats)
		},
	}

	// create the transport
	t.Log("creating transport")
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
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Cleanup(cancelTimeoutFn)

	// stop the server, wait a bit for it to finish
	cancelFn()
	time.Sleep(time.Millisecond * 200)
}
