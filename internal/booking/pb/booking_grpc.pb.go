// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/booking/pb/booking.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	ListHistory(ctx context.Context, in *ListHistoryRequest, opts ...grpc.CallOption) (*ListHistoryResponse, error)
	SendRequest(ctx context.Context, in *SendRequestRequest, opts ...grpc.CallOption) (*SendRequestResponse, error)
	// passenger
	CreateRequest(ctx context.Context, in *CreateRequestRequest, opts ...grpc.CallOption) (*CreateRequestResponse, error)
	CloseRequest(ctx context.Context, in *CloseRequestRequest, opts ...grpc.CallOption) (*CloseRequestResponse, error)
	GetResponse(ctx context.Context, in *GetResponseRequest, opts ...grpc.CallOption) (*GetResponseResponse, error)
	// driver
	AcceptRequest(ctx context.Context, in *AcceptRequestRequest, opts ...grpc.CallOption) (*AcceptRequestResponse, error)
	RejectRequest(ctx context.Context, in *RejectRequestRequest, opts ...grpc.CallOption) (*RejectRequestResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) ListHistory(ctx context.Context, in *ListHistoryRequest, opts ...grpc.CallOption) (*ListHistoryResponse, error) {
	out := new(ListHistoryResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/ListHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) SendRequest(ctx context.Context, in *SendRequestRequest, opts ...grpc.CallOption) (*SendRequestResponse, error) {
	out := new(SendRequestResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateRequest(ctx context.Context, in *CreateRequestRequest, opts ...grpc.CallOption) (*CreateRequestResponse, error) {
	out := new(CreateRequestResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/CreateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CloseRequest(ctx context.Context, in *CloseRequestRequest, opts ...grpc.CallOption) (*CloseRequestResponse, error) {
	out := new(CloseRequestResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/CloseRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetResponse(ctx context.Context, in *GetResponseRequest, opts ...grpc.CallOption) (*GetResponseResponse, error) {
	out := new(GetResponseResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/GetResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) AcceptRequest(ctx context.Context, in *AcceptRequestRequest, opts ...grpc.CallOption) (*AcceptRequestResponse, error) {
	out := new(AcceptRequestResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/AcceptRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) RejectRequest(ctx context.Context, in *RejectRequestRequest, opts ...grpc.CallOption) (*RejectRequestResponse, error) {
	out := new(RejectRequestResponse)
	err := c.cc.Invoke(ctx, "/passenger.BookingService/RejectRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	ListHistory(context.Context, *ListHistoryRequest) (*ListHistoryResponse, error)
	SendRequest(context.Context, *SendRequestRequest) (*SendRequestResponse, error)
	// passenger
	CreateRequest(context.Context, *CreateRequestRequest) (*CreateRequestResponse, error)
	CloseRequest(context.Context, *CloseRequestRequest) (*CloseRequestResponse, error)
	GetResponse(context.Context, *GetResponseRequest) (*GetResponseResponse, error)
	// driver
	AcceptRequest(context.Context, *AcceptRequestRequest) (*AcceptRequestResponse, error)
	RejectRequest(context.Context, *RejectRequestRequest) (*RejectRequestResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) ListHistory(context.Context, *ListHistoryRequest) (*ListHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHistory not implemented")
}
func (UnimplementedBookingServiceServer) SendRequest(context.Context, *SendRequestRequest) (*SendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedBookingServiceServer) CreateRequest(context.Context, *CreateRequestRequest) (*CreateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequest not implemented")
}
func (UnimplementedBookingServiceServer) CloseRequest(context.Context, *CloseRequestRequest) (*CloseRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseRequest not implemented")
}
func (UnimplementedBookingServiceServer) GetResponse(context.Context, *GetResponseRequest) (*GetResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResponse not implemented")
}
func (UnimplementedBookingServiceServer) AcceptRequest(context.Context, *AcceptRequestRequest) (*AcceptRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRequest not implemented")
}
func (UnimplementedBookingServiceServer) RejectRequest(context.Context, *RejectRequestRequest) (*RejectRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectRequest not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_ListHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ListHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/ListHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ListHistory(ctx, req.(*ListHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).SendRequest(ctx, req.(*SendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/CreateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateRequest(ctx, req.(*CreateRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CloseRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CloseRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/CloseRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CloseRequest(ctx, req.(*CloseRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/GetResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetResponse(ctx, req.(*GetResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_AcceptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).AcceptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/AcceptRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).AcceptRequest(ctx, req.(*AcceptRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_RejectRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).RejectRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.BookingService/RejectRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).RejectRequest(ctx, req.(*RejectRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "passenger.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHistory",
			Handler:    _BookingService_ListHistory_Handler,
		},
		{
			MethodName: "SendRequest",
			Handler:    _BookingService_SendRequest_Handler,
		},
		{
			MethodName: "CreateRequest",
			Handler:    _BookingService_CreateRequest_Handler,
		},
		{
			MethodName: "CloseRequest",
			Handler:    _BookingService_CloseRequest_Handler,
		},
		{
			MethodName: "GetResponse",
			Handler:    _BookingService_GetResponse_Handler,
		},
		{
			MethodName: "AcceptRequest",
			Handler:    _BookingService_AcceptRequest_Handler,
		},
		{
			MethodName: "RejectRequest",
			Handler:    _BookingService_RejectRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/booking/pb/booking.proto",
}
