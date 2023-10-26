// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/index_v1/commodity.proto

package index_v1

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

// StreamIndexCommodity allows selecting kind of index update wanted.
type StreamIndexCommodity int32

const (
	// Unknown commodity.
	StreamIndexCommodity_SIC_UNKNOWN StreamIndexCommodity = 0
	// REAL_TIME commodity.
	StreamIndexCommodity_SIC_REAL_TIME StreamIndexCommodity = 1
	// DAILY_FIXING commodity.
	StreamIndexCommodity_SIC_DAILY_FIXING StreamIndexCommodity = 2
)

// Enum value maps for StreamIndexCommodity.
var (
	StreamIndexCommodity_name = map[int32]string{
		0: "SIC_UNKNOWN",
		1: "SIC_REAL_TIME",
		2: "SIC_DAILY_FIXING",
	}
	StreamIndexCommodity_value = map[string]int32{
		"SIC_UNKNOWN":      0,
		"SIC_REAL_TIME":    1,
		"SIC_DAILY_FIXING": 2,
	}
)

func (x StreamIndexCommodity) Enum() *StreamIndexCommodity {
	p := new(StreamIndexCommodity)
	*p = x
	return p
}

func (x StreamIndexCommodity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamIndexCommodity) Descriptor() protoreflect.EnumDescriptor {
	return file_sdk_stream_index_v1_commodity_proto_enumTypes[0].Descriptor()
}

func (StreamIndexCommodity) Type() protoreflect.EnumType {
	return &file_sdk_stream_index_v1_commodity_proto_enumTypes[0]
}

func (x StreamIndexCommodity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamIndexCommodity.Descriptor instead.
func (StreamIndexCommodity) EnumDescriptor() ([]byte, []int) {
	return file_sdk_stream_index_v1_commodity_proto_rawDescGZIP(), []int{0}
}

var File_sdk_stream_index_v1_commodity_proto protoreflect.FileDescriptor

var file_sdk_stream_index_v1_commodity_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2a,
	0x50, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x43, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x49, 0x43, 0x5f,
	0x52, 0x45, 0x41, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x49, 0x43, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x46, 0x49, 0x58, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x42, 0x77, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x31,
	0xaa, 0x02, 0x17, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sdk_stream_index_v1_commodity_proto_rawDescOnce sync.Once
	file_sdk_stream_index_v1_commodity_proto_rawDescData = file_sdk_stream_index_v1_commodity_proto_rawDesc
)

func file_sdk_stream_index_v1_commodity_proto_rawDescGZIP() []byte {
	file_sdk_stream_index_v1_commodity_proto_rawDescOnce.Do(func() {
		file_sdk_stream_index_v1_commodity_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_index_v1_commodity_proto_rawDescData)
	})
	return file_sdk_stream_index_v1_commodity_proto_rawDescData
}

var file_sdk_stream_index_v1_commodity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sdk_stream_index_v1_commodity_proto_goTypes = []interface{}{
	(StreamIndexCommodity)(0), // 0: kaikosdk.StreamIndexCommodity
}
var file_sdk_stream_index_v1_commodity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sdk_stream_index_v1_commodity_proto_init() }
func file_sdk_stream_index_v1_commodity_proto_init() {
	if File_sdk_stream_index_v1_commodity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sdk_stream_index_v1_commodity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_index_v1_commodity_proto_goTypes,
		DependencyIndexes: file_sdk_stream_index_v1_commodity_proto_depIdxs,
		EnumInfos:         file_sdk_stream_index_v1_commodity_proto_enumTypes,
	}.Build()
	File_sdk_stream_index_v1_commodity_proto = out.File
	file_sdk_stream_index_v1_commodity_proto_rawDesc = nil
	file_sdk_stream_index_v1_commodity_proto_goTypes = nil
	file_sdk_stream_index_v1_commodity_proto_depIdxs = nil
}
