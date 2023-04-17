// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: sdk/stream/aggregated_quote_v2/response.proto

package aggregated_quote_v2

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

// EventType is event category for a value.
type StreamAggregatedQuoteResponseV2_EventType int32

const (
	// Unknown type.
	StreamAggregatedQuoteResponseV2_UNKNOWN StreamAggregatedQuoteResponseV2_EventType = 0
	// BEST_ASK type.
	StreamAggregatedQuoteResponseV2_BEST_ASK StreamAggregatedQuoteResponseV2_EventType = 1
	// BEST_BID type.
	StreamAggregatedQuoteResponseV2_BEST_BID StreamAggregatedQuoteResponseV2_EventType = 2
)

// Enum value maps for StreamAggregatedQuoteResponseV2_EventType.
var (
	StreamAggregatedQuoteResponseV2_EventType_name = map[int32]string{
		0: "UNKNOWN",
		1: "BEST_ASK",
		2: "BEST_BID",
	}
	StreamAggregatedQuoteResponseV2_EventType_value = map[string]int32{
		"UNKNOWN":  0,
		"BEST_ASK": 1,
		"BEST_BID": 2,
	}
)

func (x StreamAggregatedQuoteResponseV2_EventType) Enum() *StreamAggregatedQuoteResponseV2_EventType {
	p := new(StreamAggregatedQuoteResponseV2_EventType)
	*p = x
	return p
}

func (x StreamAggregatedQuoteResponseV2_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamAggregatedQuoteResponseV2_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_sdk_stream_aggregated_quote_v2_response_proto_enumTypes[0].Descriptor()
}

func (StreamAggregatedQuoteResponseV2_EventType) Type() protoreflect.EnumType {
	return &file_sdk_stream_aggregated_quote_v2_response_proto_enumTypes[0]
}

func (x StreamAggregatedQuoteResponseV2_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamAggregatedQuoteResponseV2_EventType.Descriptor instead.
func (StreamAggregatedQuoteResponseV2_EventType) EnumDescriptor() ([]byte, []int) {
	return file_sdk_stream_aggregated_quote_v2_response_proto_rawDescGZIP(), []int{1, 0}
}

// StreamAggregatedQuoteValue is wrapper for update values.
type StreamAggregatedQuoteValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Price of top of book entry.
	Price string `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	// Volume of top of book entry.
	Volume string `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *StreamAggregatedQuoteValue) Reset() {
	*x = StreamAggregatedQuoteValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAggregatedQuoteValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAggregatedQuoteValue) ProtoMessage() {}

func (x *StreamAggregatedQuoteValue) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAggregatedQuoteValue.ProtoReflect.Descriptor instead.
func (*StreamAggregatedQuoteValue) Descriptor() ([]byte, []int) {
	return file_sdk_stream_aggregated_quote_v2_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamAggregatedQuoteValue) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *StreamAggregatedQuoteValue) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

