// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: tests.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_ENUM_UNSPECIFIED Enum = 0
	Enum_ENUM_ONE         Enum = 1
	Enum_ENUM_TWO         Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_ONE",
		2: "ENUM_TWO",
	}
	Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_ONE":         1,
		"ENUM_TWO":         2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_tests_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{0}
}

type Singulars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bool     bool       `protobuf:"varint,1,opt,name=bool,proto3" json:"bool,omitempty"`
	Float    float32    `protobuf:"fixed32,2,opt,name=float,proto3" json:"float,omitempty"`
	Double   float64    `protobuf:"fixed64,3,opt,name=double,proto3" json:"double,omitempty"`
	Bytes    []byte     `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	String_  string     `protobuf:"bytes,5,opt,name=string,proto3" json:"string,omitempty"`
	Enum     Enum       `protobuf:"varint,6,opt,name=enum,proto3,enum=Enum" json:"enum,omitempty"`
	Int32    int32      `protobuf:"varint,7,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64    int64      `protobuf:"varint,8,opt,name=int64,proto3" json:"int64,omitempty"`
	Sint32   int32      `protobuf:"zigzag32,9,opt,name=sint32,proto3" json:"sint32,omitempty"`
	Sint64   int64      `protobuf:"zigzag64,10,opt,name=sint64,proto3" json:"sint64,omitempty"`
	Sfixed32 int32      `protobuf:"fixed32,11,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64 int64      `protobuf:"fixed64,12,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Uint32   uint32     `protobuf:"varint,13,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64   uint64     `protobuf:"varint,14,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Fixed32  uint32     `protobuf:"fixed32,15,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64  uint64     `protobuf:"fixed64,16,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Message  *Singulars `protobuf:"bytes,17,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Singulars) Reset() {
	*x = Singulars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Singulars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Singulars) ProtoMessage() {}

func (x *Singulars) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Singulars.ProtoReflect.Descriptor instead.
func (*Singulars) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{0}
}

func (x *Singulars) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Singulars) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Singulars) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Singulars) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Singulars) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Singulars) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_ENUM_UNSPECIFIED
}

func (x *Singulars) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Singulars) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Singulars) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Singulars) GetSint64() int64 {
	if x != nil {
		return x.Sint64
	}
	return 0
}

func (x *Singulars) GetSfixed32() int32 {
	if x != nil {
		return x.Sfixed32
	}
	return 0
}

func (x *Singulars) GetSfixed64() int64 {
	if x != nil {
		return x.Sfixed64
	}
	return 0
}

func (x *Singulars) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Singulars) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Singulars) GetFixed32() uint32 {
	if x != nil {
		return x.Fixed32
	}
	return 0
}

func (x *Singulars) GetFixed64() uint64 {
	if x != nil {
		return x.Fixed64
	}
	return 0
}

func (x *Singulars) GetMessage() *Singulars {
	if x != nil {
		return x.Message
	}
	return nil
}

type Oneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to O:
	//
	//	*Oneof_Bool
	//	*Oneof_Int32
	O isOneof_O `protobuf_oneof:"o"`
}

func (x *Oneof) Reset() {
	*x = Oneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oneof) ProtoMessage() {}

func (x *Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oneof.ProtoReflect.Descriptor instead.
func (*Oneof) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{1}
}

func (m *Oneof) GetO() isOneof_O {
	if m != nil {
		return m.O
	}
	return nil
}

func (x *Oneof) GetBool() bool {
	if x, ok := x.GetO().(*Oneof_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Oneof) GetInt32() int32 {
	if x, ok := x.GetO().(*Oneof_Int32); ok {
		return x.Int32
	}
	return 0
}

type isOneof_O interface {
	isOneof_O()
}

type Oneof_Bool struct {
	Bool bool `protobuf:"varint,1,opt,name=bool,proto3,oneof"`
}

type Oneof_Int32 struct {
	Int32 int32 `protobuf:"varint,2,opt,name=int32,proto3,oneof"`
}

func (*Oneof_Bool) isOneof_O() {}

func (*Oneof_Int32) isOneof_O() {}

type Lists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Singulars []int32      `protobuf:"varint,1,rep,packed,name=singulars,proto3" json:"singulars,omitempty"`
	Messages  []*Singulars `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Lists) Reset() {
	*x = Lists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lists) ProtoMessage() {}

func (x *Lists) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lists.ProtoReflect.Descriptor instead.
func (*Lists) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{2}
}

func (x *Lists) GetSingulars() []int32 {
	if x != nil {
		return x.Singulars
	}
	return nil
}

func (x *Lists) GetMessages() []*Singulars {
	if x != nil {
		return x.Messages
	}
	return nil
}

