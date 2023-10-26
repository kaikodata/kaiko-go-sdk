// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: sdk/stream/index_multi_assets_v1/response.proto

package index_multi_assets_v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	core "github.com/kaikodata/kaiko-go-sdk/core"
	index_v1 "github.com/kaikodata/kaiko-go-sdk/stream/index_v1"
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

// StreamIndexMultiAssetsServiceResponseComposition is the composition used to compute the index.
type StreamIndexMultiAssetsServiceResponseComposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Underlying rate of the indices.
	UnderlyingInstrument string `protobuf:"bytes,1,opt,name=underlying_instrument,json=underlyingInstrument,proto3" json:"underlying_instrument,omitempty"`
	// Base.
	Base string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	// Quote.
	Quote string `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	// Exchanges.
	Exchanges []string `protobuf:"bytes,4,rep,name=exchanges,proto3" json:"exchanges,omitempty"`
	// Data interval.
	InstrumentInterval *core.DataInterval `protobuf:"bytes,5,opt,name=instrument_interval,json=instrumentInterval,proto3" json:"instrument_interval,omitempty"`
	// Currency conversion.
	CurrencyConversion string `protobuf:"bytes,6,opt,name=currency_conversion,json=currencyConversion,proto3" json:"currency_conversion,omitempty"`
	// Timestamp (tick) of underlying rate ts.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) Reset() {
	*x = StreamIndexMultiAssetsServiceResponseComposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIndexMultiAssetsServiceResponseComposition) ProtoMessage() {}

func (x *StreamIndexMultiAssetsServiceResponseComposition) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIndexMultiAssetsServiceResponseComposition.ProtoReflect.Descriptor instead.
func (*StreamIndexMultiAssetsServiceResponseComposition) Descriptor() ([]byte, []int) {
	return file_sdk_stream_index_multi_assets_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetUnderlyingInstrument() string {
	if x != nil {
		return x.UnderlyingInstrument
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetExchanges() []string {
	if x != nil {
		return x.Exchanges
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetInstrumentInterval() *core.DataInterval {
	if x != nil {
		return x.InstrumentInterval
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetCurrencyConversion() string {
	if x != nil {
		return x.CurrencyConversion
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponseComposition) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

// StreamIndexMultiAssetsServiceResponsePair is the pair information for the rates used
type StreamIndexMultiAssetsServiceResponsePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Underlying instrument of the indices.
	UnderlyingInstrument string `protobuf:"bytes,1,opt,name=underlying_instrument,json=underlyingInstrument,proto3" json:"underlying_instrument,omitempty"`
	// Underlying price of the instrument.
	UnderlyingPrice *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=underlying_price,json=underlyingPrice,proto3" json:"underlying_price,omitempty"`
	// Weighting factor of the instrument.
	WeightingFactor float64 `protobuf:"fixed64,3,opt,name=weighting_factor,json=weightingFactor,proto3" json:"weighting_factor,omitempty"`
	// Capping factor of the instrument.
	CappingFactor float64 `protobuf:"fixed64,4,opt,name=capping_factor,json=cappingFactor,proto3" json:"capping_factor,omitempty"`
	// Currency conversion factor of the instrument.
	CurrencyConversionFactor float64 `protobuf:"fixed64,5,opt,name=currency_conversion_factor,json=currencyConversionFactor,proto3" json:"currency_conversion_factor,omitempty"`
}

func (x *StreamIndexMultiAssetsServiceResponsePair) Reset() {
	*x = StreamIndexMultiAssetsServiceResponsePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIndexMultiAssetsServiceResponsePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIndexMultiAssetsServiceResponsePair) ProtoMessage() {}

func (x *StreamIndexMultiAssetsServiceResponsePair) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIndexMultiAssetsServiceResponsePair.ProtoReflect.Descriptor instead.
func (*StreamIndexMultiAssetsServiceResponsePair) Descriptor() ([]byte, []int) {
	return file_sdk_stream_index_multi_assets_v1_response_proto_rawDescGZIP(), []int{1}
}

