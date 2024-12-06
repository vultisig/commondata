// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: vultisig/keygen/v1/lib_type_message.proto

package v1

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

// buf:lint:ignore ENUM_ZERO_VALUE_SUFFIX
type LibType int32

const (
	LibType_LIB_TYPE_GG20 LibType = 0 // Default to GG20
	LibType_LIB_TYPE_DKLS LibType = 1
)

// Enum value maps for LibType.
var (
	LibType_name = map[int32]string{
		0: "LIB_TYPE_GG20",
		1: "LIB_TYPE_DKLS",
	}
	LibType_value = map[string]int32{
		"LIB_TYPE_GG20": 0,
		"LIB_TYPE_DKLS": 1,
	}
)

func (x LibType) Enum() *LibType {
	p := new(LibType)
	*p = x
	return p
}

func (x LibType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LibType) Descriptor() protoreflect.EnumDescriptor {
	return file_vultisig_keygen_v1_lib_type_message_proto_enumTypes[0].Descriptor()
}

func (LibType) Type() protoreflect.EnumType {
	return &file_vultisig_keygen_v1_lib_type_message_proto_enumTypes[0]
}

func (x LibType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LibType.Descriptor instead.
func (LibType) EnumDescriptor() ([]byte, []int) {
	return file_vultisig_keygen_v1_lib_type_message_proto_rawDescGZIP(), []int{0}
}

var File_vultisig_keygen_v1_lib_type_message_proto protoreflect.FileDescriptor

var file_vultisig_keygen_v1_lib_type_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x79, 0x67, 0x65,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x75, 0x6c,
	0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2a,
	0x2f, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x49,
	0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x47, 0x32, 0x30, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x4c, 0x49, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4b, 0x4c, 0x53, 0x10, 0x01,
	0x42, 0x2d, 0x0a, 0x12, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x6b, 0x65, 0x79,
	0x67, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x12, 0x76, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67,
	0x2f, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xba, 0x02, 0x02, 0x56, 0x53, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vultisig_keygen_v1_lib_type_message_proto_rawDescOnce sync.Once
	file_vultisig_keygen_v1_lib_type_message_proto_rawDescData = file_vultisig_keygen_v1_lib_type_message_proto_rawDesc
)

func file_vultisig_keygen_v1_lib_type_message_proto_rawDescGZIP() []byte {
	file_vultisig_keygen_v1_lib_type_message_proto_rawDescOnce.Do(func() {
		file_vultisig_keygen_v1_lib_type_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_vultisig_keygen_v1_lib_type_message_proto_rawDescData)
	})
	return file_vultisig_keygen_v1_lib_type_message_proto_rawDescData
}

var file_vultisig_keygen_v1_lib_type_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vultisig_keygen_v1_lib_type_message_proto_goTypes = []any{
	(LibType)(0), // 0: vultisig.keygen.v1.LibType
}
var file_vultisig_keygen_v1_lib_type_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vultisig_keygen_v1_lib_type_message_proto_init() }
func file_vultisig_keygen_v1_lib_type_message_proto_init() {
	if File_vultisig_keygen_v1_lib_type_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vultisig_keygen_v1_lib_type_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vultisig_keygen_v1_lib_type_message_proto_goTypes,
		DependencyIndexes: file_vultisig_keygen_v1_lib_type_message_proto_depIdxs,
		EnumInfos:         file_vultisig_keygen_v1_lib_type_message_proto_enumTypes,
	}.Build()
	File_vultisig_keygen_v1_lib_type_message_proto = out.File
	file_vultisig_keygen_v1_lib_type_message_proto_rawDesc = nil
	file_vultisig_keygen_v1_lib_type_message_proto_goTypes = nil
	file_vultisig_keygen_v1_lib_type_message_proto_depIdxs = nil
}
