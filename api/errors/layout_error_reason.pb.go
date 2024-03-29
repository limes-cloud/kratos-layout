// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: layout_error_reason.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_NotFound  ErrorReason = 0
	ErrorReason_Database  ErrorReason = 1
	ErrorReason_Transform ErrorReason = 2
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "NotFound",
		1: "Database",
		2: "Transform",
	}
	ErrorReason_value = map[string]int32{
		"NotFound":  0,
		"Database":  1,
		"Transform": 2,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_layout_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_layout_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_layout_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_layout_error_reason_proto protoreflect.FileDescriptor

var file_layout_error_reason_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x7d, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0x00, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae,
	0xe4, 0xb8, 0x8d, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x20, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x10, 0x01, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe6, 0x95, 0xb0, 0xe6,
	0x8d, 0xae, 0xe5, 0xba, 0x93, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x24, 0x0a, 0x09, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x10, 0x02, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe6,
	0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe8, 0xbd, 0xac, 0xe6, 0x8d, 0xa2, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4,
	0xa5, 0x1a, 0x04, 0xa0, 0x45, 0xc8, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_layout_error_reason_proto_rawDescOnce sync.Once
	file_layout_error_reason_proto_rawDescData = file_layout_error_reason_proto_rawDesc
)

func file_layout_error_reason_proto_rawDescGZIP() []byte {
	file_layout_error_reason_proto_rawDescOnce.Do(func() {
		file_layout_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_layout_error_reason_proto_rawDescData)
	})
	return file_layout_error_reason_proto_rawDescData
}

var file_layout_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_layout_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: errors.ErrorReason
}
var file_layout_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_layout_error_reason_proto_init() }
func file_layout_error_reason_proto_init() {
	if File_layout_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_layout_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_layout_error_reason_proto_goTypes,
		DependencyIndexes: file_layout_error_reason_proto_depIdxs,
		EnumInfos:         file_layout_error_reason_proto_enumTypes,
	}.Build()
	File_layout_error_reason_proto = out.File
	file_layout_error_reason_proto_rawDesc = nil
	file_layout_error_reason_proto_goTypes = nil
	file_layout_error_reason_proto_depIdxs = nil
}