func (x *StreamIndexMultiAssetsServiceResponsePair) GetUnderlyingInstrument() string {
	if x != nil {
		return x.UnderlyingInstrument
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponsePair) GetUnderlyingPrice() *wrappers.DoubleValue {
	if x != nil {
		return x.UnderlyingPrice
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponsePair) GetWeightingFactor() float64 {
	if x != nil {
		return x.WeightingFactor
	}
	return 0
}

func (x *StreamIndexMultiAssetsServiceResponsePair) GetCappingFactor() float64 {
	if x != nil {
		return x.CappingFactor
	}
	return 0
}

func (x *StreamIndexMultiAssetsServiceResponsePair) GetCurrencyConversionFactor() float64 {
	if x != nil {
		return x.CurrencyConversionFactor
	}
	return 0
}

// StreamIndexMultiAssetsServiceResponsePrices is the prices informations on the pair used
type StreamIndexMultiAssetsServiceResponsePrices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index value.
	IndexValue float64 `protobuf:"fixed64,1,opt,name=index_value,json=indexValue,proto3" json:"index_value,omitempty"`
	// Divisor.
	Divisor float64 `protobuf:"fixed64,2,opt,name=divisor,proto3" json:"divisor,omitempty"`
	// StreamIndexMultiAssetsServiceResponsePair is the pair information for the rates used.
	Pairs []*StreamIndexMultiAssetsServiceResponsePair `protobuf:"bytes,3,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *StreamIndexMultiAssetsServiceResponsePrices) Reset() {
	*x = StreamIndexMultiAssetsServiceResponsePrices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIndexMultiAssetsServiceResponsePrices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIndexMultiAssetsServiceResponsePrices) ProtoMessage() {}

func (x *StreamIndexMultiAssetsServiceResponsePrices) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIndexMultiAssetsServiceResponsePrices.ProtoReflect.Descriptor instead.
func (*StreamIndexMultiAssetsServiceResponsePrices) Descriptor() ([]byte, []int) {
	return file_sdk_stream_index_multi_assets_v1_response_proto_rawDescGZIP(), []int{2}
}

func (x *StreamIndexMultiAssetsServiceResponsePrices) GetIndexValue() float64 {
	if x != nil {
		return x.IndexValue
	}
	return 0
}

func (x *StreamIndexMultiAssetsServiceResponsePrices) GetDivisor() float64 {
	if x != nil {
		return x.Divisor
	}
	return 0
}

func (x *StreamIndexMultiAssetsServiceResponsePrices) GetPairs() []*StreamIndexMultiAssetsServiceResponsePair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// StreamIndexServiceResponseV1
type StreamIndexMultiAssetsServiceResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event type.
	Commodity index_v1.StreamIndexCommodity `protobuf:"varint,1,opt,name=commodity,proto3,enum=kaikosdk.StreamIndexCommodity" json:"commodity,omitempty"`
	// Index code.
	IndexCode string `protobuf:"bytes,2,opt,name=index_code,json=indexCode,proto3" json:"index_code,omitempty"`
	// Data interval.
	Interval *core.DataInterval `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	// Quote.
	MainQuote string `protobuf:"bytes,4,opt,name=main_quote,json=mainQuote,proto3" json:"main_quote,omitempty"`
	// List of rates used in indices.
	Compositions []*StreamIndexMultiAssetsServiceResponseComposition `protobuf:"bytes,5,rep,name=compositions,proto3" json:"compositions,omitempty"`
	// Price of the indice.
	Price *StreamIndexMultiAssetsServiceResponsePrices `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	// Event generation timestamp (event created by Kaiko), after normalization.
	TsEvent *timestamp.Timestamp `protobuf:"bytes,7,opt,name=ts_event,json=tsEvent,proto3" json:"ts_event,omitempty"`
	// Timestamp of computation (differs from ts_event only if a buffer is applied, especially for real time data).
	TsCompute *timestamp.Timestamp `protobuf:"bytes,8,opt,name=ts_compute,json=tsCompute,proto3" json:"ts_compute,omitempty"`
}

func (x *StreamIndexMultiAssetsServiceResponseV1) Reset() {
	*x = StreamIndexMultiAssetsServiceResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamIndexMultiAssetsServiceResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamIndexMultiAssetsServiceResponseV1) ProtoMessage() {}

func (x *StreamIndexMultiAssetsServiceResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamIndexMultiAssetsServiceResponseV1.ProtoReflect.Descriptor instead.
func (*StreamIndexMultiAssetsServiceResponseV1) Descriptor() ([]byte, []int) {
	return file_sdk_stream_index_multi_assets_v1_response_proto_rawDescGZIP(), []int{3}
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetCommodity() index_v1.StreamIndexCommodity {
	if x != nil {
		return x.Commodity
	}
	return index_v1.StreamIndexCommodity(0)
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetIndexCode() string {
	if x != nil {
		return x.IndexCode
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetInterval() *core.DataInterval {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetMainQuote() string {
	if x != nil {
		return x.MainQuote
	}
	return ""
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetCompositions() []*StreamIndexMultiAssetsServiceResponseComposition {
	if x != nil {
		return x.Compositions
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetPrice() *StreamIndexMultiAssetsServiceResponsePrices {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetTsEvent() *timestamp.Timestamp {
	if x != nil {
		return x.TsEvent
	}
	return nil
}

func (x *StreamIndexMultiAssetsServiceResponseV1) GetTsCompute() *timestamp.Timestamp {
	if x != nil {
		return x.TsCompute
	}
	return nil
}

var File_sdk_stream_index_multi_assets_v1_response_proto protoreflect.FileDescriptor

var file_sdk_stream_index_multi_assets_v1_response_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x64,
	0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x73, 0x64, 0x6b, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe0, 0x02, 0x0a, 0x30, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x15, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x47, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x74,
	0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x73, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0xb9, 0x02, 0x0a, 0x29, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x33, 0x0a, 0x15, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x10, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x75,
	0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x69, 0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x3c, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x18, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xb3,
	0x01, 0x0a, 0x2b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x69, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x64, 0x69, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f,
	0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x69, 0x72, 0x73, 0x22, 0xf8, 0x03, 0x0a, 0x27, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31,
	0x12, 0x3c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x12, 0x5e, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64,
	0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4b, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x74, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x73, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x42,
	0xa9, 0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x69,
	0x6b, 0x6f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6b, 0x61, 0x69, 0x6b, 0x6f, 0x2d, 0x67, 0x6f, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x76, 0x31,
	0x3b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x5f, 0x76, 0x31, 0xaa, 0x02, 0x22, 0x4b, 0x61, 0x69, 0x6b, 0x6f, 0x53, 0x64,
	0x6b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sdk_stream_index_multi_assets_v1_response_proto_rawDescOnce sync.Once
	file_sdk_stream_index_multi_assets_v1_response_proto_rawDescData = file_sdk_stream_index_multi_assets_v1_response_proto_rawDesc
)

func file_sdk_stream_index_multi_assets_v1_response_proto_rawDescGZIP() []byte {
	file_sdk_stream_index_multi_assets_v1_response_proto_rawDescOnce.Do(func() {
		file_sdk_stream_index_multi_assets_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_stream_index_multi_assets_v1_response_proto_rawDescData)
	})
	return file_sdk_stream_index_multi_assets_v1_response_proto_rawDescData
}

var file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sdk_stream_index_multi_assets_v1_response_proto_goTypes = []interface{}{
	(*StreamIndexMultiAssetsServiceResponseComposition)(nil), // 0: kaikosdk.StreamIndexMultiAssetsServiceResponseComposition
	(*StreamIndexMultiAssetsServiceResponsePair)(nil),        // 1: kaikosdk.StreamIndexMultiAssetsServiceResponsePair
	(*StreamIndexMultiAssetsServiceResponsePrices)(nil),      // 2: kaikosdk.StreamIndexMultiAssetsServiceResponsePrices
	(*StreamIndexMultiAssetsServiceResponseV1)(nil),          // 3: kaikosdk.StreamIndexMultiAssetsServiceResponseV1
	(*core.DataInterval)(nil),                                // 4: kaikosdk.DataInterval
	(*timestamp.Timestamp)(nil),                              // 5: google.protobuf.Timestamp
	(*wrappers.DoubleValue)(nil),                             // 6: google.protobuf.DoubleValue
	(index_v1.StreamIndexCommodity)(0),                       // 7: kaikosdk.StreamIndexCommodity
}
var file_sdk_stream_index_multi_assets_v1_response_proto_depIdxs = []int32{
	4,  // 0: kaikosdk.StreamIndexMultiAssetsServiceResponseComposition.instrument_interval:type_name -> kaikosdk.DataInterval
	5,  // 1: kaikosdk.StreamIndexMultiAssetsServiceResponseComposition.ts_event:type_name -> google.protobuf.Timestamp
	6,  // 2: kaikosdk.StreamIndexMultiAssetsServiceResponsePair.underlying_price:type_name -> google.protobuf.DoubleValue
	1,  // 3: kaikosdk.StreamIndexMultiAssetsServiceResponsePrices.pairs:type_name -> kaikosdk.StreamIndexMultiAssetsServiceResponsePair
	7,  // 4: kaikosdk.StreamIndexMultiAssetsServiceResponseV1.commodity:type_name -> kaikosdk.StreamIndexCommodity
	4,  // 5: kaikosdk.StreamIndexMultiAssetsServiceResponseV1.interval:type_name -> kaikosdk.DataInterval
	0,  // 6: kaikosdk.StreamIndexMultiAssetsServiceResponseV1.compositions:type_name -> kaikosdk.StreamIndexMultiAssetsServiceResponseComposition
	2,  // 7: kaikosdk.StreamIndexMultiAssetsServiceResponseV1.price:type_name -> kaikosdk.StreamIndexMultiAssetsServiceResponsePrices
	5,  // 8: kaikosdk.StreamIndexMultiAssetsServiceResponseV1.ts_event:type_name -> google.protobuf.Timestamp
	5,  // 9: kaikosdk.StreamIndexMultiAssetsServiceResponseV1.ts_compute:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_sdk_stream_index_multi_assets_v1_response_proto_init() }
func file_sdk_stream_index_multi_assets_v1_response_proto_init() {
	if File_sdk_stream_index_multi_assets_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIndexMultiAssetsServiceResponseComposition); i {
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
		file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIndexMultiAssetsServiceResponsePair); i {
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
		file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIndexMultiAssetsServiceResponsePrices); i {
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
		file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamIndexMultiAssetsServiceResponseV1); i {
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
			RawDescriptor: file_sdk_stream_index_multi_assets_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_stream_index_multi_assets_v1_response_proto_goTypes,
		DependencyIndexes: file_sdk_stream_index_multi_assets_v1_response_proto_depIdxs,
		MessageInfos:      file_sdk_stream_index_multi_assets_v1_response_proto_msgTypes,
	}.Build()
	File_sdk_stream_index_multi_assets_v1_response_proto = out.File
	file_sdk_stream_index_multi_assets_v1_response_proto_rawDesc = nil
	file_sdk_stream_index_multi_assets_v1_response_proto_goTypes = nil
	file_sdk_stream_index_multi_assets_v1_response_proto_depIdxs = nil
}
