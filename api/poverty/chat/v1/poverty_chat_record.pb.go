// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/poverty/chat/poverty_chat_record.proto

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

type SendChatMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	SessionId *string `protobuf:"bytes,2,opt,name=sessionId,proto3,oneof" json:"sessionId,omitempty"`
}

func (x *SendChatMessageRequest) Reset() {
	*x = SendChatMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatMessageRequest) ProtoMessage() {}

func (x *SendChatMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatMessageRequest.ProtoReflect.Descriptor instead.
func (*SendChatMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{0}
}

func (x *SendChatMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendChatMessageRequest) GetSessionId() string {
	if x != nil && x.SessionId != nil {
		return *x.SessionId
	}
	return ""
}

type SendChatMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *SendChatMessageReply) Reset() {
	*x = SendChatMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChatMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChatMessageReply) ProtoMessage() {}

func (x *SendChatMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChatMessageReply.ProtoReflect.Descriptor instead.
func (*SendChatMessageReply) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{1}
}

func (x *SendChatMessageReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendChatMessageReply) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type ListChatRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  uint32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Order     *string `protobuf:"bytes,3,opt,name=order,proto3,oneof" json:"order,omitempty"`
	OrderBy   *string `protobuf:"bytes,4,opt,name=orderBy,proto3,oneof" json:"orderBy,omitempty"`
	UserId    *uint32 `protobuf:"varint,5,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
	SessionId *string `protobuf:"bytes,6,opt,name=sessionId,proto3,oneof" json:"sessionId,omitempty"`
	Distinct  *bool   `protobuf:"varint,7,opt,name=distinct,proto3,oneof" json:"distinct,omitempty"`
	UserName  *string `protobuf:"bytes,8,opt,name=userName,proto3,oneof" json:"userName,omitempty"`
}

func (x *ListChatRecordRequest) Reset() {
	*x = ListChatRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatRecordRequest) ProtoMessage() {}

func (x *ListChatRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatRecordRequest.ProtoReflect.Descriptor instead.
func (*ListChatRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{2}
}

func (x *ListChatRecordRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListChatRecordRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListChatRecordRequest) GetOrder() string {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return ""
}

func (x *ListChatRecordRequest) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *ListChatRecordRequest) GetUserId() uint32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *ListChatRecordRequest) GetSessionId() string {
	if x != nil && x.SessionId != nil {
		return *x.SessionId
	}
	return ""
}

func (x *ListChatRecordRequest) GetDistinct() bool {
	if x != nil && x.Distinct != nil {
		return *x.Distinct
	}
	return false
}

func (x *ListChatRecordRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

type ListChatRecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32                            `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*ListChatRecordReply_ChatRecord `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListChatRecordReply) Reset() {
	*x = ListChatRecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatRecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatRecordReply) ProtoMessage() {}

func (x *ListChatRecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatRecordReply.ProtoReflect.Descriptor instead.
func (*ListChatRecordReply) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{3}
}

func (x *ListChatRecordReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListChatRecordReply) GetList() []*ListChatRecordReply_ChatRecord {
	if x != nil {
		return x.List
	}
	return nil
}

type DeleteChatRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteChatRecordRequest) Reset() {
	*x = DeleteChatRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRecordRequest) ProtoMessage() {}

func (x *DeleteChatRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteChatRecordRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteChatRecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteChatRecordReply) Reset() {
	*x = DeleteChatRecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRecordReply) ProtoMessage() {}

func (x *DeleteChatRecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRecordReply.ProtoReflect.Descriptor instead.
func (*DeleteChatRecordReply) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{5}
}

type ListChatRecordReply_ChatRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     uint32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName   string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	UserAvatar string `protobuf:"bytes,4,opt,name=userAvatar,proto3" json:"userAvatar,omitempty"`
	SessionId  string `protobuf:"bytes,5,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Message    string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Type       string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt  uint32 `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ListChatRecordReply_ChatRecord) Reset() {
	*x = ListChatRecordReply_ChatRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatRecordReply_ChatRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatRecordReply_ChatRecord) ProtoMessage() {}

