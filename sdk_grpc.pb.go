// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kaikosdk

import (
	context "context"
	aggregates_ohlcv_v1 "github.com/challengerdeep/kaiko-go-sdk/stream/aggregates_ohlcv_v1"
	aggregates_spot_exchange_rate_v1 "github.com/challengerdeep/kaiko-go-sdk/stream/aggregates_spot_exchange_rate_v1"
	aggregates_vwap_v1 "github.com/challengerdeep/kaiko-go-sdk/stream/aggregates_vwap_v1"
	market_update_v1 "github.com/challengerdeep/kaiko-go-sdk/stream/market_update_v1"
	trades_v1 "github.com/challengerdeep/kaiko-go-sdk/stream/trades_v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamAggregatesOHLCVServiceV1Client is the client API for StreamAggregatesOHLCVServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamAggregatesOHLCVServiceV1Client interface {
	// Subscribe
	Subscribe(ctx context.Context, in *aggregates_ohlcv_v1.StreamAggregatesOHLCVRequestV1, opts ...grpc.CallOption) (StreamAggregatesOHLCVServiceV1_SubscribeClient, error)
}

type streamAggregatesOHLCVServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStreamAggregatesOHLCVServiceV1Client(cc grpc.ClientConnInterface) StreamAggregatesOHLCVServiceV1Client {
	return &streamAggregatesOHLCVServiceV1Client{cc}
}

func (c *streamAggregatesOHLCVServiceV1Client) Subscribe(ctx context.Context, in *aggregates_ohlcv_v1.StreamAggregatesOHLCVRequestV1, opts ...grpc.CallOption) (StreamAggregatesOHLCVServiceV1_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamAggregatesOHLCVServiceV1_ServiceDesc.Streams[0], "/kaikosdk.StreamAggregatesOHLCVServiceV1/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamAggregatesOHLCVServiceV1SubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamAggregatesOHLCVServiceV1_SubscribeClient interface {
	Recv() (*aggregates_ohlcv_v1.StreamAggregatesOHLCVResponseV1, error)
	grpc.ClientStream
}

type streamAggregatesOHLCVServiceV1SubscribeClient struct {
	grpc.ClientStream
}

func (x *streamAggregatesOHLCVServiceV1SubscribeClient) Recv() (*aggregates_ohlcv_v1.StreamAggregatesOHLCVResponseV1, error) {
	m := new(aggregates_ohlcv_v1.StreamAggregatesOHLCVResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamAggregatesOHLCVServiceV1Server is the server API for StreamAggregatesOHLCVServiceV1 service.
// All implementations must embed UnimplementedStreamAggregatesOHLCVServiceV1Server
// for forward compatibility
type StreamAggregatesOHLCVServiceV1Server interface {
	// Subscribe
	Subscribe(*aggregates_ohlcv_v1.StreamAggregatesOHLCVRequestV1, StreamAggregatesOHLCVServiceV1_SubscribeServer) error
	mustEmbedUnimplementedStreamAggregatesOHLCVServiceV1Server()
}

// UnimplementedStreamAggregatesOHLCVServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedStreamAggregatesOHLCVServiceV1Server struct {
}

func (UnimplementedStreamAggregatesOHLCVServiceV1Server) Subscribe(*aggregates_ohlcv_v1.StreamAggregatesOHLCVRequestV1, StreamAggregatesOHLCVServiceV1_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedStreamAggregatesOHLCVServiceV1Server) mustEmbedUnimplementedStreamAggregatesOHLCVServiceV1Server() {
}

// UnsafeStreamAggregatesOHLCVServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamAggregatesOHLCVServiceV1Server will
// result in compilation errors.
type UnsafeStreamAggregatesOHLCVServiceV1Server interface {
	mustEmbedUnimplementedStreamAggregatesOHLCVServiceV1Server()
}

func RegisterStreamAggregatesOHLCVServiceV1Server(s grpc.ServiceRegistrar, srv StreamAggregatesOHLCVServiceV1Server) {
	s.RegisterService(&StreamAggregatesOHLCVServiceV1_ServiceDesc, srv)
}

func _StreamAggregatesOHLCVServiceV1_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(aggregates_ohlcv_v1.StreamAggregatesOHLCVRequestV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamAggregatesOHLCVServiceV1Server).Subscribe(m, &streamAggregatesOHLCVServiceV1SubscribeServer{stream})
}

