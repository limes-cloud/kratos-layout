// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/partyaffairs/information/partyaffairs_information_classify.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ListInformationClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListInformationClassifyRequest) Reset() {
	*x = ListInformationClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInformationClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInformationClassifyRequest) ProtoMessage() {}

func (x *ListInformationClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInformationClassifyRequest.ProtoReflect.Descriptor instead.
func (*ListInformationClassifyRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{0}
}

type ListInformationClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ListInformationClassifyReply_InformationClassify `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListInformationClassifyReply) Reset() {
	*x = ListInformationClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInformationClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInformationClassifyReply) ProtoMessage() {}

func (x *ListInformationClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInformationClassifyReply.ProtoReflect.Descriptor instead.
func (*ListInformationClassifyReply) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{1}
}

func (x *ListInformationClassifyReply) GetList() []*ListInformationClassifyReply_InformationClassify {
	if x != nil {
		return x.List
	}
	return nil
}

type CreateInformationClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *CreateInformationClassifyRequest) Reset() {
	*x = CreateInformationClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInformationClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInformationClassifyRequest) ProtoMessage() {}

func (x *CreateInformationClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInformationClassifyRequest.ProtoReflect.Descriptor instead.
func (*CreateInformationClassifyRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{2}
}

func (x *CreateInformationClassifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateInformationClassifyRequest) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type CreateInformationClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateInformationClassifyReply) Reset() {
	*x = CreateInformationClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInformationClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInformationClassifyReply) ProtoMessage() {}

func (x *CreateInformationClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInformationClassifyReply.ProtoReflect.Descriptor instead.
func (*CreateInformationClassifyReply) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{3}
}

func (x *CreateInformationClassifyReply) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateInformationClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *UpdateInformationClassifyRequest) Reset() {
	*x = UpdateInformationClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInformationClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInformationClassifyRequest) ProtoMessage() {}

func (x *UpdateInformationClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInformationClassifyRequest.ProtoReflect.Descriptor instead.
func (*UpdateInformationClassifyRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInformationClassifyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateInformationClassifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateInformationClassifyRequest) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type UpdateInformationClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInformationClassifyReply) Reset() {
	*x = UpdateInformationClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInformationClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInformationClassifyReply) ProtoMessage() {}

func (x *UpdateInformationClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInformationClassifyReply.ProtoReflect.Descriptor instead.
func (*UpdateInformationClassifyReply) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{5}
}

type DeleteInformationClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInformationClassifyRequest) Reset() {
	*x = DeleteInformationClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInformationClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInformationClassifyRequest) ProtoMessage() {}

func (x *DeleteInformationClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInformationClassifyRequest.ProtoReflect.Descriptor instead.
func (*DeleteInformationClassifyRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteInformationClassifyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteInformationClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInformationClassifyReply) Reset() {
	*x = DeleteInformationClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInformationClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInformationClassifyReply) ProtoMessage() {}

func (x *DeleteInformationClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInformationClassifyReply.ProtoReflect.Descriptor instead.
func (*DeleteInformationClassifyReply) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{7}
}

type ListInformationClassifyReply_InformationClassify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Weight    uint32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	CreatedAt uint32 `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt uint32 `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ListInformationClassifyReply_InformationClassify) Reset() {
	*x = ListInformationClassifyReply_InformationClassify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInformationClassifyReply_InformationClassify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInformationClassifyReply_InformationClassify) ProtoMessage() {}

func (x *ListInformationClassifyReply_InformationClassify) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInformationClassifyReply_InformationClassify.ProtoReflect.Descriptor instead.
func (*ListInformationClassifyReply_InformationClassify) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListInformationClassifyReply_InformationClassify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListInformationClassifyReply_InformationClassify) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListInformationClassifyReply_InformationClassify) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ListInformationClassifyReply_InformationClassify) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListInformationClassifyReply_InformationClassify) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_api_partyaffairs_information_partyaffairs_information_classify_proto protoreflect.FileDescriptor

var file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDesc = []byte{
	0x0a, 0x44, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69,
	0x72, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66,
	0x61, 0x69, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66,
	0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a,
	0x1e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xa2, 0x02, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x72, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5e,
	0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x1a, 0x8d, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x57, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x30, 0x0a,
	0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x70, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x22, 0x20, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x3b, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x20, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x49, 0x0a, 0x30, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69,
	0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescOnce sync.Once
	file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescData = file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDesc
)

func file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescGZIP() []byte {
	file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescOnce.Do(func() {
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescData)
	})
	return file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDescData
}

var file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_partyaffairs_information_partyaffairs_information_classify_proto_goTypes = []interface{}{
	(*ListInformationClassifyRequest)(nil),                   // 0: partyaffairs.api.partyaffairs.information.v1.ListInformationClassifyRequest
	(*ListInformationClassifyReply)(nil),                     // 1: partyaffairs.api.partyaffairs.information.v1.ListInformationClassifyReply
	(*CreateInformationClassifyRequest)(nil),                 // 2: partyaffairs.api.partyaffairs.information.v1.CreateInformationClassifyRequest
	(*CreateInformationClassifyReply)(nil),                   // 3: partyaffairs.api.partyaffairs.information.v1.CreateInformationClassifyReply
	(*UpdateInformationClassifyRequest)(nil),                 // 4: partyaffairs.api.partyaffairs.information.v1.UpdateInformationClassifyRequest
	(*UpdateInformationClassifyReply)(nil),                   // 5: partyaffairs.api.partyaffairs.information.v1.UpdateInformationClassifyReply
	(*DeleteInformationClassifyRequest)(nil),                 // 6: partyaffairs.api.partyaffairs.information.v1.DeleteInformationClassifyRequest
	(*DeleteInformationClassifyReply)(nil),                   // 7: partyaffairs.api.partyaffairs.information.v1.DeleteInformationClassifyReply
	(*ListInformationClassifyReply_InformationClassify)(nil), // 8: partyaffairs.api.partyaffairs.information.v1.ListInformationClassifyReply.InformationClassify
}
var file_api_partyaffairs_information_partyaffairs_information_classify_proto_depIdxs = []int32{
	8, // 0: partyaffairs.api.partyaffairs.information.v1.ListInformationClassifyReply.list:type_name -> partyaffairs.api.partyaffairs.information.v1.ListInformationClassifyReply.InformationClassify
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_partyaffairs_information_partyaffairs_information_classify_proto_init() }
func file_api_partyaffairs_information_partyaffairs_information_classify_proto_init() {
	if File_api_partyaffairs_information_partyaffairs_information_classify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInformationClassifyRequest); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInformationClassifyReply); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInformationClassifyRequest); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInformationClassifyReply); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInformationClassifyRequest); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInformationClassifyReply); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInformationClassifyRequest); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInformationClassifyReply); i {
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
		file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInformationClassifyReply_InformationClassify); i {
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
			RawDescriptor: file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_partyaffairs_information_partyaffairs_information_classify_proto_goTypes,
		DependencyIndexes: file_api_partyaffairs_information_partyaffairs_information_classify_proto_depIdxs,
		MessageInfos:      file_api_partyaffairs_information_partyaffairs_information_classify_proto_msgTypes,
	}.Build()
	File_api_partyaffairs_information_partyaffairs_information_classify_proto = out.File
	file_api_partyaffairs_information_partyaffairs_information_classify_proto_rawDesc = nil
	file_api_partyaffairs_information_partyaffairs_information_classify_proto_goTypes = nil
	file_api_partyaffairs_information_partyaffairs_information_classify_proto_depIdxs = nil
}
