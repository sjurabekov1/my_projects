// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: fix_service.proto

package integration_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FixServiceClient is the client API for FixService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FixServiceClient interface {
	GetTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTypesResponse, error)
	GetInstrumentsByType(ctx context.Context, in *GetInstrumentsByTypeRequest, opts ...grpc.CallOption) (*GetInstrumentsResponse, error)
	GetUserAccounts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetUserAccountsResponse, error)
	PlaceOrder(ctx context.Context, in *FixOrderParameters, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	GetAccountSummary(ctx context.Context, in *GetAccountSummaryRequest, opts ...grpc.CallOption) (*FixAccountSummary, error)
	GetGroups(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	GetInstrumentByExchange(ctx context.Context, in *GetInstrumentsByExchangeRequest, opts ...grpc.CallOption) (*GetInstrumentsResponse, error)
	GetExchanges(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetExchangesResponse, error)
	GetInstrumentByGroup(ctx context.Context, in *GetInstrumentsByGroupRequest, opts ...grpc.CallOption) (*GetInstrumentsResponse, error)
	// rpc GetInstruments(google.protobuf.Empty) returns (stream GetInstrumentsResponse) {} // sending via stream because of max limit 4mb
	GetOHLC(ctx context.Context, in *GetOHLCRequest, opts ...grpc.CallOption) (*GetOHLCResponse, error)
	GetDailyChange(ctx context.Context, in *GetDailyChangeRequest, opts ...grpc.CallOption) (*GetDailyChangeResponse, error)
	GetActiveOrders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetActiveOrdersResponse, error)
	GetTicks(ctx context.Context, in *GetTicksRequest, opts ...grpc.CallOption) (*GetTicksResponse, error)
	GetLastQuote(ctx context.Context, in *GetLastQuoteRequest, opts ...grpc.CallOption) (*GetLastQuoteResponse, error)
	GetLastQuoteMap(ctx context.Context, in *GetLastQuoteRequest, opts ...grpc.CallOption) (*GetLastQuoteMapResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	ModifyOrder(ctx context.Context, in *ModifyOrderRequest, opts ...grpc.CallOption) (*ModifyOrderResponse, error)
	GetInstrumentSchedule(ctx context.Context, in *GetInstrumentScheduleRequest, opts ...grpc.CallOption) (*GetInstrumentScheduleResponse, error)
	GetOrderUpdatesStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (FixService_GetOrderUpdatesStreamClient, error)
	GetCurrencies(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCurrenciesResponse, error)
}

type fixServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFixServiceClient(cc grpc.ClientConnInterface) FixServiceClient {
	return &fixServiceClient{cc}
}

func (c *fixServiceClient) GetTypes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetTypesResponse, error) {
	out := new(GetTypesResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetInstrumentsByType(ctx context.Context, in *GetInstrumentsByTypeRequest, opts ...grpc.CallOption) (*GetInstrumentsResponse, error) {
	out := new(GetInstrumentsResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetInstrumentsByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetUserAccounts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetUserAccountsResponse, error) {
	out := new(GetUserAccountsResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetUserAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) PlaceOrder(ctx context.Context, in *FixOrderParameters, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetAccountSummary(ctx context.Context, in *GetAccountSummaryRequest, opts ...grpc.CallOption) (*FixAccountSummary, error) {
	out := new(FixAccountSummary)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetAccountSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetGroups(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetInstrumentByExchange(ctx context.Context, in *GetInstrumentsByExchangeRequest, opts ...grpc.CallOption) (*GetInstrumentsResponse, error) {
	out := new(GetInstrumentsResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetInstrumentByExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetExchanges(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetExchangesResponse, error) {
	out := new(GetExchangesResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetExchanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetInstrumentByGroup(ctx context.Context, in *GetInstrumentsByGroupRequest, opts ...grpc.CallOption) (*GetInstrumentsResponse, error) {
	out := new(GetInstrumentsResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetInstrumentByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetOHLC(ctx context.Context, in *GetOHLCRequest, opts ...grpc.CallOption) (*GetOHLCResponse, error) {
	out := new(GetOHLCResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetOHLC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetDailyChange(ctx context.Context, in *GetDailyChangeRequest, opts ...grpc.CallOption) (*GetDailyChangeResponse, error) {
	out := new(GetDailyChangeResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetDailyChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetActiveOrders(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetActiveOrdersResponse, error) {
	out := new(GetActiveOrdersResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetActiveOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetTicks(ctx context.Context, in *GetTicksRequest, opts ...grpc.CallOption) (*GetTicksResponse, error) {
	out := new(GetTicksResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetTicks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetLastQuote(ctx context.Context, in *GetLastQuoteRequest, opts ...grpc.CallOption) (*GetLastQuoteResponse, error) {
	out := new(GetLastQuoteResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetLastQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetLastQuoteMap(ctx context.Context, in *GetLastQuoteRequest, opts ...grpc.CallOption) (*GetLastQuoteMapResponse, error) {
	out := new(GetLastQuoteMapResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetLastQuoteMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) ModifyOrder(ctx context.Context, in *ModifyOrderRequest, opts ...grpc.CallOption) (*ModifyOrderResponse, error) {
	out := new(ModifyOrderResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/ModifyOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetInstrumentSchedule(ctx context.Context, in *GetInstrumentScheduleRequest, opts ...grpc.CallOption) (*GetInstrumentScheduleResponse, error) {
	out := new(GetInstrumentScheduleResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetInstrumentSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixServiceClient) GetOrderUpdatesStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (FixService_GetOrderUpdatesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FixService_ServiceDesc.Streams[0], "/integration_service.FixService/GetOrderUpdatesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fixServiceGetOrderUpdatesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FixService_GetOrderUpdatesStreamClient interface {
	Recv() (*GetOrderUpdatesStreamResponse, error)
	grpc.ClientStream
}

type fixServiceGetOrderUpdatesStreamClient struct {
	grpc.ClientStream
}

func (x *fixServiceGetOrderUpdatesStreamClient) Recv() (*GetOrderUpdatesStreamResponse, error) {
	m := new(GetOrderUpdatesStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fixServiceClient) GetCurrencies(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetCurrenciesResponse, error) {
	out := new(GetCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/integration_service.FixService/GetCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FixServiceServer is the server API for FixService service.
// All implementations must embed UnimplementedFixServiceServer
// for forward compatibility
type FixServiceServer interface {
	GetTypes(context.Context, *empty.Empty) (*GetTypesResponse, error)
	GetInstrumentsByType(context.Context, *GetInstrumentsByTypeRequest) (*GetInstrumentsResponse, error)
	GetUserAccounts(context.Context, *empty.Empty) (*GetUserAccountsResponse, error)
	PlaceOrder(context.Context, *FixOrderParameters) (*PlaceOrderResponse, error)
	GetAccountSummary(context.Context, *GetAccountSummaryRequest) (*FixAccountSummary, error)
	GetGroups(context.Context, *empty.Empty) (*GetGroupsResponse, error)
	GetInstrumentByExchange(context.Context, *GetInstrumentsByExchangeRequest) (*GetInstrumentsResponse, error)
	GetExchanges(context.Context, *empty.Empty) (*GetExchangesResponse, error)
	GetInstrumentByGroup(context.Context, *GetInstrumentsByGroupRequest) (*GetInstrumentsResponse, error)
	// rpc GetInstruments(google.protobuf.Empty) returns (stream GetInstrumentsResponse) {} // sending via stream because of max limit 4mb
	GetOHLC(context.Context, *GetOHLCRequest) (*GetOHLCResponse, error)
	GetDailyChange(context.Context, *GetDailyChangeRequest) (*GetDailyChangeResponse, error)
	GetActiveOrders(context.Context, *empty.Empty) (*GetActiveOrdersResponse, error)
	GetTicks(context.Context, *GetTicksRequest) (*GetTicksResponse, error)
	GetLastQuote(context.Context, *GetLastQuoteRequest) (*GetLastQuoteResponse, error)
	GetLastQuoteMap(context.Context, *GetLastQuoteRequest) (*GetLastQuoteMapResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	ModifyOrder(context.Context, *ModifyOrderRequest) (*ModifyOrderResponse, error)
	GetInstrumentSchedule(context.Context, *GetInstrumentScheduleRequest) (*GetInstrumentScheduleResponse, error)
	GetOrderUpdatesStream(*empty.Empty, FixService_GetOrderUpdatesStreamServer) error
	GetCurrencies(context.Context, *empty.Empty) (*GetCurrenciesResponse, error)
	mustEmbedUnimplementedFixServiceServer()
}

// UnimplementedFixServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFixServiceServer struct {
}

func (UnimplementedFixServiceServer) GetTypes(context.Context, *empty.Empty) (*GetTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypes not implemented")
}
func (UnimplementedFixServiceServer) GetInstrumentsByType(context.Context, *GetInstrumentsByTypeRequest) (*GetInstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentsByType not implemented")
}
func (UnimplementedFixServiceServer) GetUserAccounts(context.Context, *empty.Empty) (*GetUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccounts not implemented")
}
func (UnimplementedFixServiceServer) PlaceOrder(context.Context, *FixOrderParameters) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedFixServiceServer) GetAccountSummary(context.Context, *GetAccountSummaryRequest) (*FixAccountSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountSummary not implemented")
}
func (UnimplementedFixServiceServer) GetGroups(context.Context, *empty.Empty) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedFixServiceServer) GetInstrumentByExchange(context.Context, *GetInstrumentsByExchangeRequest) (*GetInstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentByExchange not implemented")
}
func (UnimplementedFixServiceServer) GetExchanges(context.Context, *empty.Empty) (*GetExchangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchanges not implemented")
}
func (UnimplementedFixServiceServer) GetInstrumentByGroup(context.Context, *GetInstrumentsByGroupRequest) (*GetInstrumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentByGroup not implemented")
}
func (UnimplementedFixServiceServer) GetOHLC(context.Context, *GetOHLCRequest) (*GetOHLCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOHLC not implemented")
}
func (UnimplementedFixServiceServer) GetDailyChange(context.Context, *GetDailyChangeRequest) (*GetDailyChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyChange not implemented")
}
func (UnimplementedFixServiceServer) GetActiveOrders(context.Context, *empty.Empty) (*GetActiveOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveOrders not implemented")
}
func (UnimplementedFixServiceServer) GetTicks(context.Context, *GetTicksRequest) (*GetTicksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicks not implemented")
}
func (UnimplementedFixServiceServer) GetLastQuote(context.Context, *GetLastQuoteRequest) (*GetLastQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastQuote not implemented")
}
func (UnimplementedFixServiceServer) GetLastQuoteMap(context.Context, *GetLastQuoteRequest) (*GetLastQuoteMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastQuoteMap not implemented")
}
func (UnimplementedFixServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedFixServiceServer) ModifyOrder(context.Context, *ModifyOrderRequest) (*ModifyOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyOrder not implemented")
}
func (UnimplementedFixServiceServer) GetInstrumentSchedule(context.Context, *GetInstrumentScheduleRequest) (*GetInstrumentScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentSchedule not implemented")
}
func (UnimplementedFixServiceServer) GetOrderUpdatesStream(*empty.Empty, FixService_GetOrderUpdatesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrderUpdatesStream not implemented")
}
func (UnimplementedFixServiceServer) GetCurrencies(context.Context, *empty.Empty) (*GetCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
}
func (UnimplementedFixServiceServer) mustEmbedUnimplementedFixServiceServer() {}

// UnsafeFixServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FixServiceServer will
// result in compilation errors.
type UnsafeFixServiceServer interface {
	mustEmbedUnimplementedFixServiceServer()
}

func RegisterFixServiceServer(s grpc.ServiceRegistrar, srv FixServiceServer) {
	s.RegisterService(&FixService_ServiceDesc, srv)
}

func _FixService_GetTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetTypes(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetInstrumentsByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstrumentsByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetInstrumentsByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetInstrumentsByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetInstrumentsByType(ctx, req.(*GetInstrumentsByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetUserAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetUserAccounts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixOrderParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).PlaceOrder(ctx, req.(*FixOrderParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetAccountSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetAccountSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetAccountSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetAccountSummary(ctx, req.(*GetAccountSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetGroups(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetInstrumentByExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstrumentsByExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetInstrumentByExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetInstrumentByExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetInstrumentByExchange(ctx, req.(*GetInstrumentsByExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetExchanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetExchanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetExchanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetExchanges(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetInstrumentByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstrumentsByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetInstrumentByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetInstrumentByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetInstrumentByGroup(ctx, req.(*GetInstrumentsByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetOHLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOHLCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetOHLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetOHLC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetOHLC(ctx, req.(*GetOHLCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetDailyChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDailyChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetDailyChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetDailyChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetDailyChange(ctx, req.(*GetDailyChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetActiveOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetActiveOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetActiveOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetActiveOrders(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetTicks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetTicks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetTicks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetTicks(ctx, req.(*GetTicksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetLastQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetLastQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetLastQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetLastQuote(ctx, req.(*GetLastQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetLastQuoteMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetLastQuoteMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetLastQuoteMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetLastQuoteMap(ctx, req.(*GetLastQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_ModifyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).ModifyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/ModifyOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).ModifyOrder(ctx, req.(*ModifyOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetInstrumentSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstrumentScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetInstrumentSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetInstrumentSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetInstrumentSchedule(ctx, req.(*GetInstrumentScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixService_GetOrderUpdatesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FixServiceServer).GetOrderUpdatesStream(m, &fixServiceGetOrderUpdatesStreamServer{stream})
}

type FixService_GetOrderUpdatesStreamServer interface {
	Send(*GetOrderUpdatesStreamResponse) error
	grpc.ServerStream
}

type fixServiceGetOrderUpdatesStreamServer struct {
	grpc.ServerStream
}

func (x *fixServiceGetOrderUpdatesStreamServer) Send(m *GetOrderUpdatesStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FixService_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixServiceServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/integration_service.FixService/GetCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixServiceServer).GetCurrencies(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FixService_ServiceDesc is the grpc.ServiceDesc for FixService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FixService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "integration_service.FixService",
	HandlerType: (*FixServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTypes",
			Handler:    _FixService_GetTypes_Handler,
		},
		{
			MethodName: "GetInstrumentsByType",
			Handler:    _FixService_GetInstrumentsByType_Handler,
		},
		{
			MethodName: "GetUserAccounts",
			Handler:    _FixService_GetUserAccounts_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _FixService_PlaceOrder_Handler,
		},
		{
			MethodName: "GetAccountSummary",
			Handler:    _FixService_GetAccountSummary_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _FixService_GetGroups_Handler,
		},
		{
			MethodName: "GetInstrumentByExchange",
			Handler:    _FixService_GetInstrumentByExchange_Handler,
		},
		{
			MethodName: "GetExchanges",
			Handler:    _FixService_GetExchanges_Handler,
		},
		{
			MethodName: "GetInstrumentByGroup",
			Handler:    _FixService_GetInstrumentByGroup_Handler,
		},
		{
			MethodName: "GetOHLC",
			Handler:    _FixService_GetOHLC_Handler,
		},
		{
			MethodName: "GetDailyChange",
			Handler:    _FixService_GetDailyChange_Handler,
		},
		{
			MethodName: "GetActiveOrders",
			Handler:    _FixService_GetActiveOrders_Handler,
		},
		{
			MethodName: "GetTicks",
			Handler:    _FixService_GetTicks_Handler,
		},
		{
			MethodName: "GetLastQuote",
			Handler:    _FixService_GetLastQuote_Handler,
		},
		{
			MethodName: "GetLastQuoteMap",
			Handler:    _FixService_GetLastQuoteMap_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _FixService_GetOrder_Handler,
		},
		{
			MethodName: "ModifyOrder",
			Handler:    _FixService_ModifyOrder_Handler,
		},
		{
			MethodName: "GetInstrumentSchedule",
			Handler:    _FixService_GetInstrumentSchedule_Handler,
		},
		{
			MethodName: "GetCurrencies",
			Handler:    _FixService_GetCurrencies_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrderUpdatesStream",
			Handler:       _FixService_GetOrderUpdatesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fix_service.proto",
}
