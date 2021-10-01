// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: sdk/core/source.proto

package core

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

// Source
type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prices for instrument code in each exchange.
	Data []*SourceData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// VWAP value for related instrument code accross exchanges.
	Price string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_core_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_core_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_sdk_core_source_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetData() []*SourceData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Source) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

var File_sdk_core_source_proto protoreflect.FileDescriptor

var file_sdk_core_source_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64,
	0x6b, 0x1a, 0x1a, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a,
	0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x58, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6b,
	0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x64, 0x65, 0x65, 0x70, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72,
	0x65, 0xaa, 0x02, 0x0d, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x43, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_core_source_proto_rawDescOnce sync.Once
	file_sdk_core_source_proto_rawDescData = file_sdk_core_source_proto_rawDesc
)

func file_sdk_core_source_proto_rawDescGZIP() []byte {
	file_sdk_core_source_proto_rawDescOnce.Do(func() {
		file_sdk_core_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_core_source_proto_rawDescData)
	})
	return file_sdk_core_source_proto_rawDescData
}

var file_sdk_core_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_core_source_proto_goTypes = []interface{}{
	(*Source)(nil),     // 0: kaikosdk.Source
	(*SourceData)(nil), // 1: kaikosdk.SourceData
}
var file_sdk_core_source_proto_depIdxs = []int32{
	1, // 0: kaikosdk.Source.data:type_name -> kaikosdk.SourceData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sdk_core_source_proto_init() }
func file_sdk_core_source_proto_init() {
	if File_sdk_core_source_proto != nil {
		return
	}
	file_sdk_core_source_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sdk_core_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
			RawDescriptor: file_sdk_core_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_core_source_proto_goTypes,
		DependencyIndexes: file_sdk_core_source_proto_depIdxs,
		MessageInfos:      file_sdk_core_source_proto_msgTypes,
	}.Build()
	File_sdk_core_source_proto = out.File
	file_sdk_core_source_proto_rawDesc = nil
	file_sdk_core_source_proto_goTypes = nil
	file_sdk_core_source_proto_depIdxs = nil
}
