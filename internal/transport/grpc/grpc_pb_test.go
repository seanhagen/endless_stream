package grpc

import (
	"context"
	"errors"
	"io"
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

	registerCalls  int
	registeredWith []grpc.ServiceRegistrar

	gatewayRegisterCalls  int
	gatewayRegisteredWith []*runtime.ServeMux

	pingCalls         int
	clientStreamCalls int
	serverStreamCalls int
	biDiStreamCalls   int

	pingHandler         func(context.Context, *proto.PingReq) (*proto.PongResp, error)
	clientStreamHandler func(proto.Test_ClientStreamServer) error
	serverStreamHandler func(*proto.TestRequest, proto.Test_ServerStreamServer) error
	biDiStreamHandler   func(proto.Test_BiDiStreamServer) error
}

// Ping ...
func (ts *testService) Ping(ctx context.Context, req *proto.PingReq) (*proto.PongResp, error) {
	ts.pingCalls++
	return ts.pingHandler(ctx, req)
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
	ts.registerCalls++
	ts.registeredWith = append(ts.registeredWith, srv)
	proto.RegisterTestServer(srv, ts)
}

// RegisterGateway ...
func (ts *testService) RegisterGateway(ctx context.Context, mux *runtime.ServeMux) {
	ts.gatewayRegisterCalls++
	ts.gatewayRegisteredWith = append(ts.gatewayRegisteredWith, mux)
	proto.RegisterTestHandlerServer(ctx, mux, ts)
}

type testPingHandler struct {
	msgs []string
}

// PingHandler ...
func (tph *testPingHandler) PingHandler(
	ctx context.Context,
	pr *proto.PingReq,
) (*proto.PongResp, error) {
	tph.msgs = append(tph.msgs, pr.GetMsg())
	return &proto.PongResp{Gsm: reverseStr(pr.GetMsg())}, nil
}

type testClientStreamHandler struct {
	closeMsg *proto.TestResponse
	msgCount int
	values   map[int]string
}

// ClientStreamHandler ...
func (tcsh *testClientStreamHandler) ClientStreamHandler(srv proto.Test_ClientStreamServer) error {
	for {
		msg, err := srv.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		tcsh.msgCount++

		id := msg.GetChunkId()
		data := msg.GetMsg()

		tcsh.values[int(id)] = data
	}

	return srv.SendAndClose(tcsh.closeMsg)
}

func reverseStr(in string) string {
	n := 0
	rune := make([]rune, len(in))
	for _, r := range in {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
