// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: sdk/stream/aggregates_ohlcv_v1/response.proto

package aggregates_ohlcv_v1

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

// StreamAggregatesOHLCVResponseV1
type StreamAggregatesOHLCVResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregate (interval).
	Aggregate string `protobuf:"bytes,1,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	// Instrument class.
	Class string `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	// Close value of OHLCV.
	Close string `protobuf:"bytes,3,opt,name=close,proto3" json:"close,omitempty"`
	// Instrument exchange.
	Exchange string `protobuf:"bytes,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	// High value of OHLCV.
	High string `protobuf:"bytes,5,opt,name=high,proto3" json:"high,omitempty"`
	// Low value of OHLCV.
	Low string `protobuf:"bytes,6,opt,name=low,proto3" json:"low,omitempty"`
	// Open value of OHLCV.
	Open string `protobuf:"bytes,7,opt,name=open,proto3" json:"open,omitempty"`
	// Sequence ID for event. Sortable in lexicographic order.
	SequenceId string `protobuf:"bytes,8,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// Instrument code.
	Code string `protobuf:"bytes,9,opt,name=code,proto3" json:"code,omitempty"`
	// Timestamp of event.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// UID is timestamp truncated to interval of OHLCV.
	Uid string `protobuf:"bytes,11,opt,name=uid,proto3" json:"uid,omitempty"`
	// Volume of OHLCV.
	Volume string `protobuf:"bytes,12,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *StreamAggregatesOHLCVResponseV1) Reset() {
	*x = StreamAggregatesOHLCVResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_aggregates_ohlcv_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAggregatesOHLCVResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAggregatesOHLCVResponseV1) ProtoMessage() {}

func (x *StreamAggregatesOHLCVResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_aggregates_ohlcv_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAggregatesOHLCVResponseV1.ProtoReflect.Descriptor instead.
func (*StreamAggregatesOHLCVResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamAggregatesOHLCVResponseV1) GetAggregate() string {
	if x != nil {
		return x.Aggregate
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetClose() string {
	if x != nil {
		return x.Close
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetHigh() string {
	if x != nil {
		return x.High
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetLow() string {
	if x != nil {
		return x.Low
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetOpen() string {
	if x != nil {
		return x.Open
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetSequenceId() string {
	if x != nil {
		return x.SequenceId
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *StreamAggregatesOHLCVResponseV1) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *StreamAggregatesOHLCVResponseV1) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

var File_sdk_stream_aggregates_ohlcv_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x68, 0x6c, 0x63, 0x76, 0x5f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02, 0x0a, 0x1f, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x4f,
	0x48, 0x4c, 0x43, 0x56, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0xa2, 0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e,
	0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x68, 0x6c, 0x63,
	0x76, 0x5f, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69,
	0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x68, 0x6c, 0x63,
	0x76, 0x5f, 0x76, 0x31, 0x3b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x5f,
	0x6f, 0x68, 0x6c, 0x63, 0x76, 0x5f, 0x76, 0x31, 0xaa, 0x02, 0x21, 0x4b, 0x61, 0x69, 0x6b, 0x6f,
	0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x48, 0x4c, 0x43, 0x56, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescData = file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDesc
)

func file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDescData
}

var file_sdk_stream_aggregates_ohlcv_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_stream_aggregates_ohlcv_v1_response_proto_goTypes = []interface{}{
	(*StreamAggregatesOHLCVResponseV1)(nil), // 0: kaikosdk.StreamAggregatesOHLCVResponseV1
	(*timestamp.Timestamp)(nil),             // 1: google.protobuf.Timestamp
}
var file_sdk_stream_aggregates_ohlcv_v1_response_proto_depIdxs = []int32{
	1, // 0: kaikosdk.StreamAggregatesOHLCVResponseV1.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sdk_stream_aggregates_ohlcv_v1_response_proto_init() }
func file_sdk_stream_aggregates_ohlcv_v1_response_proto_init() {
	if File_sdk_stream_aggregates_ohlcv_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_aggregates_ohlcv_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAggregatesOHLCVResponseV1); i {
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
			RawDescriptor: file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_aggregates_ohlcv_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_aggregates_ohlcv_v1_response_proto_depIdxs,
		MessageInfos:      file_sdk_stream_aggregates_ohlcv_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_aggregates_ohlcv_v1_response_proto = out.File
	file_sdk_stream_aggregates_ohlcv_v1_response_proto_rawDesc = nil
	file_sdk_stream_aggregates_ohlcv_v1_response_proto_goTypes = nil
	file_sdk_stream_aggregates_ohlcv_v1_response_proto_depIdxs = nil
}