type Maps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bools   map[bool]string   `protobuf:"bytes,1,rep,name=bools,proto3" json:"bools,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ints    map[int32]string  `protobuf:"bytes,2,rep,name=ints,proto3" json:"ints,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uints   map[uint32]string `protobuf:"bytes,3,rep,name=uints,proto3" json:"uints,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Strings map[string]string `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Maps) Reset() {
	*x = Maps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maps) ProtoMessage() {}

func (x *Maps) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maps.ProtoReflect.Descriptor instead.
func (*Maps) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{3}
}

func (x *Maps) GetBools() map[bool]string {
	if x != nil {
		return x.Bools
	}
	return nil
}

func (x *Maps) GetInts() map[int32]string {
	if x != nil {
		return x.Ints
	}
	return nil
}

func (x *Maps) GetUints() map[uint32]string {
	if x != nil {
		return x.Uints
	}
	return nil
}

func (x *Maps) GetStrings() map[string]string {
	if x != nil {
		return x.Strings
	}
	return nil
}

type Redaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Redaction) Reset() {
	*x = Redaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redaction) ProtoMessage() {}

func (x *Redaction) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redaction.ProtoReflect.Descriptor instead.
func (*Redaction) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{4}
}

func (x *Redaction) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type OuterNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OuterNested) Reset() {
	*x = OuterNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterNested) ProtoMessage() {}

func (x *OuterNested) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterNested.ProtoReflect.Descriptor instead.
func (*OuterNested) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{5}
}

type OuterNested_InnerNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OuterNested_InnerNested) Reset() {
	*x = OuterNested_InnerNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterNested_InnerNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterNested_InnerNested) ProtoMessage() {}

func (x *OuterNested_InnerNested) ProtoReflect() protoreflect.Message {
	mi := &file_tests_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterNested_InnerNested.ProtoReflect.Descriptor instead.
func (*OuterNested_InnerNested) Descriptor() ([]byte, []int) {
	return file_tests_proto_rawDescGZIP(), []int{5, 0}
}

var File_tests_proto protoreflect.FileDescriptor

var file_tests_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x03,
	0x0a, 0x09, 0x53, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x04, 0x65,
	0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10,
	0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x24,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x05, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x14, 0x0a,
	0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x03, 0x0a, 0x01, 0x6f,
	0x22, 0x4d, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x6e,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x73, 0x69,
	0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x6e, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x92, 0x03, 0x0a, 0x04, 0x4d, 0x61, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x73,
	0x12, 0x23, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x75, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x75, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x42,
	0x6f, 0x6f, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38,
	0x0a, 0x0a, 0x55, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x09, 0x52, 0x65, 0x64, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x15, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0x80, 0x01, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1c, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x65,
	0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2a, 0x38, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54, 0x57, 0x4f, 0x10, 0x02,
	0x42, 0x39, 0x42, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x64,
	0x61, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tests_proto_rawDescOnce sync.Once
	file_tests_proto_rawDescData = file_tests_proto_rawDesc
)

func file_tests_proto_rawDescGZIP() []byte {
	file_tests_proto_rawDescOnce.Do(func() {
		file_tests_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_proto_rawDescData)
	})
	return file_tests_proto_rawDescData
}

var file_tests_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tests_proto_goTypes = []interface{}{
	(Enum)(0),                       // 0: Enum
	(*Singulars)(nil),               // 1: Singulars
	(*Oneof)(nil),                   // 2: Oneof
	(*Lists)(nil),                   // 3: Lists
	(*Maps)(nil),                    // 4: Maps
	(*Redaction)(nil),               // 5: Redaction
	(*OuterNested)(nil),             // 6: OuterNested
	nil,                             // 7: Maps.BoolsEntry
	nil,                             // 8: Maps.IntsEntry
	nil,                             // 9: Maps.UintsEntry
	nil,                             // 10: Maps.StringsEntry
	(*OuterNested_InnerNested)(nil), // 11: OuterNested.InnerNested
}
var file_tests_proto_depIdxs = []int32{
	0,  // 0: Singulars.enum:type_name -> Enum
	1,  // 1: Singulars.message:type_name -> Singulars
	1,  // 2: Lists.messages:type_name -> Singulars
	7,  // 3: Maps.bools:type_name -> Maps.BoolsEntry
	8,  // 4: Maps.ints:type_name -> Maps.IntsEntry
	9,  // 5: Maps.uints:type_name -> Maps.UintsEntry
	10, // 6: Maps.strings:type_name -> Maps.StringsEntry
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_tests_proto_init() }
func file_tests_proto_init() {
	if File_tests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Singulars); i {
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
		file_tests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oneof); i {
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
		file_tests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lists); i {
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
		file_tests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maps); i {
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
		file_tests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redaction); i {
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
		file_tests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterNested); i {
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
		file_tests_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterNested_InnerNested); i {
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
	file_tests_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Oneof_Bool)(nil),
		(*Oneof_Int32)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_proto_goTypes,
		DependencyIndexes: file_tests_proto_depIdxs,
		EnumInfos:         file_tests_proto_enumTypes,
		MessageInfos:      file_tests_proto_msgTypes,
	}.Build()
	File_tests_proto = out.File
	file_tests_proto_rawDesc = nil
	file_tests_proto_goTypes = nil
	file_tests_proto_depIdxs = nil
}
