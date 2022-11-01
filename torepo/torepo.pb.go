// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: torepo/torepo.proto

package torepo

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tkn       string `protobuf:"bytes,1,opt,name=tkn,proto3" json:"tkn,omitempty"`
	Id        int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Inn       string `protobuf:"bytes,3,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp       string `protobuf:"bytes,4,opt,name=kpp,proto3" json:"kpp,omitempty"`
	ShortName string `protobuf:"bytes,5,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	FullName  string `protobuf:"bytes,6,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetTkn() string {
	if x != nil {
		return x.Tkn
	}
	return ""
}

func (x *Entity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Entity) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *Entity) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *Entity) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Entity) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

type NewEntityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *NewEntityReply) Reset() {
	*x = NewEntityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEntityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEntityReply) ProtoMessage() {}

func (x *NewEntityReply) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEntityReply.ProtoReflect.Descriptor instead.
func (*NewEntityReply) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{1}
}

func (x *NewEntityReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewEntityReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tkn         string `protobuf:"bytes,1,opt,name=tkn,proto3" json:"tkn,omitempty"`
	Id          int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{2}
}

func (x *Origin) GetTkn() string {
	if x != nil {
		return x.Tkn
	}
	return ""
}

func (x *Origin) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Origin) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type NewOriginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *NewOriginReply) Reset() {
	*x = NewOriginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOriginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOriginReply) ProtoMessage() {}

func (x *NewOriginReply) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOriginReply.ProtoReflect.Descriptor instead.
func (*NewOriginReply) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{3}
}

