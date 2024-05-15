// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/game.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_common_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_common_game_proto_rawDescGZIP(), []int{0}
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_common_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_common_game_proto_rawDescGZIP(), []int{1}
}

// Level is a single layer of tiles laid out to create a level.
type Level struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tiles is an array of all the tiles laid out on a level.
	Tiles []*Tile `protobuf:"bytes,1,rep,name=tiles,proto3" json:"tiles,omitempty"`
	// Tileset tells Godot what tileset to use.
	Tileset Tileset `protobuf:"varint,2,opt,name=tileset,proto3,enum=endless.Tileset" json:"tileset,omitempty"`
}

func (x *Level) Reset() {
	*x = Level{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Level) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Level) ProtoMessage() {}

func (x *Level) ProtoReflect() protoreflect.Message {
	mi := &file_common_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Level.ProtoReflect.Descriptor instead.
func (*Level) Descriptor() ([]byte, []int) {
	return file_common_game_proto_rawDescGZIP(), []int{2}
}

func (x *Level) GetTiles() []*Tile {
	if x != nil {
		return x.Tiles
	}
	return nil
}

func (x *Level) GetTileset() Tileset {
	if x != nil {
		return x.Tileset
	}
	return Tileset_Dungeon
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beat *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=beat,proto3" json:"beat,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_common_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_common_game_proto_rawDescGZIP(), []int{3}
}

func (x *Heartbeat) GetBeat() *timestamppb.Timestamp {
	if x != nil {
		return x.Beat
	}
	return nil
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Types that are assignable to Request:
	//
	//	*GameRequest_Ping
	Request isGameRequest_Request `protobuf_oneof:"request"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_common_game_proto_rawDescGZIP(), []int{4}
}

func (x *GameRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (m *GameRequest) GetRequest() isGameRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *GameRequest) GetPing() *Ping {
	if x, ok := x.GetRequest().(*GameRequest_Ping); ok {
		return x.Ping
	}
	return nil
}

type isGameRequest_Request interface {
	isGameRequest_Request()
}

type GameRequest_Ping struct {
	Ping *Ping `protobuf:"bytes,100,opt,name=ping,proto3,oneof"`
}

func (*GameRequest_Ping) isGameRequest_Request() {}

type GameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Types that are assignable to Message:
	//
	//	*GameResponse_Heartbeat
	//	*GameResponse_Log
	//	*GameResponse_Level
	//	*GameResponse_Pong
	Message isGameResponse_Message `protobuf_oneof:"message"`
}

func (x *GameResponse) Reset() {
	*x = GameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResponse) ProtoMessage() {}

func (x *GameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResponse.ProtoReflect.Descriptor instead.
func (*GameResponse) Descriptor() ([]byte, []int) {
	return file_common_game_proto_rawDescGZIP(), []int{5}
}

func (x *GameResponse) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (m *GameResponse) GetMessage() isGameResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *GameResponse) GetHeartbeat() *Heartbeat {
	if x, ok := x.GetMessage().(*GameResponse_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (x *GameResponse) GetLog() *Log {
	if x, ok := x.GetMessage().(*GameResponse_Log); ok {
		return x.Log
	}
	return nil
}

func (x *GameResponse) GetLevel() *Level {
	if x, ok := x.GetMessage().(*GameResponse_Level); ok {
		return x.Level
	}
	return nil
}

func (x *GameResponse) GetPong() *Pong {
	if x, ok := x.GetMessage().(*GameResponse_Pong); ok {
		return x.Pong
	}
	return nil
}

type isGameResponse_Message interface {
	isGameResponse_Message()
}

type GameResponse_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,2,opt,name=heartbeat,proto3,oneof"`
}

type GameResponse_Log struct {
	Log *Log `protobuf:"bytes,3,opt,name=log,proto3,oneof"`
}

type GameResponse_Level struct {
	Level *Level `protobuf:"bytes,5,opt,name=level,proto3,oneof"`
}

type GameResponse_Pong struct {
	Pong *Pong `protobuf:"bytes,100,opt,name=pong,proto3,oneof"`
}

func (*GameResponse_Heartbeat) isGameResponse_Message() {}

func (*GameResponse_Log) isGameResponse_Message() {}

func (*GameResponse_Level) isGameResponse_Message() {}

func (*GameResponse_Pong) isGameResponse_Message() {}

var File_common_game_proto protoreflect.FileDescriptor

var file_common_game_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x06,
	0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x58, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x23, 0x0a, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x74,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x54, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x74,
	0x22, 0x3b, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x62, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x62, 0x65, 0x61, 0x74, 0x22, 0x5a, 0x0a,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x09,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd9, 0x01, 0x0a, 0x0c, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00,
	0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x6c,
	0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x26, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x6f,
	0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x8f, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x73, 0x42, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x65, 0x61, 0x6e, 0x68, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03,
	0x45, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0xca, 0x02, 0x07,
	0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x13, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07,
	0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_game_proto_rawDescOnce sync.Once
	file_common_game_proto_rawDescData = file_common_game_proto_rawDesc
)

func file_common_game_proto_rawDescGZIP() []byte {
	file_common_game_proto_rawDescOnce.Do(func() {
		file_common_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_game_proto_rawDescData)
	})
	return file_common_game_proto_rawDescData
}

var file_common_game_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_game_proto_goTypes = []interface{}{
	(*Ping)(nil),                  // 0: endless.Ping
	(*Pong)(nil),                  // 1: endless.Pong
	(*Level)(nil),                 // 2: endless.Level
	(*Heartbeat)(nil),             // 3: endless.Heartbeat
	(*GameRequest)(nil),           // 4: endless.GameRequest
	(*GameResponse)(nil),          // 5: endless.GameResponse
	(*Tile)(nil),                  // 6: endless.Tile
	(Tileset)(0),                  // 7: endless.Tileset
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Log)(nil),                   // 9: endless.Log
}
var file_common_game_proto_depIdxs = []int32{
	6, // 0: endless.Level.tiles:type_name -> endless.Tile
	7, // 1: endless.Level.tileset:type_name -> endless.Tileset
	8, // 2: endless.Heartbeat.beat:type_name -> google.protobuf.Timestamp
	0, // 3: endless.GameRequest.ping:type_name -> endless.Ping
	3, // 4: endless.GameResponse.heartbeat:type_name -> endless.Heartbeat
	9, // 5: endless.GameResponse.log:type_name -> endless.Log
	2, // 6: endless.GameResponse.level:type_name -> endless.Level
	1, // 7: endless.GameResponse.pong:type_name -> endless.Pong
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_common_game_proto_init() }
func file_common_game_proto_init() {
	if File_common_game_proto != nil {
		return
	}
	file_common_tile_proto_init()
	file_common_logs_proto_init()
	file_common_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Level); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_common_game_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GameRequest_Ping)(nil),
	}
	file_common_game_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*GameResponse_Heartbeat)(nil),
		(*GameResponse_Log)(nil),
		(*GameResponse_Level)(nil),
		(*GameResponse_Pong)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_game_proto_goTypes,
		DependencyIndexes: file_common_game_proto_depIdxs,
		MessageInfos:      file_common_game_proto_msgTypes,
	}.Build()
	File_common_game_proto = out.File
	file_common_game_proto_rawDesc = nil
	file_common_game_proto_goTypes = nil
	file_common_game_proto_depIdxs = nil
}
