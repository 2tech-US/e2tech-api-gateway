// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: internal/location/pb/location.proto

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

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	// unary - synchronous
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	GetAddressList(ctx context.Context, in *GetListAddressRequest, opts ...grpc.CallOption) (*GetListAddressResponse, error)
	CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error)
	UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationResponse, error)
	CreateRequest(ctx context.Context, in *CreateCallCenterRequest, opts ...grpc.CallOption) (*CreateCallCenterRequestResponse, error)
	GetRequest(ctx context.Context, in *GetCallCenterRequest, opts ...grpc.CallOption) (*GetCallCenterRequestResponse, error)
	GetListRequest(ctx context.Context, in *GetListCallCenterRequest, opts ...grpc.CallOption) (*GetListRequestResponse, error)
	GetListRequestByPhone(ctx context.Context, in *GetListCallCenterRequestByPhone, opts ...grpc.CallOption) (*GetListRequestResponse, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/getAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetAddressList(ctx context.Context, in *GetListAddressRequest, opts ...grpc.CallOption) (*GetListAddressResponse, error) {
	out := new(GetListAddressResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/getAddressList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/createAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error) {
	out := new(UpdateAddressResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/updateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationResponse, error) {
	out := new(UpdateLocationResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/updateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) CreateRequest(ctx context.Context, in *CreateCallCenterRequest, opts ...grpc.CallOption) (*CreateCallCenterRequestResponse, error) {
	out := new(CreateCallCenterRequestResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/createRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetRequest(ctx context.Context, in *GetCallCenterRequest, opts ...grpc.CallOption) (*GetCallCenterRequestResponse, error) {
	out := new(GetCallCenterRequestResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/getRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetListRequest(ctx context.Context, in *GetListCallCenterRequest, opts ...grpc.CallOption) (*GetListRequestResponse, error) {
	out := new(GetListRequestResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/getListRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetListRequestByPhone(ctx context.Context, in *GetListCallCenterRequestByPhone, opts ...grpc.CallOption) (*GetListRequestResponse, error) {
	out := new(GetListRequestResponse)
	err := c.cc.Invoke(ctx, "/tech2.microservice.LocationService/getListRequestByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility
type LocationServiceServer interface {
	// unary - synchronous
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	GetAddressList(context.Context, *GetListAddressRequest) (*GetListAddressResponse, error)
	CreateAddress(context.Context, *CreateAddressRequest) (*GetAddressResponse, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error)
	UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error)
	CreateRequest(context.Context, *CreateCallCenterRequest) (*CreateCallCenterRequestResponse, error)
	GetRequest(context.Context, *GetCallCenterRequest) (*GetCallCenterRequestResponse, error)
	GetListRequest(context.Context, *GetListCallCenterRequest) (*GetListRequestResponse, error)
	GetListRequestByPhone(context.Context, *GetListCallCenterRequestByPhone) (*GetListRequestResponse, error)
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (UnimplementedLocationServiceServer) GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedLocationServiceServer) GetAddressList(context.Context, *GetListAddressRequest) (*GetListAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressList not implemented")
}
func (UnimplementedLocationServiceServer) CreateAddress(context.Context, *CreateAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedLocationServiceServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedLocationServiceServer) UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedLocationServiceServer) CreateRequest(context.Context, *CreateCallCenterRequest) (*CreateCallCenterRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequest not implemented")
}
func (UnimplementedLocationServiceServer) GetRequest(context.Context, *GetCallCenterRequest) (*GetCallCenterRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequest not implemented")
}
func (UnimplementedLocationServiceServer) GetListRequest(context.Context, *GetListCallCenterRequest) (*GetListRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListRequest not implemented")
}
func (UnimplementedLocationServiceServer) GetListRequestByPhone(context.Context, *GetListCallCenterRequestByPhone) (*GetListRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListRequestByPhone not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/getAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/getAddressList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetAddressList(ctx, req.(*GetListAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/createAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).CreateAddress(ctx, req.(*CreateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/updateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/updateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).UpdateLocation(ctx, req.(*UpdateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_CreateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCallCenterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).CreateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/createRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).CreateRequest(ctx, req.(*CreateCallCenterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCallCenterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/getRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetRequest(ctx, req.(*GetCallCenterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetListRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCallCenterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetListRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/getListRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetListRequest(ctx, req.(*GetListCallCenterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetListRequestByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListCallCenterRequestByPhone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetListRequestByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tech2.microservice.LocationService/getListRequestByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetListRequestByPhone(ctx, req.(*GetListCallCenterRequestByPhone))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tech2.microservice.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAddress",
			Handler:    _LocationService_GetAddress_Handler,
		},
		{
			MethodName: "getAddressList",
			Handler:    _LocationService_GetAddressList_Handler,
		},
		{
			MethodName: "createAddress",
			Handler:    _LocationService_CreateAddress_Handler,
		},
		{
			MethodName: "updateAddress",
			Handler:    _LocationService_UpdateAddress_Handler,
		},
		{
			MethodName: "updateLocation",
			Handler:    _LocationService_UpdateLocation_Handler,
		},
		{
			MethodName: "createRequest",
			Handler:    _LocationService_CreateRequest_Handler,
		},
		{
			MethodName: "getRequest",
			Handler:    _LocationService_GetRequest_Handler,
		},
		{
			MethodName: "getListRequest",
			Handler:    _LocationService_GetListRequest_Handler,
		},
		{
			MethodName: "getListRequestByPhone",
			Handler:    _LocationService_GetListRequestByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/location/pb/location.proto",
}
