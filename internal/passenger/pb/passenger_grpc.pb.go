// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/passenger/pb/passenger.proto

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

// PassengerServiceClient is the client API for PassengerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PassengerServiceClient interface {
	CreatePassenger(ctx context.Context, in *CreatePassengerRequest, opts ...grpc.CallOption) (*CreatePassengerResponse, error)
	GetPassengerByPhone(ctx context.Context, in *GetPassengerByPhoneRequest, opts ...grpc.CallOption) (*GetPassengerByPhoneResponse, error)
	ListPassengers(ctx context.Context, in *ListPassengersRequest, opts ...grpc.CallOption) (*ListPassengersResponse, error)
	UpdatePassenger(ctx context.Context, in *UpdatePassengerRequest, opts ...grpc.CallOption) (*UpdatePassengerResponse, error)
	DeletePassenger(ctx context.Context, in *DeletePassengerRequest, opts ...grpc.CallOption) (*DeletePassengerResponse, error)
}

type passengerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPassengerServiceClient(cc grpc.ClientConnInterface) PassengerServiceClient {
	return &passengerServiceClient{cc}
}

func (c *passengerServiceClient) CreatePassenger(ctx context.Context, in *CreatePassengerRequest, opts ...grpc.CallOption) (*CreatePassengerResponse, error) {
	out := new(CreatePassengerResponse)
	err := c.cc.Invoke(ctx, "/passenger.PassengerService/CreatePassenger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerServiceClient) GetPassengerByPhone(ctx context.Context, in *GetPassengerByPhoneRequest, opts ...grpc.CallOption) (*GetPassengerByPhoneResponse, error) {
	out := new(GetPassengerByPhoneResponse)
	err := c.cc.Invoke(ctx, "/passenger.PassengerService/GetPassengerByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerServiceClient) ListPassengers(ctx context.Context, in *ListPassengersRequest, opts ...grpc.CallOption) (*ListPassengersResponse, error) {
	out := new(ListPassengersResponse)
	err := c.cc.Invoke(ctx, "/passenger.PassengerService/ListPassengers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerServiceClient) UpdatePassenger(ctx context.Context, in *UpdatePassengerRequest, opts ...grpc.CallOption) (*UpdatePassengerResponse, error) {
	out := new(UpdatePassengerResponse)
	err := c.cc.Invoke(ctx, "/passenger.PassengerService/UpdatePassenger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerServiceClient) DeletePassenger(ctx context.Context, in *DeletePassengerRequest, opts ...grpc.CallOption) (*DeletePassengerResponse, error) {
	out := new(DeletePassengerResponse)
	err := c.cc.Invoke(ctx, "/passenger.PassengerService/DeletePassenger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerServiceServer is the server API for PassengerService service.
// All implementations must embed UnimplementedPassengerServiceServer
// for forward compatibility
type PassengerServiceServer interface {
	CreatePassenger(context.Context, *CreatePassengerRequest) (*CreatePassengerResponse, error)
	GetPassengerByPhone(context.Context, *GetPassengerByPhoneRequest) (*GetPassengerByPhoneResponse, error)
	ListPassengers(context.Context, *ListPassengersRequest) (*ListPassengersResponse, error)
	UpdatePassenger(context.Context, *UpdatePassengerRequest) (*UpdatePassengerResponse, error)
	DeletePassenger(context.Context, *DeletePassengerRequest) (*DeletePassengerResponse, error)
	mustEmbedUnimplementedPassengerServiceServer()
}

// UnimplementedPassengerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPassengerServiceServer struct {
}

func (UnimplementedPassengerServiceServer) CreatePassenger(context.Context, *CreatePassengerRequest) (*CreatePassengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePassenger not implemented")
}
func (UnimplementedPassengerServiceServer) GetPassengerByPhone(context.Context, *GetPassengerByPhoneRequest) (*GetPassengerByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPassengerByPhone not implemented")
}
func (UnimplementedPassengerServiceServer) ListPassengers(context.Context, *ListPassengersRequest) (*ListPassengersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPassengers not implemented")
}
func (UnimplementedPassengerServiceServer) UpdatePassenger(context.Context, *UpdatePassengerRequest) (*UpdatePassengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassenger not implemented")
}
func (UnimplementedPassengerServiceServer) DeletePassenger(context.Context, *DeletePassengerRequest) (*DeletePassengerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassenger not implemented")
}
func (UnimplementedPassengerServiceServer) mustEmbedUnimplementedPassengerServiceServer() {}

// UnsafePassengerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PassengerServiceServer will
// result in compilation errors.
type UnsafePassengerServiceServer interface {
	mustEmbedUnimplementedPassengerServiceServer()
}

func RegisterPassengerServiceServer(s grpc.ServiceRegistrar, srv PassengerServiceServer) {
	s.RegisterService(&PassengerService_ServiceDesc, srv)
}

func _PassengerService_CreatePassenger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePassengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).CreatePassenger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.PassengerService/CreatePassenger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).CreatePassenger(ctx, req.(*CreatePassengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerService_GetPassengerByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPassengerByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).GetPassengerByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.PassengerService/GetPassengerByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).GetPassengerByPhone(ctx, req.(*GetPassengerByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerService_ListPassengers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPassengersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).ListPassengers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.PassengerService/ListPassengers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).ListPassengers(ctx, req.(*ListPassengersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerService_UpdatePassenger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePassengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).UpdatePassenger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.PassengerService/UpdatePassenger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).UpdatePassenger(ctx, req.(*UpdatePassengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerService_DeletePassenger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePassengerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerServiceServer).DeletePassenger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passenger.PassengerService/DeletePassenger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerServiceServer).DeletePassenger(ctx, req.(*DeletePassengerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PassengerService_ServiceDesc is the grpc.ServiceDesc for PassengerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PassengerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "passenger.PassengerService",
	HandlerType: (*PassengerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePassenger",
			Handler:    _PassengerService_CreatePassenger_Handler,
		},
		{
			MethodName: "GetPassengerByPhone",
			Handler:    _PassengerService_GetPassengerByPhone_Handler,
		},
		{
			MethodName: "ListPassengers",
			Handler:    _PassengerService_ListPassengers_Handler,
		},
		{
			MethodName: "UpdatePassenger",
			Handler:    _PassengerService_UpdatePassenger_Handler,
		},
		{
			MethodName: "DeletePassenger",
			Handler:    _PassengerService_DeletePassenger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/passenger/pb/passenger.proto",
}