func (x *NewOriginReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewOriginReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Obligation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tkn        string               `protobuf:"bytes,1,opt,name=tkn,proto3" json:"tkn,omitempty"`
	Id         int32                `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Dbtorid    int32                `protobuf:"varint,3,opt,name=dbtorid,proto3" json:"dbtorid,omitempty"`
	Creditorid int32                `protobuf:"varint,4,opt,name=creditorid,proto3" json:"creditorid,omitempty"`
	Cost       float32              `protobuf:"fixed32,5,opt,name=cost,proto3" json:"cost,omitempty"`
	Originid   int32                `protobuf:"varint,6,opt,name=originid,proto3" json:"originid,omitempty"`
	Date       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Obligation) Reset() {
	*x = Obligation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Obligation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Obligation) ProtoMessage() {}

func (x *Obligation) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Obligation.ProtoReflect.Descriptor instead.
func (*Obligation) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{4}
}

func (x *Obligation) GetTkn() string {
	if x != nil {
		return x.Tkn
	}
	return ""
}

func (x *Obligation) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Obligation) GetDbtorid() int32 {
	if x != nil {
		return x.Dbtorid
	}
	return 0
}

func (x *Obligation) GetCreditorid() int32 {
	if x != nil {
		return x.Creditorid
	}
	return 0
}

func (x *Obligation) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Obligation) GetOriginid() int32 {
	if x != nil {
		return x.Originid
	}
	return 0
}

func (x *Obligation) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type NewObligationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *NewObligationReply) Reset() {
	*x = NewObligationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewObligationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewObligationReply) ProtoMessage() {}

func (x *NewObligationReply) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewObligationReply.ProtoReflect.Descriptor instead.
func (*NewObligationReply) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{5}
}

func (x *NewObligationReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewObligationReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Tkn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tkn string `protobuf:"bytes,1,opt,name=tkn,proto3" json:"tkn,omitempty"`
}

func (x *Tkn) Reset() {
	*x = Tkn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torepo_torepo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tkn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tkn) ProtoMessage() {}

func (x *Tkn) ProtoReflect() protoreflect.Message {
	mi := &file_torepo_torepo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tkn.ProtoReflect.Descriptor instead.
func (*Tkn) Descriptor() ([]byte, []int) {
	return file_torepo_torepo_proto_rawDescGZIP(), []int{6}
}

func (x *Tkn) GetTkn() string {
	if x != nil {
		return x.Tkn
	}
	return ""
}

var File_torepo_torepo_proto protoreflect.FileDescriptor

var file_torepo_torepo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6b, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x6b, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x6e, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x70, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x70, 0x70, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0e, 0x4e,
	0x65, 0x77, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x4c, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6b, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x6b, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a,
	0x0e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x6b, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x6b, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x12,
	0x4e, 0x65, 0x77, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x22, 0x17, 0x0a, 0x03, 0x54, 0x6b, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x6b, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x6b, 0x6e, 0x32, 0xa5, 0x02,
	0x0a, 0x0c, 0x42, 0x32, 0x62, 0x32, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35,
	0x0a, 0x09, 0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x2e, 0x74, 0x6f,
	0x72, 0x65, 0x70, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x74, 0x6f,
	0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x0b, 0x2e, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x54, 0x6b, 0x6e,
	0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x0e, 0x2e, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x4e,
	0x65, 0x77, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x74,
	0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x62, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0b, 0x2e, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x54, 0x6b, 0x6e, 0x1a, 0x12, 0x2e,
	0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x4f, 0x62, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x32, 0x62, 0x32, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x2f, 0x6c, 0x69,
	0x62, 0x2f, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_torepo_torepo_proto_rawDescOnce sync.Once
	file_torepo_torepo_proto_rawDescData = file_torepo_torepo_proto_rawDesc
)

func file_torepo_torepo_proto_rawDescGZIP() []byte {
	file_torepo_torepo_proto_rawDescOnce.Do(func() {
		file_torepo_torepo_proto_rawDescData = protoimpl.X.CompressGZIP(file_torepo_torepo_proto_rawDescData)
	})
	return file_torepo_torepo_proto_rawDescData
}

var file_torepo_torepo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_torepo_torepo_proto_goTypes = []interface{}{
	(*Entity)(nil),              // 0: torepo.Entity
	(*NewEntityReply)(nil),      // 1: torepo.NewEntityReply
	(*Origin)(nil),              // 2: torepo.Origin
	(*NewOriginReply)(nil),      // 3: torepo.NewOriginReply
	(*Obligation)(nil),          // 4: torepo.Obligation
	(*NewObligationReply)(nil),  // 5: torepo.NewObligationReply
	(*Tkn)(nil),                 // 6: torepo.Tkn
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_torepo_torepo_proto_depIdxs = []int32{
	7, // 0: torepo.Obligation.date:type_name -> google.protobuf.Timestamp
	0, // 1: torepo.B2b2bService.NewEntity:input_type -> torepo.Entity
	6, // 2: torepo.B2b2bService.ListEntity:input_type -> torepo.Tkn
	2, // 3: torepo.B2b2bService.NewOrigin:input_type -> torepo.Origin
	4, // 4: torepo.B2b2bService.NewObligation:input_type -> torepo.Obligation
	6, // 5: torepo.B2b2bService.ListObligation:input_type -> torepo.Tkn
	1, // 6: torepo.B2b2bService.NewEntity:output_type -> torepo.NewEntityReply
	0, // 7: torepo.B2b2bService.ListEntity:output_type -> torepo.Entity
	3, // 8: torepo.B2b2bService.NewOrigin:output_type -> torepo.NewOriginReply
	5, // 9: torepo.B2b2bService.NewObligation:output_type -> torepo.NewObligationReply
	4, // 10: torepo.B2b2bService.ListObligation:output_type -> torepo.Obligation
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_torepo_torepo_proto_init() }
func file_torepo_torepo_proto_init() {
	if File_torepo_torepo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_torepo_torepo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_torepo_torepo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEntityReply); i {
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
		file_torepo_torepo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
		file_torepo_torepo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOriginReply); i {
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
		file_torepo_torepo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Obligation); i {
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
		file_torepo_torepo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewObligationReply); i {
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
		file_torepo_torepo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tkn); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_torepo_torepo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_torepo_torepo_proto_goTypes,
		DependencyIndexes: file_torepo_torepo_proto_depIdxs,
		MessageInfos:      file_torepo_torepo_proto_msgTypes,
	}.Build()
	File_torepo_torepo_proto = out.File
	file_torepo_torepo_proto_rawDesc = nil
	file_torepo_torepo_proto_goTypes = nil
	file_torepo_torepo_proto_depIdxs = nil
}
