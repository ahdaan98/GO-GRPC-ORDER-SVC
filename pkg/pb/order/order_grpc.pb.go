// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: order.proto

package order

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	OrderItemsFromCart(ctx context.Context, in *OrderItemsFromCartRequest, opts ...grpc.CallOption) (*OrderItemsFromCartResponse, error)
	GetOrderDetails(ctx context.Context, in *GetOrderDetailsRequest, opts ...grpc.CallOption) (*GetOrderDetailsResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) OrderItemsFromCart(ctx context.Context, in *OrderItemsFromCartRequest, opts ...grpc.CallOption) (*OrderItemsFromCartResponse, error) {
	out := new(OrderItemsFromCartResponse)
	err := c.cc.Invoke(ctx, "/order.Order/OrderItemsFromCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderDetails(ctx context.Context, in *GetOrderDetailsRequest, opts ...grpc.CallOption) (*GetOrderDetailsResponse, error) {
	out := new(GetOrderDetailsResponse)
	err := c.cc.Invoke(ctx, "/order.Order/GetOrderDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	OrderItemsFromCart(context.Context, *OrderItemsFromCartRequest) (*OrderItemsFromCartResponse, error)
	GetOrderDetails(context.Context, *GetOrderDetailsRequest) (*GetOrderDetailsResponse, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) OrderItemsFromCart(context.Context, *OrderItemsFromCartRequest) (*OrderItemsFromCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderItemsFromCart not implemented")
}
func (UnimplementedOrderServer) GetOrderDetails(context.Context, *GetOrderDetailsRequest) (*GetOrderDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetails not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_OrderItemsFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemsFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderItemsFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/OrderItemsFromCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderItemsFromCart(ctx, req.(*OrderItemsFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/GetOrderDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderDetails(ctx, req.(*GetOrderDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderItemsFromCart",
			Handler:    _Order_OrderItemsFromCart_Handler,
		},
		{
			MethodName: "GetOrderDetails",
			Handler:    _Order_GetOrderDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
