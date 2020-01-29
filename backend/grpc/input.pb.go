// Code generated by protoc-gen-go. DO NOT EDIT.
// source: input.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Move_Dir int32

const (
	Move_Left  Move_Dir = 0
	Move_Right Move_Dir = 1
)

var Move_Dir_name = map[int32]string{
	0: "Left",
	1: "Right",
}

var Move_Dir_value = map[string]int32{
	"Left":  0,
	"Right": 1,
}

func (x Move_Dir) String() string {
	return proto.EnumName(Move_Dir_name, int32(x))
}

func (Move_Dir) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{5, 0}
}

type Register struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{0}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Register) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CharSelect struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Choice               Class    `protobuf:"varint,2,opt,name=choice,proto3,enum=endless.stream.v1.Class" json:"choice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharSelect) Reset()         { *m = CharSelect{} }
func (m *CharSelect) String() string { return proto.CompactTextString(m) }
func (*CharSelect) ProtoMessage()    {}
func (*CharSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{1}
}

func (m *CharSelect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharSelect.Unmarshal(m, b)
}
func (m *CharSelect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharSelect.Marshal(b, m, deterministic)
}
func (m *CharSelect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharSelect.Merge(m, src)
}
func (m *CharSelect) XXX_Size() int {
	return xxx_messageInfo_CharSelect.Size(m)
}
func (m *CharSelect) XXX_DiscardUnknown() {
	xxx_messageInfo_CharSelect.DiscardUnknown(m)
}

var xxx_messageInfo_CharSelect proto.InternalMessageInfo

func (m *CharSelect) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *CharSelect) GetChoice() Class {
	if m != nil {
		return m.Choice
	}
	return Class_Unknown
}

type GameStart struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStart) Reset()         { *m = GameStart{} }
func (m *GameStart) String() string { return proto.CompactTextString(m) }
func (*GameStart) ProtoMessage()    {}
func (*GameStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{2}
}

func (m *GameStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStart.Unmarshal(m, b)
}
func (m *GameStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStart.Marshal(b, m, deterministic)
}
func (m *GameStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStart.Merge(m, src)
}
func (m *GameStart) XXX_Size() int {
	return xxx_messageInfo_GameStart.Size(m)
}
func (m *GameStart) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStart.DiscardUnknown(m)
}

var xxx_messageInfo_GameStart proto.InternalMessageInfo

type UseSkill struct {
	SkillId              string   `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	TargetId             string   `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseSkill) Reset()         { *m = UseSkill{} }
func (m *UseSkill) String() string { return proto.CompactTextString(m) }
func (*UseSkill) ProtoMessage()    {}
func (*UseSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{3}
}

func (m *UseSkill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseSkill.Unmarshal(m, b)
}
func (m *UseSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseSkill.Marshal(b, m, deterministic)
}
func (m *UseSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseSkill.Merge(m, src)
}
func (m *UseSkill) XXX_Size() int {
	return xxx_messageInfo_UseSkill.Size(m)
}
func (m *UseSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_UseSkill.DiscardUnknown(m)
}

var xxx_messageInfo_UseSkill proto.InternalMessageInfo

func (m *UseSkill) GetSkillId() string {
	if m != nil {
		return m.SkillId
	}
	return ""
}

func (m *UseSkill) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

type UseItem struct {
	ItemId               string   `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseItem) Reset()         { *m = UseItem{} }
func (m *UseItem) String() string { return proto.CompactTextString(m) }
func (*UseItem) ProtoMessage()    {}
func (*UseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{4}
}

func (m *UseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseItem.Unmarshal(m, b)
}
func (m *UseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseItem.Marshal(b, m, deterministic)
}
func (m *UseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseItem.Merge(m, src)
}
func (m *UseItem) XXX_Size() int {
	return xxx_messageInfo_UseItem.Size(m)
}
func (m *UseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UseItem.DiscardUnknown(m)
}

var xxx_messageInfo_UseItem proto.InternalMessageInfo

func (m *UseItem) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type Move struct {
	Dir                  Move_Dir `protobuf:"varint,1,opt,name=dir,proto3,enum=endless.stream.v1.Move_Dir" json:"dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Move) Reset()         { *m = Move{} }
func (m *Move) String() string { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()    {}
func (*Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{5}
}

func (m *Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Move.Unmarshal(m, b)
}
func (m *Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Move.Marshal(b, m, deterministic)
}
func (m *Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Move.Merge(m, src)
}
func (m *Move) XXX_Size() int {
	return xxx_messageInfo_Move.Size(m)
}
func (m *Move) XXX_DiscardUnknown() {
	xxx_messageInfo_Move.DiscardUnknown(m)
}

var xxx_messageInfo_Move proto.InternalMessageInfo

func (m *Move) GetDir() Move_Dir {
	if m != nil {
		return m.Dir
	}
	return Move_Left
}

type Input struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	// Types that are valid to be assigned to Input:
	//	*Input_Register
	//	*Input_CharSelect
	//	*Input_GameStart
	//	*Input_Skill
	//	*Input_Item
	//	*Input_Move
	Input                isInput_Input `protobuf_oneof:"input"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_db6f7669dced820e, []int{6}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Input) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type isInput_Input interface {
	isInput_Input()
}