type StreamAggregatesOHLCVServiceV1_SubscribeServer interface {
	Send(*aggregates_ohlcv_v1.StreamAggregatesOHLCVResponseV1) error
	grpc.ServerStream
}

type streamAggregatesOHLCVServiceV1SubscribeServer struct {
	grpc.ServerStream
}

func (x *streamAggregatesOHLCVServiceV1SubscribeServer) Send(m *aggregates_ohlcv_v1.StreamAggregatesOHLCVResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

// StreamAggregatesOHLCVServiceV1_ServiceDesc is the grpc.ServiceDesc for StreamAggregatesOHLCVServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamAggregatesOHLCVServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaikosdk.StreamAggregatesOHLCVServiceV1",
	HandlerType: (*StreamAggregatesOHLCVServiceV1Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _StreamAggregatesOHLCVServiceV1_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk/sdk.proto",
}

// StreamAggregatesSpotExchangeRateServiceV1Client is the client API for StreamAggregatesSpotExchangeRateServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamAggregatesSpotExchangeRateServiceV1Client interface {
	// Subscribe
	Subscribe(ctx context.Context, in *aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateRequestV1, opts ...grpc.CallOption) (StreamAggregatesSpotExchangeRateServiceV1_SubscribeClient, error)
}

type streamAggregatesSpotExchangeRateServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStreamAggregatesSpotExchangeRateServiceV1Client(cc grpc.ClientConnInterface) StreamAggregatesSpotExchangeRateServiceV1Client {
	return &streamAggregatesSpotExchangeRateServiceV1Client{cc}
}

func (c *streamAggregatesSpotExchangeRateServiceV1Client) Subscribe(ctx context.Context, in *aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateRequestV1, opts ...grpc.CallOption) (StreamAggregatesSpotExchangeRateServiceV1_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamAggregatesSpotExchangeRateServiceV1_ServiceDesc.Streams[0], "/kaikosdk.StreamAggregatesSpotExchangeRateServiceV1/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamAggregatesSpotExchangeRateServiceV1SubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamAggregatesSpotExchangeRateServiceV1_SubscribeClient interface {
	Recv() (*aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateResponseV1, error)
	grpc.ClientStream
}

type streamAggregatesSpotExchangeRateServiceV1SubscribeClient struct {
	grpc.ClientStream
}

func (x *streamAggregatesSpotExchangeRateServiceV1SubscribeClient) Recv() (*aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateResponseV1, error) {
	m := new(aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamAggregatesSpotExchangeRateServiceV1Server is the server API for StreamAggregatesSpotExchangeRateServiceV1 service.
// All implementations must embed UnimplementedStreamAggregatesSpotExchangeRateServiceV1Server
// for forward compatibility
type StreamAggregatesSpotExchangeRateServiceV1Server interface {
	// Subscribe
	Subscribe(*aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateRequestV1, StreamAggregatesSpotExchangeRateServiceV1_SubscribeServer) error
	mustEmbedUnimplementedStreamAggregatesSpotExchangeRateServiceV1Server()
}

// UnimplementedStreamAggregatesSpotExchangeRateServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedStreamAggregatesSpotExchangeRateServiceV1Server struct {
}

func (UnimplementedStreamAggregatesSpotExchangeRateServiceV1Server) Subscribe(*aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateRequestV1, StreamAggregatesSpotExchangeRateServiceV1_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedStreamAggregatesSpotExchangeRateServiceV1Server) mustEmbedUnimplementedStreamAggregatesSpotExchangeRateServiceV1Server() {
}

// UnsafeStreamAggregatesSpotExchangeRateServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamAggregatesSpotExchangeRateServiceV1Server will
// result in compilation errors.
type UnsafeStreamAggregatesSpotExchangeRateServiceV1Server interface {
	mustEmbedUnimplementedStreamAggregatesSpotExchangeRateServiceV1Server()
}

func RegisterStreamAggregatesSpotExchangeRateServiceV1Server(s grpc.ServiceRegistrar, srv StreamAggregatesSpotExchangeRateServiceV1Server) {
	s.RegisterService(&StreamAggregatesSpotExchangeRateServiceV1_ServiceDesc, srv)
}

func _StreamAggregatesSpotExchangeRateServiceV1_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateRequestV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamAggregatesSpotExchangeRateServiceV1Server).Subscribe(m, &streamAggregatesSpotExchangeRateServiceV1SubscribeServer{stream})
}

