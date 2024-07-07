// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/partyaffairs/partyaffairs_banner.proto

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

type Banner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Src       string       `protobuf:"bytes,3,opt,name=src,proto3" json:"src,omitempty"`
	Path      string       `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Weight    uint32       `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	CreatedAt uint32       `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint32       `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Resource  *Banner_File `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Banner) Reset() {
	*x = Banner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Banner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Banner) ProtoMessage() {}

func (x *Banner) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Banner.ProtoReflect.Descriptor instead.
func (*Banner) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP(), []int{0}
}

func (x *Banner) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Banner) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Banner) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *Banner) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Banner) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Banner) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Banner) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Banner) GetResource() *Banner_File {
	if x != nil {
		return x.Resource
	}
	return nil
}

type AllBannerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Banner `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AllBannerReply) Reset() {
	*x = AllBannerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllBannerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllBannerReply) ProtoMessage() {}

func (x *AllBannerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllBannerReply.ProtoReflect.Descriptor instead.
func (*AllBannerReply) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP(), []int{1}
}

func (x *AllBannerReply) GetList() []*Banner {
	if x != nil {
		return x.List
	}
	return nil
}

type AddBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Src   string  `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	Path  *string `protobuf:"bytes,3,opt,name=path,proto3,oneof" json:"path,omitempty"`
}

func (x *AddBannerRequest) Reset() {
	*x = AddBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBannerRequest) ProtoMessage() {}

func (x *AddBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBannerRequest.ProtoReflect.Descriptor instead.
func (*AddBannerRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP(), []int{2}
}

func (x *AddBannerRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddBannerRequest) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *AddBannerRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

type UpdateBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Src   string  `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	Path  *string `protobuf:"bytes,3,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Id    uint32  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateBannerRequest) Reset() {
	*x = UpdateBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBannerRequest) ProtoMessage() {}

func (x *UpdateBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBannerRequest.ProtoReflect.Descriptor instead.
func (*UpdateBannerRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBannerRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateBannerRequest) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *UpdateBannerRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *UpdateBannerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBannerRequest) Reset() {
	*x = DeleteBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBannerRequest) ProtoMessage() {}

func (x *DeleteBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBannerRequest.ProtoReflect.Descriptor instead.
func (*DeleteBannerRequest) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBannerRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Banner_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sha  string `protobuf:"bytes,2,opt,name=sha,proto3" json:"sha,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Banner_File) Reset() {
	*x = Banner_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Banner_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Banner_File) ProtoMessage() {}

func (x *Banner_File) ProtoReflect() protoreflect.Message {
	mi := &file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Banner_File.ProtoReflect.Descriptor instead.
func (*Banner_File) Descriptor() ([]byte, []int) {
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Banner_File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Banner_File) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *Banner_File) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_api_partyaffairs_partyaffairs_banner_proto protoreflect.FileDescriptor

var file_api_partyaffairs_partyaffairs_banner_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69,
	0x72, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66,
	0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a,
	0x3e, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x4e, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x3c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x6e, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x17, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x8a, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x73, 0x72,
	0x63, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_partyaffairs_partyaffairs_banner_proto_rawDescOnce sync.Once
	file_api_partyaffairs_partyaffairs_banner_proto_rawDescData = file_api_partyaffairs_partyaffairs_banner_proto_rawDesc
)

func file_api_partyaffairs_partyaffairs_banner_proto_rawDescGZIP() []byte {
	file_api_partyaffairs_partyaffairs_banner_proto_rawDescOnce.Do(func() {
		file_api_partyaffairs_partyaffairs_banner_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_partyaffairs_partyaffairs_banner_proto_rawDescData)
	})
	return file_api_partyaffairs_partyaffairs_banner_proto_rawDescData
}

var file_api_partyaffairs_partyaffairs_banner_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_partyaffairs_partyaffairs_banner_proto_goTypes = []interface{}{
	(*Banner)(nil),              // 0: partyaffairs.api.partyaffairs.v1.Banner
	(*AllBannerReply)(nil),      // 1: partyaffairs.api.partyaffairs.v1.AllBannerReply
	(*AddBannerRequest)(nil),    // 2: partyaffairs.api.partyaffairs.v1.AddBannerRequest
	(*UpdateBannerRequest)(nil), // 3: partyaffairs.api.partyaffairs.v1.UpdateBannerRequest
	(*DeleteBannerRequest)(nil), // 4: partyaffairs.api.partyaffairs.v1.DeleteBannerRequest
	(*Banner_File)(nil),         // 5: partyaffairs.api.partyaffairs.v1.Banner.File
}
var file_api_partyaffairs_partyaffairs_banner_proto_depIdxs = []int32{
	5, // 0: partyaffairs.api.partyaffairs.v1.Banner.resource:type_name -> partyaffairs.api.partyaffairs.v1.Banner.File
	0, // 1: partyaffairs.api.partyaffairs.v1.AllBannerReply.list:type_name -> partyaffairs.api.partyaffairs.v1.Banner
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_partyaffairs_partyaffairs_banner_proto_init() }
func file_api_partyaffairs_partyaffairs_banner_proto_init() {
	if File_api_partyaffairs_partyaffairs_banner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Banner); i {
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
		file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllBannerReply); i {
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
		file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBannerRequest); i {
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
		file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBannerRequest); i {
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
		file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBannerRequest); i {
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
		file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Banner_File); i {
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
	file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_api_partyaffairs_partyaffairs_banner_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_partyaffairs_partyaffairs_banner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_partyaffairs_partyaffairs_banner_proto_goTypes,
		DependencyIndexes: file_api_partyaffairs_partyaffairs_banner_proto_depIdxs,
		MessageInfos:      file_api_partyaffairs_partyaffairs_banner_proto_msgTypes,
	}.Build()
	File_api_partyaffairs_partyaffairs_banner_proto = out.File
	file_api_partyaffairs_partyaffairs_banner_proto_rawDesc = nil
	file_api_partyaffairs_partyaffairs_banner_proto_goTypes = nil
	file_api_partyaffairs_partyaffairs_banner_proto_depIdxs = nil
}