type Input_Register struct {
	Register *Register `protobuf:"bytes,10,opt,name=register,proto3,oneof"`
}

type Input_CharSelect struct {
	CharSelect *CharSelect `protobuf:"bytes,20,opt,name=char_select,json=charSelect,proto3,oneof"`
}

type Input_GameStart struct {
	GameStart *GameStart `protobuf:"bytes,30,opt,name=game_start,json=gameStart,proto3,oneof"`
}

type Input_Skill struct {
	Skill *UseSkill `protobuf:"bytes,40,opt,name=skill,proto3,oneof"`
}

type Input_Item struct {
	Item *UseItem `protobuf:"bytes,50,opt,name=item,proto3,oneof"`
}

type Input_Move struct {
	Move *Move `protobuf:"bytes,60,opt,name=move,proto3,oneof"`
}

func (*Input_Register) isInput_Input() {}

func (*Input_CharSelect) isInput_Input() {}

func (*Input_GameStart) isInput_Input() {}

func (*Input_Skill) isInput_Input() {}

func (*Input_Item) isInput_Input() {}

func (*Input_Move) isInput_Input() {}

func (m *Input) GetInput() isInput_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Input) GetRegister() *Register {
	if x, ok := m.GetInput().(*Input_Register); ok {
		return x.Register
	}
	return nil
}

func (m *Input) GetCharSelect() *CharSelect {
	if x, ok := m.GetInput().(*Input_CharSelect); ok {
		return x.CharSelect
	}
	return nil
}

func (m *Input) GetGameStart() *GameStart {
	if x, ok := m.GetInput().(*Input_GameStart); ok {
		return x.GameStart
	}
	return nil
}

func (m *Input) GetSkill() *UseSkill {
	if x, ok := m.GetInput().(*Input_Skill); ok {
		return x.Skill
	}
	return nil
}

func (m *Input) GetItem() *UseItem {
	if x, ok := m.GetInput().(*Input_Item); ok {
		return x.Item
	}
	return nil
}

func (m *Input) GetMove() *Move {
	if x, ok := m.GetInput().(*Input_Move); ok {
		return x.Move
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Input) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Input_Register)(nil),
		(*Input_CharSelect)(nil),
		(*Input_GameStart)(nil),
		(*Input_Skill)(nil),
		(*Input_Item)(nil),
		(*Input_Move)(nil),
	}
}

func init() {
	proto.RegisterEnum("endless.stream.v1.Move_Dir", Move_Dir_name, Move_Dir_value)
	proto.RegisterType((*Register)(nil), "endless.stream.v1.Register")
	proto.RegisterType((*CharSelect)(nil), "endless.stream.v1.CharSelect")
	proto.RegisterType((*GameStart)(nil), "endless.stream.v1.GameStart")
	proto.RegisterType((*UseSkill)(nil), "endless.stream.v1.UseSkill")
	proto.RegisterType((*UseItem)(nil), "endless.stream.v1.UseItem")
	proto.RegisterType((*Move)(nil), "endless.stream.v1.Move")
	proto.RegisterType((*Input)(nil), "endless.stream.v1.Input")
}

func init() { proto.RegisterFile("input.proto", fileDescriptor_db6f7669dced820e) }