type StreamAggregatesSpotExchangeRateServiceV1_SubscribeServer interface {
	Send(*aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateResponseV1) error
	grpc.ServerStream
}

type streamAggregatesSpotExchangeRateServiceV1SubscribeServer struct {
	grpc.ServerStream
}

func (x *streamAggregatesSpotExchangeRateServiceV1SubscribeServer) Send(m *aggregates_spot_exchange_rate_v1.StreamAggregatesSpotExchangeRateResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

// StreamAggregatesSpotExchangeRateServiceV1_ServiceDesc is the grpc.ServiceDesc for StreamAggregatesSpotExchangeRateServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamAggregatesSpotExchangeRateServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaikosdk.StreamAggregatesSpotExchangeRateServiceV1",
	HandlerType: (*StreamAggregatesSpotExchangeRateServiceV1Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _StreamAggregatesSpotExchangeRateServiceV1_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk/sdk.proto",
}

// StreamTradesServiceV1Client is the client API for StreamTradesServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamTradesServiceV1Client interface {
	// Subscribe
	Subscribe(ctx context.Context, in *trades_v1.StreamTradesRequestV1, opts ...grpc.CallOption) (StreamTradesServiceV1_SubscribeClient, error)
}

type streamTradesServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStreamTradesServiceV1Client(cc grpc.ClientConnInterface) StreamTradesServiceV1Client {
	return &streamTradesServiceV1Client{cc}
}

func (c *streamTradesServiceV1Client) Subscribe(ctx context.Context, in *trades_v1.StreamTradesRequestV1, opts ...grpc.CallOption) (StreamTradesServiceV1_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamTradesServiceV1_ServiceDesc.Streams[0], "/kaikosdk.StreamTradesServiceV1/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamTradesServiceV1SubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamTradesServiceV1_SubscribeClient interface {
	Recv() (*trades_v1.StreamTradesResponseV1, error)
	grpc.ClientStream
}

type streamTradesServiceV1SubscribeClient struct {
	grpc.ClientStream
}

func (x *streamTradesServiceV1SubscribeClient) Recv() (*trades_v1.StreamTradesResponseV1, error) {
	m := new(trades_v1.StreamTradesResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamTradesServiceV1Server is the server API for StreamTradesServiceV1 service.
// All implementations must embed UnimplementedStreamTradesServiceV1Server
// for forward compatibility
type StreamTradesServiceV1Server interface {
	// Subscribe
	Subscribe(*trades_v1.StreamTradesRequestV1, StreamTradesServiceV1_SubscribeServer) error
	mustEmbedUnimplementedStreamTradesServiceV1Server()
}

// UnimplementedStreamTradesServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedStreamTradesServiceV1Server struct {
}

func (UnimplementedStreamTradesServiceV1Server) Subscribe(*trades_v1.StreamTradesRequestV1, StreamTradesServiceV1_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedStreamTradesServiceV1Server) mustEmbedUnimplementedStreamTradesServiceV1Server() {}

// UnsafeStreamTradesServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamTradesServiceV1Server will
// result in compilation errors.
type UnsafeStreamTradesServiceV1Server interface {
	mustEmbedUnimplementedStreamTradesServiceV1Server()
}

func RegisterStreamTradesServiceV1Server(s grpc.ServiceRegistrar, srv StreamTradesServiceV1Server) {
	s.RegisterService(&StreamTradesServiceV1_ServiceDesc, srv)
}

func _StreamTradesServiceV1_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(trades_v1.StreamTradesRequestV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamTradesServiceV1Server).Subscribe(m, &streamTradesServiceV1SubscribeServer{stream})
}

type StreamTradesServiceV1_SubscribeServer interface {
	Send(*trades_v1.StreamTradesResponseV1) error
	grpc.ServerStream
}

type streamTradesServiceV1SubscribeServer struct {
	grpc.ServerStream
}

