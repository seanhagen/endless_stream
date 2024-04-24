package grpc

import (
	"context"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func Test_testService_ImplementsGRPCTestServer(t *testing.T) {
	assert.Implements(t, (*proto.TestServer)(nil), &testService{})
}

type testService struct {
	proto.UnimplementedTestServer

	pingCalls         int
	clientStreamCalls int
	serverStreamCalls int
	biDiStreamCalls   int

	pingHandler         func(*proto.PingReq) (*proto.PongResp, error)
	clientStreamHandler func(proto.Test_ClientStreamServer) error
	serverStreamHandler func(*proto.TestRequest, proto.Test_ServerStreamServer) error
	biDiStreamHandler   func(proto.Test_BiDiStreamServer) error
}

// Ping ...
func (ts *testService) Ping(_ context.Context, req *proto.PingReq) (*proto.PongResp, error) {
	ts.pingCalls++
	return ts.pingHandler(req)
}

// ClientStream ...
func (ts *testService) ClientStream(srv proto.Test_ClientStreamServer) error {
	ts.clientStreamCalls++
	return ts.clientStreamHandler(srv)
}

// ServerStream ...
func (ts *testService) ServerStream(
	req *proto.TestRequest,
	srv proto.Test_ServerStreamServer,
) error {
	ts.serverStreamCalls++
	return ts.serverStreamHandler(req, srv)
}

// BiDiStream ...
func (ts *testService) BiDiStream(srv proto.Test_BiDiStreamServer) error {
	ts.biDiStreamCalls++
	return ts.biDiStreamHandler(srv)
}

// Register ...
func (ts *testService) Register(srv grpc.ServiceRegistrar) {
	proto.RegisterTestServer(srv, ts)
}

// RegisterGateway ...
func (ts *testService) RegisterGateway(ctx context.Context, mux *runtime.ServeMux) {
	proto.RegisterTestHandlerServer(ctx, mux, ts)
}

type testPingHandler struct {
	msgs []string
}

// PingHandler ...
func (tph *testPingHandler) PingHandler(pr *proto.PingReq) (*proto.PongResp, error) {
	msg := pr.GetMsg()

	tph.msgs = append(tph.msgs, msg)

	n := 0
	rune := make([]rune, len(msg))
	for _, r := range msg {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)

	return &proto.PongResp{Gsm: output}, nil
}