func (x *ListChatRecordReply_ChatRecord) ProtoReflect() protoreflect.Message {
	mi := &file_api_poverty_chat_poverty_chat_record_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatRecordReply_ChatRecord.ProtoReflect.Descriptor instead.
func (*ListChatRecordReply_ChatRecord) Descriptor() ([]byte, []int) {
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListChatRecordReply_ChatRecord) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListChatRecordReply_ChatRecord) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListChatRecordReply_ChatRecord) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ListChatRecordReply_ChatRecord) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *ListChatRecordReply_ChatRecord) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ListChatRecordReply_ChatRecord) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListChatRecordReply_ChatRecord) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ListChatRecordReply_ChatRecord) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_api_poverty_chat_poverty_chat_record_proto protoreflect.FileDescriptor

var file_api_poverty_chat_poverty_chat_record_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x6f,
	0x76, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x74,
	0x79, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xac, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xa6, 0x03, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x28, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42,
	0x06, 0x2a, 0x04, 0x18, 0x32, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x2b, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x10, 0xfa, 0x42, 0x0d, 0x72, 0x0b, 0x52, 0x03, 0x61, 0x73, 0x63, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x15, 0xfa, 0x42, 0x12, 0x72, 0x10, 0x52, 0x02, 0x69, 0x64, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x48, 0x01, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x48, 0x02, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x28, 0x48, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x69, 0x6e, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x48, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd9, 0x02, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x4f, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70, 0x6f, 0x76, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0xda, 0x01, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x32, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x30, 0x0a, 0x1b, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x06, 0x43, 0x68, 0x61, 0x74, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x07,
	0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_poverty_chat_poverty_chat_record_proto_rawDescOnce sync.Once
	file_api_poverty_chat_poverty_chat_record_proto_rawDescData = file_api_poverty_chat_poverty_chat_record_proto_rawDesc
)

func file_api_poverty_chat_poverty_chat_record_proto_rawDescGZIP() []byte {
	file_api_poverty_chat_poverty_chat_record_proto_rawDescOnce.Do(func() {
		file_api_poverty_chat_poverty_chat_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_poverty_chat_poverty_chat_record_proto_rawDescData)
	})
	return file_api_poverty_chat_poverty_chat_record_proto_rawDescData
}

var file_api_poverty_chat_poverty_chat_record_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_poverty_chat_poverty_chat_record_proto_goTypes = []interface{}{
	(*SendChatMessageRequest)(nil),         // 0: poverty.api.poverty.chat.v1.SendChatMessageRequest
	(*SendChatMessageReply)(nil),           // 1: poverty.api.poverty.chat.v1.SendChatMessageReply
	(*ListChatRecordRequest)(nil),          // 2: poverty.api.poverty.chat.v1.ListChatRecordRequest
	(*ListChatRecordReply)(nil),            // 3: poverty.api.poverty.chat.v1.ListChatRecordReply
	(*DeleteChatRecordRequest)(nil),        // 4: poverty.api.poverty.chat.v1.DeleteChatRecordRequest
	(*DeleteChatRecordReply)(nil),          // 5: poverty.api.poverty.chat.v1.DeleteChatRecordReply
	(*ListChatRecordReply_ChatRecord)(nil), // 6: poverty.api.poverty.chat.v1.ListChatRecordReply.ChatRecord
}
var file_api_poverty_chat_poverty_chat_record_proto_depIdxs = []int32{
	6, // 0: poverty.api.poverty.chat.v1.ListChatRecordReply.list:type_name -> poverty.api.poverty.chat.v1.ListChatRecordReply.ChatRecord
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_poverty_chat_poverty_chat_record_proto_init() }
func file_api_poverty_chat_poverty_chat_record_proto_init() {
	if File_api_poverty_chat_poverty_chat_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatMessageRequest); i {
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
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChatMessageReply); i {
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
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatRecordRequest); i {
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
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatRecordReply); i {
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
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRecordRequest); i {
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
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRecordReply); i {
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
		file_api_poverty_chat_poverty_chat_record_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatRecordReply_ChatRecord); i {
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
	file_api_poverty_chat_poverty_chat_record_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_api_poverty_chat_poverty_chat_record_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_poverty_chat_poverty_chat_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_poverty_chat_poverty_chat_record_proto_goTypes,
		DependencyIndexes: file_api_poverty_chat_poverty_chat_record_proto_depIdxs,
		MessageInfos:      file_api_poverty_chat_poverty_chat_record_proto_msgTypes,
	}.Build()
	File_api_poverty_chat_poverty_chat_record_proto = out.File
	file_api_poverty_chat_poverty_chat_record_proto_rawDesc = nil
	file_api_poverty_chat_poverty_chat_record_proto_goTypes = nil
	file_api_poverty_chat_poverty_chat_record_proto_depIdxs = nil
}