// StreamAggregatedQuoteResponseV2
type StreamAggregatedQuoteResponseV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Aggregate (interval).
	Aggregate string `protobuf:"bytes,1,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	// Instrument class. See https://docs.kaiko.com/?python#instruments.
	InstrumentClass string `protobuf:"bytes,2,opt,name=instrument_class,json=instrumentClass,proto3" json:"instrument_class,omitempty"`
	// Instrument code. See https://docs.kaiko.com/?python#instruments.
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// Event type is ASK or BID.
	EventType StreamAggregatedQuoteResponseV2_EventType `protobuf:"varint,4,opt,name=event_type,json=eventType,proto3,enum=kaikosdk.StreamAggregatedQuoteResponseV2_EventType" json:"event_type,omitempty"`
	// Timestamp of event.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,5,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
	// Vetted is update value (price and volume) using Kaiko vetted exchanges list.
	// Field is absent if the ticker has no update during the aggregate inerval or asset is not covered by vetted exchange list.
	Vetted *StreamAggregatedQuoteValue `protobuf:"bytes,6,opt,name=vetted,proto3" json:"vetted,omitempty"`
	// Unvetted value is same as value, but includes all exchanges covered by Kaiko Top Of Book Stream for computation.
	// This should be used whenever an asset is not part of the vetted list.
	Unvetted *StreamAggregatedQuoteValue `protobuf:"bytes,7,opt,name=unvetted,proto3" json:"unvetted,omitempty"`
}

func (x *StreamAggregatedQuoteResponseV2) Reset() {
	*x = StreamAggregatedQuoteResponseV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAggregatedQuoteResponseV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAggregatedQuoteResponseV2) ProtoMessage() {}

func (x *StreamAggregatedQuoteResponseV2) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAggregatedQuoteResponseV2.ProtoReflect.Descriptor instead.
func (*StreamAggregatedQuoteResponseV2) Descriptor() ([]byte, []int) {
	return file_sdk_stream_aggregated_quote_v2_response_proto_rawDescGZIP(), []int{1}
}

func (x *StreamAggregatedQuoteResponseV2) GetAggregate() string {
	if x != nil {
		return x.Aggregate
	}
	return ""
}

func (x *StreamAggregatedQuoteResponseV2) GetInstrumentClass() string {
	if x != nil {
		return x.InstrumentClass
	}
	return ""
}

func (x *StreamAggregatedQuoteResponseV2) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StreamAggregatedQuoteResponseV2) GetEventType() StreamAggregatedQuoteResponseV2_EventType {
	if x != nil {
		return x.EventType
	}
	return StreamAggregatedQuoteResponseV2_UNKNOWN
}

func (x *StreamAggregatedQuoteResponseV2) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

func (x *StreamAggregatedQuoteResponseV2) GetVetted() *StreamAggregatedQuoteValue {
	if x != nil {
		return x.Vetted
	}
	return nil
}

func (x *StreamAggregatedQuoteResponseV2) GetUnvetted() *StreamAggregatedQuoteValue {
	if x != nil {
		return x.Unvetted
	}
	return nil
}

var File_sdk_stream_aggregated_quote_v2_response_proto protoreflect.FileDescriptor

var file_sdk_stream_aggregated_quote_v2_response_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x32,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x1a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0xbf, 0x03, 0x0a, 0x1f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x74,
	0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x76, 0x65, 0x74, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x65, 0x74, 0x74, 0x65, 0x64,
	0x12, 0x40, 0x0a, 0x08, 0x75, 0x6e, 0x76, 0x65, 0x74, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x75, 0x6e, 0x76, 0x65, 0x74, 0x74,
	0x65, 0x64, 0x22, 0x34, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x42, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x53, 0x4b, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x45,
	0x53, 0x54, 0x5f, 0x42, 0x49, 0x44, 0x10, 0x02, 0x42, 0xa2, 0x01, 0x0a, 0x28, 0x63, 0x6f, 0x6d,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61,
	0x69, 0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x76, 0x32, 0x3b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x76, 0x32, 0xaa, 0x02, 0x21, 0x4b, 0x61, 0x69, 0x6b,
	0x6f, 0x53, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_aggregated_quote_v2_response_proto_rawDescOnce sync.Once
	file_sdk_stream_aggregated_quote_v2_response_proto_rawDescData = file_sdk_stream_aggregated_quote_v2_response_proto_rawDesc
)

func file_sdk_stream_aggregated_quote_v2_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_aggregated_quote_v2_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_aggregated_quote_v2_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_aggregated_quote_v2_response_proto_rawDescData)
	})
	return file_sdk_stream_aggregated_quote_v2_response_proto_rawDescData
}

var file_sdk_stream_aggregated_quote_v2_response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sdk_stream_aggregated_quote_v2_response_proto_goTypes = []interface{}{
	(StreamAggregatedQuoteResponseV2_EventType)(0), // 0: kaikosdk.StreamAggregatedQuoteResponseV2.EventType
	(*StreamAggregatedQuoteValue)(nil),             // 1: kaikosdk.StreamAggregatedQuoteValue
	(*StreamAggregatedQuoteResponseV2)(nil),        // 2: kaikosdk.StreamAggregatedQuoteResponseV2
	(*timestamp.Timestamp)(nil),                    // 3: google.protobuf.Timestamp
}
var file_sdk_stream_aggregated_quote_v2_response_proto_depIdxs = []int32{
	0, // 0: kaikosdk.StreamAggregatedQuoteResponseV2.event_type:type_name -> kaikosdk.StreamAggregatedQuoteResponseV2.EventType
	3, // 1: kaikosdk.StreamAggregatedQuoteResponseV2.ts_event:type_name -> google.protobuf.Timestamp
	1, // 2: kaikosdk.StreamAggregatedQuoteResponseV2.vetted:type_name -> kaikosdk.StreamAggregatedQuoteValue
	1, // 3: kaikosdk.StreamAggregatedQuoteResponseV2.unvetted:type_name -> kaikosdk.StreamAggregatedQuoteValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sdk_stream_aggregated_quote_v2_response_proto_init() }
func file_sdk_stream_aggregated_quote_v2_response_proto_init() {
	if File_sdk_stream_aggregated_quote_v2_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAggregatedQuoteValue); i {
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
		file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAggregatedQuoteResponseV2); i {
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
			RawDescriptor: file_sdk_stream_aggregated_quote_v2_response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_aggregated_quote_v2_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_aggregated_quote_v2_response_proto_depIdxs,
		EnumInfos:         file_sdk_stream_aggregated_quote_v2_response_proto_enumTypes,
		MessageInfos:      file_sdk_stream_aggregated_quote_v2_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_aggregated_quote_v2_response_proto = out.File
	file_sdk_stream_aggregated_quote_v2_response_proto_rawDesc = nil
	file_sdk_stream_aggregated_quote_v2_response_proto_goTypes = nil
	file_sdk_stream_aggregated_quote_v2_response_proto_depIdxs = nil
}