func (x *streamTradesServiceV1SubscribeServer) Send(m *trades_v1.StreamTradesResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

// StreamTradesServiceV1_ServiceDesc is the grpc.ServiceDesc for StreamTradesServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamTradesServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaikosdk.StreamTradesServiceV1",
	HandlerType: (*StreamTradesServiceV1Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _StreamTradesServiceV1_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk/sdk.proto",
}

// StreamAggregatesVWAPServiceV1Client is the client API for StreamAggregatesVWAPServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamAggregatesVWAPServiceV1Client interface {
	// Subscribe
	Subscribe(ctx context.Context, in *aggregates_vwap_v1.StreamAggregatesVWAPRequestV1, opts ...grpc.CallOption) (StreamAggregatesVWAPServiceV1_SubscribeClient, error)
}

type streamAggregatesVWAPServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStreamAggregatesVWAPServiceV1Client(cc grpc.ClientConnInterface) StreamAggregatesVWAPServiceV1Client {
	return &streamAggregatesVWAPServiceV1Client{cc}
}

func (c *streamAggregatesVWAPServiceV1Client) Subscribe(ctx context.Context, in *aggregates_vwap_v1.StreamAggregatesVWAPRequestV1, opts ...grpc.CallOption) (StreamAggregatesVWAPServiceV1_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamAggregatesVWAPServiceV1_ServiceDesc.Streams[0], "/kaikosdk.StreamAggregatesVWAPServiceV1/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamAggregatesVWAPServiceV1SubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamAggregatesVWAPServiceV1_SubscribeClient interface {
	Recv() (*aggregates_vwap_v1.StreamAggregatesVWAPResponseV1, error)
	grpc.ClientStream
}

type streamAggregatesVWAPServiceV1SubscribeClient struct {
	grpc.ClientStream
}

func (x *streamAggregatesVWAPServiceV1SubscribeClient) Recv() (*aggregates_vwap_v1.StreamAggregatesVWAPResponseV1, error) {
	m := new(aggregates_vwap_v1.StreamAggregatesVWAPResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamAggregatesVWAPServiceV1Server is the server API for StreamAggregatesVWAPServiceV1 service.
// All implementations must embed UnimplementedStreamAggregatesVWAPServiceV1Server
// for forward compatibility
type StreamAggregatesVWAPServiceV1Server interface {
	// Subscribe
	Subscribe(*aggregates_vwap_v1.StreamAggregatesVWAPRequestV1, StreamAggregatesVWAPServiceV1_SubscribeServer) error
	mustEmbedUnimplementedStreamAggregatesVWAPServiceV1Server()
}

// UnimplementedStreamAggregatesVWAPServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedStreamAggregatesVWAPServiceV1Server struct {
}

func (UnimplementedStreamAggregatesVWAPServiceV1Server) Subscribe(*aggregates_vwap_v1.StreamAggregatesVWAPRequestV1, StreamAggregatesVWAPServiceV1_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedStreamAggregatesVWAPServiceV1Server) mustEmbedUnimplementedStreamAggregatesVWAPServiceV1Server() {
}

// UnsafeStreamAggregatesVWAPServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamAggregatesVWAPServiceV1Server will
// result in compilation errors.
type UnsafeStreamAggregatesVWAPServiceV1Server interface {
	mustEmbedUnimplementedStreamAggregatesVWAPServiceV1Server()
}

func RegisterStreamAggregatesVWAPServiceV1Server(s grpc.ServiceRegistrar, srv StreamAggregatesVWAPServiceV1Server) {
	s.RegisterService(&StreamAggregatesVWAPServiceV1_ServiceDesc, srv)
}

func _StreamAggregatesVWAPServiceV1_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(aggregates_vwap_v1.StreamAggregatesVWAPRequestV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamAggregatesVWAPServiceV1Server).Subscribe(m, &streamAggregatesVWAPServiceV1SubscribeServer{stream})
}

type StreamAggregatesVWAPServiceV1_SubscribeServer interface {
	Send(*aggregates_vwap_v1.StreamAggregatesVWAPResponseV1) error
	grpc.ServerStream
}

type streamAggregatesVWAPServiceV1SubscribeServer struct {
	grpc.ServerStream
}

