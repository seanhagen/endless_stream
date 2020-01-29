// Code generated by protoc-gen-go. DO NOT EDIT.
// source: endless.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateGame struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGame) Reset()         { *m = CreateGame{} }
func (m *CreateGame) String() string { return proto.CompactTextString(m) }
func (*CreateGame) ProtoMessage()    {}
func (*CreateGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_602ff85d9c6c040c, []int{0}
}

func (m *CreateGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGame.Unmarshal(m, b)
}
func (m *CreateGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGame.Marshal(b, m, deterministic)
}
func (m *CreateGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGame.Merge(m, src)
}
func (m *CreateGame) XXX_Size() int {
	return xxx_messageInfo_CreateGame.Size(m)
}
func (m *CreateGame) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGame.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGame proto.InternalMessageInfo

type GameCreated struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameCreated) Reset()         { *m = GameCreated{} }
func (m *GameCreated) String() string { return proto.CompactTextString(m) }
func (*GameCreated) ProtoMessage()    {}
func (*GameCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_602ff85d9c6c040c, []int{1}
}

func (m *GameCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCreated.Unmarshal(m, b)
}
func (m *GameCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCreated.Marshal(b, m, deterministic)
}
func (m *GameCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCreated.Merge(m, src)
}
func (m *GameCreated) XXX_Size() int {
	return xxx_messageInfo_GameCreated.Size(m)
}
func (m *GameCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCreated.DiscardUnknown(m)
}

var xxx_messageInfo_GameCreated proto.InternalMessageInfo

func (m *GameCreated) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateGame)(nil), "endless.stream.v1.CreateGame")
	proto.RegisterType((*GameCreated)(nil), "endless.stream.v1.GameCreated")
}

func init() { proto.RegisterFile("endless.proto", fileDescriptor_602ff85d9c6c040c) }

var fileDescriptor_602ff85d9c6c040c = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcd, 0x4b, 0xc9,
	0x49, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x71, 0x8b, 0x4b, 0x8a,
	0x52, 0x13, 0x73, 0xf5, 0xca, 0x0c, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13,
	0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xa0, 0x1a, 0xa4,
	0xb8, 0x33, 0xf3, 0x0a, 0x4a, 0x4b, 0xa0, 0x1c, 0x9e, 0xfc, 0xd2, 0x12, 0x38, 0x4f, 0x89, 0x87,
	0x8b, 0xcb, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xd5, 0x3d, 0x31, 0x37, 0x55, 0x49, 0x91, 0x8b, 0x1b,
	0x44, 0x43, 0x44, 0x52, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa3, 0xb5, 0x8c, 0x5c, 0x5c, 0x20, 0x35, 0xc1, 0xa9, 0x45, 0x65,
	0xa9, 0x45, 0x42, 0xd1, 0x5c, 0x6c, 0x10, 0xd5, 0x42, 0xb2, 0x7a, 0x18, 0xce, 0xd2, 0x43, 0x18,
	0x2d, 0x25, 0x87, 0x45, 0x1a, 0xc9, 0x2e, 0x25, 0xa1, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xf1, 0x08,
	0x71, 0xe9, 0x97, 0x19, 0xea, 0x27, 0x43, 0x8c, 0x74, 0xe2, 0x62, 0x0d, 0x2e, 0x01, 0x31, 0x24,
	0xb0, 0x68, 0xf6, 0x04, 0xf9, 0x49, 0x4a, 0x12, 0x8b, 0x8c, 0x3f, 0xd8, 0x83, 0x4a, 0x0c, 0x1a,
	0x8c, 0x06, 0x8c, 0x4e, 0x46, 0x51, 0x06, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0xc5, 0xa9, 0x89, 0x79, 0x19, 0x89, 0xe9, 0xa9, 0x79, 0xfa, 0x50, 0x6d, 0xf1, 0x10,
	0x6d, 0xfa, 0x49, 0x89, 0xc9, 0xd9, 0xa9, 0x79, 0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0x49, 0x6c,
	0xe0, 0xb0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x73, 0x5b, 0xce, 0x33, 0x78, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServerClient is the client API for GameServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServerClient interface {
	Create(ctx context.Context, in *CreateGame, opts ...grpc.CallOption) (*GameCreated, error)
	State(ctx context.Context, opts ...grpc.CallOption) (GameServer_StateClient, error)
}

type gameServerClient struct {
	cc *grpc.ClientConn
}

func NewGameServerClient(cc *grpc.ClientConn) GameServerClient {
	return &gameServerClient{cc}
}

func (c *gameServerClient) Create(ctx context.Context, in *CreateGame, opts ...grpc.CallOption) (*GameCreated, error) {
	out := new(GameCreated)
	err := c.cc.Invoke(ctx, "/endless.stream.v1.GameServer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerClient) State(ctx context.Context, opts ...grpc.CallOption) (GameServer_StateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GameServer_serviceDesc.Streams[0], "/endless.stream.v1.GameServer/State", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServerStateClient{stream}
	return x, nil
}

type GameServer_StateClient interface {
	Send(*Input) error
	Recv() (*Output, error)
	grpc.ClientStream
}

type gameServerStateClient struct {
	grpc.ClientStream
}

func (x *gameServerStateClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServerStateClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServerServer is the server API for GameServer service.
type GameServerServer interface {
	Create(context.Context, *CreateGame) (*GameCreated, error)
	State(GameServer_StateServer) error
}

// UnimplementedGameServerServer can be embedded to have forward compatible implementations.
type UnimplementedGameServerServer struct {
}

func (*UnimplementedGameServerServer) Create(ctx context.Context, req *CreateGame) (*GameCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGameServerServer) State(srv GameServer_StateServer) error {
	return status.Errorf(codes.Unimplemented, "method State not implemented")
}

func RegisterGameServerServer(s *grpc.Server, srv GameServerServer) {
	s.RegisterService(&_GameServer_serviceDesc, srv)
}

func _GameServer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endless.stream.v1.GameServer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerServer).Create(ctx, req.(*CreateGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServer_State_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServerServer).State(&gameServerStateServer{stream})
}

type GameServer_StateServer interface {
	Send(*Output) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type gameServerStateServer struct {
	grpc.ServerStream
}

func (x *gameServerStateServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServerStateServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GameServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "endless.stream.v1.GameServer",
	HandlerType: (*GameServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GameServer_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "State",
			Handler:       _GameServer_State_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "endless.proto",
}
