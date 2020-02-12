// Code generated by protoc-gen-go. DO NOT EDIT.
// source: endless.proto

package endless

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
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xae, 0xa5, 0x52, 0x89, 0x6b, 0x19, 0xf0, 0x54, 0x22, 0x40, 0x90, 0xa9, 0x93, 0x0d, 0x54,
	0xbc, 0x40, 0x19, 0x50, 0x27, 0x24, 0xd8, 0x58, 0xd0, 0x25, 0x39, 0xa5, 0x11, 0xcd, 0x5d, 0x94,
	0x5c, 0x78, 0x13, 0xde, 0x17, 0x39, 0x49, 0x01, 0xa9, 0x99, 0xec, 0xef, 0xbe, 0x1f, 0x7f, 0x3e,
	0x38, 0x23, 0xce, 0xf6, 0xd4, 0x34, 0xae, 0xaa, 0x45, 0xc5, 0x9e, 0x1f, 0x60, 0xa3, 0x35, 0x61,
	0xe9, 0xbe, 0xee, 0xa3, 0xcb, 0x5c, 0x24, 0xdf, 0x93, 0xc7, 0xaa, 0xf0, 0xc8, 0x2c, 0x8a, 0x5a,
	0x08, 0x0f, 0x86, 0x68, 0x5e, 0x70, 0xd5, 0xea, 0x00, 0x16, 0xd2, 0xea, 0x2f, 0x8a, 0x17, 0x00,
	0x4f, 0x35, 0xa1, 0xd2, 0x33, 0x96, 0x14, 0xdf, 0xc2, 0x3c, 0x9c, 0xfd, 0x24, 0xb3, 0x16, 0xa6,
	0xa9, 0x64, 0xb4, 0x34, 0x37, 0x66, 0x75, 0xfa, 0xda, 0xdd, 0x1f, 0xbe, 0x0d, 0x4c, 0x83, 0xc6,
	0x6e, 0x61, 0xd6, 0xeb, 0xec, 0x95, 0x3b, 0x2a, 0xe4, 0xfe, 0x42, 0xa3, 0xeb, 0x11, 0xfa, 0xdf,
	0x2b, 0xf1, 0xc4, 0x6e, 0xe0, 0xe4, 0x4d, 0x43, 0xd2, 0x72, 0x44, 0xba, 0x0d, 0xdd, 0xa3, 0x8b,
	0x11, 0xe6, 0xa5, 0xfb, 0x48, 0x3c, 0x59, 0x99, 0x3b, 0xb3, 0x79, 0x7c, 0x5f, 0xe7, 0x85, 0xee,
	0xda, 0xc4, 0xa5, 0x52, 0xfa, 0x86, 0x90, 0x77, 0x98, 0x13, 0xfb, 0xc1, 0xf6, 0xd1, 0xdb, 0x7c,
	0x82, 0xe9, 0x27, 0x71, 0x76, 0x18, 0x27, 0xb3, 0x6e, 0x0d, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0x03, 0xf6, 0x65, 0x63, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	Create(ctx context.Context, in *CreateGame, opts ...grpc.CallOption) (*GameCreated, error)
	State(ctx context.Context, opts ...grpc.CallOption) (Game_StateClient, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Create(ctx context.Context, in *CreateGame, opts ...grpc.CallOption) (*GameCreated, error) {
	out := new(GameCreated)
	err := c.cc.Invoke(ctx, "/endless.stream.v1.Game/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) State(ctx context.Context, opts ...grpc.CallOption) (Game_StateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Game_serviceDesc.Streams[0], "/endless.stream.v1.Game/State", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameStateClient{stream}
	return x, nil
}

type Game_StateClient interface {
	Send(*Input) error
	Recv() (*Output, error)
	grpc.ClientStream
}

type gameStateClient struct {
	grpc.ClientStream
}

func (x *gameStateClient) Send(m *Input) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameStateClient) Recv() (*Output, error) {
	m := new(Output)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	Create(context.Context, *CreateGame) (*GameCreated, error)
	State(Game_StateServer) error
}

// UnimplementedGameServer can be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (*UnimplementedGameServer) Create(ctx context.Context, req *CreateGame) (*GameCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGameServer) State(srv Game_StateServer) error {
	return status.Errorf(codes.Unimplemented, "method State not implemented")
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endless.stream.v1.Game/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Create(ctx, req.(*CreateGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_State_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).State(&gameStateServer{stream})
}

type Game_StateServer interface {
	Send(*Output) error
	Recv() (*Input, error)
	grpc.ServerStream
}

type gameStateServer struct {
	grpc.ServerStream
}

func (x *gameStateServer) Send(m *Output) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameStateServer) Recv() (*Input, error) {
	m := new(Input)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "endless.stream.v1.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Game_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "State",
			Handler:       _Game_State_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "endless.proto",
}