var fileDescriptor_db6f7669dced820e = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x9d, 0xac, 0x93, 0xd8, 0x13, 0x08, 0x5b, 0x51, 0x58, 0x37, 0xdb, 0x96, 0xc5, 0xa7,
	0x5c, 0xd6, 0x49, 0xbd, 0xa7, 0x42, 0x0b, 0x25, 0xbb, 0x50, 0x1b, 0xda, 0x43, 0x15, 0xf6, 0xd2,
	0x1e, 0x82, 0x62, 0x4d, 0x6d, 0xb1, 0xfe, 0x13, 0x24, 0x25, 0xd0, 0x4f, 0xd1, 0xaf, 0x5c, 0x24,
	0xe7, 0x0f, 0x6d, 0xbd, 0x37, 0x69, 0xe6, 0xbd, 0xd1, 0x63, 0x7e, 0x08, 0xc6, 0xa2, 0xde, 0xee,
	0x74, 0xb4, 0x95, 0x8d, 0x6e, 0xc8, 0x0b, 0xac, 0x79, 0x89, 0x4a, 0x45, 0x4a, 0x4b, 0x64, 0x55,
	0xb4, 0x7f, 0x37, 0x85, 0x9d, 0x16, 0x65, 0xdb, 0x0e, 0x63, 0xf0, 0x28, 0xe6, 0x42, 0x69, 0x94,
	0x84, 0x80, 0x9b, 0x35, 0x1c, 0x83, 0xde, 0x4d, 0x6f, 0xe6, 0x53, 0x7b, 0x36, 0xb5, 0x9a, 0x55,
	0x18, 0xf4, 0xdb, 0x9a, 0x39, 0x87, 0x3f, 0x00, 0xee, 0x0b, 0x26, 0x57, 0x58, 0x62, 0xa6, 0xc9,
	0x35, 0xf8, 0xdb, 0x92, 0xfd, 0x42, 0xb9, 0x16, 0xfc, 0x60, 0xf5, 0xda, 0x42, 0xca, 0xc9, 0x02,
	0x86, 0x59, 0xd1, 0x88, 0xac, 0x1d, 0x30, 0x89, 0x83, 0xe8, 0xbf, 0x38, 0xd1, 0x7d, 0xc9, 0x94,
	0xa2, 0x07, 0x5d, 0x38, 0x06, 0xff, 0x33, 0xab, 0x70, 0xa5, 0x99, 0xd4, 0xe1, 0x12, 0xbc, 0x47,
	0x85, 0xab, 0x27, 0x51, 0x96, 0xe4, 0x15, 0x78, 0xca, 0x1c, 0xce, 0xcf, 0x8c, 0xec, 0x3d, 0xe5,
	0x26, 0x82, 0x66, 0x32, 0x47, 0x6d, 0x7a, 0x6d, 0x52, 0xaf, 0x2d, 0xa4, 0x3c, 0x0c, 0x61, 0xf4,
	0xa8, 0x30, 0xd5, 0x58, 0x91, 0x2b, 0x18, 0x09, 0x8d, 0xd5, 0x79, 0xc2, 0xd0, 0x5c, 0x53, 0x1e,
	0x7e, 0x03, 0xf7, 0x6b, 0xb3, 0x47, 0x72, 0x0b, 0x17, 0x5c, 0x48, 0xdb, 0x9c, 0xc4, 0xd7, 0x1d,
	0x59, 0x8d, 0x2a, 0x7a, 0x10, 0x92, 0x1a, 0x5d, 0x38, 0x85, 0x8b, 0x07, 0x21, 0x89, 0x07, 0xee,
	0x17, 0xfc, 0xa9, 0x2f, 0x1d, 0xe2, 0xc3, 0x80, 0x8a, 0xbc, 0xd0, 0x97, 0xbd, 0xf0, 0xf7, 0x05,
	0x0c, 0x52, 0xc3, 0x81, 0x4c, 0xa0, 0x7f, 0x7a, 0xb0, 0x2f, 0xf8, 0xdf, 0x0b, 0xeb, 0xff, 0xb3,
	0xb0, 0xf7, 0xe0, 0xc9, 0x03, 0x8f, 0x00, 0x6e, 0x7a, 0xb3, 0x71, 0x67, 0x8c, 0x23, 0xb2, 0xc4,
	0xa1, 0x27, 0x39, 0xf9, 0x04, 0xe3, 0xac, 0x60, 0x72, 0xad, 0x2c, 0x97, 0xe0, 0xa5, 0x75, 0xbf,
	0xe9, 0x5a, 0xf8, 0x09, 0x5e, 0xe2, 0x50, 0xc8, 0xce, 0x28, 0x3f, 0x02, 0xe4, 0xac, 0xc2, 0xb5,
	0x32, 0xcb, 0x0f, 0xde, 0xda, 0x01, 0xaf, 0x3b, 0x06, 0x9c, 0x00, 0x25, 0x0e, 0xf5, 0xf3, 0xe3,
	0x85, 0xdc, 0xc1, 0xc0, 0x12, 0x09, 0x66, 0xcf, 0x06, 0x3f, 0xd2, 0x4c, 0x1c, 0xda, 0x6a, 0xc9,
	0x02, 0x5c, 0x03, 0x21, 0x88, 0xad, 0x67, 0xda, 0xed, 0x31, 0xf4, 0x12, 0x87, 0x5a, 0x25, 0xb9,
	0x05, 0xb7, 0x6a, 0xf6, 0x18, 0x7c, 0xb0, 0x8e, 0xab, 0x67, 0x28, 0x19, 0xb9, 0x91, 0x2d, 0x47,
	0x30, 0xb0, 0xff, 0x61, 0x19, 0x7f, 0x5f, 0xe4, 0x42, 0x17, 0xbb, 0x4d, 0x94, 0x35, 0xd5, 0x5c,
	0x21, 0xab, 0x0b, 0x96, 0x63, 0x3d, 0x3f, 0xf8, 0xd7, 0xad, 0x7f, 0xbe, 0x61, 0xd9, 0x13, 0xd6,
	0x7c, 0x9e, 0xcb, 0x6d, 0xb6, 0x19, 0xda, 0x5f, 0x72, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x94,
	0x67, 0x78, 0xc9, 0x53, 0x03, 0x00, 0x00,
}