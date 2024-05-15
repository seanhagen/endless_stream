// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proxy/proxy.proto

package proxy

import (
	context "context"
	common "github.com/seanhagen/endless_stream/internal/proto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Proxy_Game_FullMethodName = "/endless.Proxy/Game"
)

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	Game(ctx context.Context, opts ...grpc.CallOption) (Proxy_GameClient, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) Game(ctx context.Context, opts ...grpc.CallOption) (Proxy_GameClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], Proxy_Game_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyGameClient{stream}
	return x, nil
}

type Proxy_GameClient interface {
	Send(*common.GameRequest) error
	Recv() (*common.GameResponse, error)
	grpc.ClientStream
}

type proxyGameClient struct {
	grpc.ClientStream
}

func (x *proxyGameClient) Send(m *common.GameRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyGameClient) Recv() (*common.GameResponse, error) {
	m := new(common.GameResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations must embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	Game(Proxy_GameServer) error
	mustEmbedUnimplementedProxyServer()
}

// UnimplementedProxyServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) Game(Proxy_GameServer) error {
	return status.Errorf(codes.Unimplemented, "method Game not implemented")
}
func (UnimplementedProxyServer) mustEmbedUnimplementedProxyServer() {}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_Game_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServer).Game(&proxyGameServer{stream})
}

type Proxy_GameServer interface {
	Send(*common.GameResponse) error
	Recv() (*common.GameRequest, error)
	grpc.ServerStream
}

type proxyGameServer struct {
	grpc.ServerStream
}

func (x *proxyGameServer) Send(m *common.GameResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyGameServer) Recv() (*common.GameRequest, error) {
	m := new(common.GameRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endless.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Game",
			Handler:       _Proxy_Game_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proxy/proxy.proto",
}
