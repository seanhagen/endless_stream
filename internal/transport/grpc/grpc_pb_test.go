package grpc

import (
	"context"
	"testing"

	"github.com/seanhagen/endless_stream/internal/proto"
	"github.com/stretchr/testify/assert"
)

func Test_testService_ImplementsGRPCTestServer(t *testing.T) {
	assert.Implements(t, (*proto.TestServer)(nil), &testService{})
}

type testService struct {
	proto.UnimplementedTestServer

	PingHandler         func(*proto.PingReq) (*proto.PongResp, error)
	ClientStreamHandler func(proto.Test_ClientStreamServer) error
	ServerStreamHandler func(*proto.TestRequest, proto.Test_ServerStreamServer) error
	BiDiStreamHandler   func(proto.Test_BiDiStreamServer) error
}

// Ping ...
func (ts *testService) Ping(_ context.Context, req *proto.PingReq) (*proto.PongResp, error) {
	return ts.PingHandler(req)
}

// ClientStream ...
func (ts *testService) ClientStream(srv proto.Test_ClientStreamServer) error {
	return ts.ClientStreamHandler(srv)
}

// ServerStream ...
func (ts *testService) ServerStream(
	req *proto.TestRequest,
	srv proto.Test_ServerStreamServer,
) error {
	return ts.ServerStreamHandler(req, srv)
}

// BiDiStream ...
func (ts *testService) BiDiStream(srv proto.Test_BiDiStreamServer) error {
	return ts.BiDiStreamHandler(srv)
}
