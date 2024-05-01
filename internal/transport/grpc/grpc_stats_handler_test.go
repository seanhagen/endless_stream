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
			// fmt.Printf("==============================\ntag rpc -- ")
			// fmt.Printf("method: %q, fail fast: %v\n", tagInfo.FullMethodName, tagInfo.FailFast)
			// fmt.Printf("==============================\n")
			return ctx
		},

		handleRpc: func(ctx context.Context, data stats.RPCStats) {
			handleRPCCalled++

			// fmt.Printf("******************************\nhandle rpc -- ")
			// switch data.(type) {
			// case *stats.InHeader:
			// 	fmt.Printf("in header!\n")
			// case *stats.Begin:
			// 	fmt.Printf("begin!\n")
			// case *stats.InPayload:
			// 	fmt.Printf("in payload!\n")
			// case *stats.OutHeader:
			// 	fmt.Printf("out header!\n")
			// case *stats.OutPayload:
			// 	fmt.Printf("out payload!\n")
			// case *stats.OutTrailer:
			// 	fmt.Printf("out trailer!\n")
			// case *stats.End:
			// 	fmt.Printf("end!\n")
			// default:
			// 	spew.Dump(data)
			// }

			// // spew.Dump(ctx, stats)
			// fmt.Printf("******************************\n")
		},

		tagConn: func(ctx context.Context, tagInfo *stats.ConnTagInfo) context.Context {
			tagConnCalled = true
			// fmt.Printf("##############################\ntag conn -- ")
			// fmt.Printf(
			// 	"tagging connection, local: %q, remote: %q\n",
			// 	tagInfo.LocalAddr.String(),
			// 	tagInfo.RemoteAddr.String(),
			// )
			// fmt.Printf("##############################\n")
			return ctx
		},

		handleConn: func(ctx context.Context, data stats.ConnStats) {
			handleConnCalled++
			// fmt.Printf("------------------------------\nhandle conn -- ")
			// switch data.(type) {
			// case *stats.ConnBegin:
			// 	fmt.Printf("conn begin!\n")
			// case *stats.ConnEnd:
			// 	fmt.Printf("conn end!\n")
			// default:
			// 	spew.Dump(data)
			// }
			// fmt.Printf("------------------------------\n")
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