func (x *streamAggregatesVWAPServiceV1SubscribeServer) Send(m *aggregates_vwap_v1.StreamAggregatesVWAPResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

// StreamAggregatesVWAPServiceV1_ServiceDesc is the grpc.ServiceDesc for StreamAggregatesVWAPServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamAggregatesVWAPServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaikosdk.StreamAggregatesVWAPServiceV1",
	HandlerType: (*StreamAggregatesVWAPServiceV1Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _StreamAggregatesVWAPServiceV1_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk/sdk.proto",
}

// StreamMarketUpdateServiceV1Client is the client API for StreamMarketUpdateServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamMarketUpdateServiceV1Client interface {
	// Subscribe
	Subscribe(ctx context.Context, in *market_update_v1.StreamMarketUpdateRequestV1, opts ...grpc.CallOption) (StreamMarketUpdateServiceV1_SubscribeClient, error)
}

type streamMarketUpdateServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStreamMarketUpdateServiceV1Client(cc grpc.ClientConnInterface) StreamMarketUpdateServiceV1Client {
	return &streamMarketUpdateServiceV1Client{cc}
}

func (c *streamMarketUpdateServiceV1Client) Subscribe(ctx context.Context, in *market_update_v1.StreamMarketUpdateRequestV1, opts ...grpc.CallOption) (StreamMarketUpdateServiceV1_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamMarketUpdateServiceV1_ServiceDesc.Streams[0], "/kaikosdk.StreamMarketUpdateServiceV1/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamMarketUpdateServiceV1SubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamMarketUpdateServiceV1_SubscribeClient interface {
	Recv() (*market_update_v1.StreamMarketUpdateResponseV1, error)
	grpc.ClientStream
}

type streamMarketUpdateServiceV1SubscribeClient struct {
	grpc.ClientStream
}

func (x *streamMarketUpdateServiceV1SubscribeClient) Recv() (*market_update_v1.StreamMarketUpdateResponseV1, error) {
	m := new(market_update_v1.StreamMarketUpdateResponseV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamMarketUpdateServiceV1Server is the server API for StreamMarketUpdateServiceV1 service.
// All implementations must embed UnimplementedStreamMarketUpdateServiceV1Server
// for forward compatibility
type StreamMarketUpdateServiceV1Server interface {
	// Subscribe
	Subscribe(*market_update_v1.StreamMarketUpdateRequestV1, StreamMarketUpdateServiceV1_SubscribeServer) error
	mustEmbedUnimplementedStreamMarketUpdateServiceV1Server()
}

// UnimplementedStreamMarketUpdateServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedStreamMarketUpdateServiceV1Server struct {
}

func (UnimplementedStreamMarketUpdateServiceV1Server) Subscribe(*market_update_v1.StreamMarketUpdateRequestV1, StreamMarketUpdateServiceV1_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedStreamMarketUpdateServiceV1Server) mustEmbedUnimplementedStreamMarketUpdateServiceV1Server() {
}

// UnsafeStreamMarketUpdateServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamMarketUpdateServiceV1Server will
// result in compilation errors.
type UnsafeStreamMarketUpdateServiceV1Server interface {
	mustEmbedUnimplementedStreamMarketUpdateServiceV1Server()
}

func RegisterStreamMarketUpdateServiceV1Server(s grpc.ServiceRegistrar, srv StreamMarketUpdateServiceV1Server) {
	s.RegisterService(&StreamMarketUpdateServiceV1_ServiceDesc, srv)
}

func _StreamMarketUpdateServiceV1_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(market_update_v1.StreamMarketUpdateRequestV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamMarketUpdateServiceV1Server).Subscribe(m, &streamMarketUpdateServiceV1SubscribeServer{stream})
}

type StreamMarketUpdateServiceV1_SubscribeServer interface {
	Send(*market_update_v1.StreamMarketUpdateResponseV1) error
	grpc.ServerStream
}

type streamMarketUpdateServiceV1SubscribeServer struct {
	grpc.ServerStream
}

func (x *streamMarketUpdateServiceV1SubscribeServer) Send(m *market_update_v1.StreamMarketUpdateResponseV1) error {
	return x.ServerStream.SendMsg(m)
}

// StreamMarketUpdateServiceV1_ServiceDesc is the grpc.ServiceDesc for StreamMarketUpdateServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamMarketUpdateServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaikosdk.StreamMarketUpdateServiceV1",
	HandlerType: (*StreamMarketUpdateServiceV1Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _StreamMarketUpdateServiceV1_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk/sdk.proto",
}
