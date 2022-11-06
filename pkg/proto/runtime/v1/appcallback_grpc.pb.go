// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: daprd/proto/runtime/v1/appcallback.proto

package runtime

import (
	context "context"
	v1 "github.com/malijoe/daprdan/pkg/proto/components/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AppCallbackClient is the client API for AppCallback service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppCallbackClient interface {
	// Invokes service method with InvokeRequest
	OnInvoke(ctx context.Context, in *v1.InvokeRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error)
	// Lists all topics subscribed by this app.
	ListTopicSubscriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TopicEventResponse, error)
	// Subscribes events from Pubsub
	OnTopicEvent(ctx context.Context, in *TopicEventRequest, opts ...grpc.CallOption) (*TopicEventResponse, error)
	// Lists all input bindings subscribed by this app.
	ListInputBindings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInputBindingsResponse, error)
	// Listens to events from the input bindings
	//
	// User application can save the states or send the events to the output
	// bindings optionally by returning BindingEventResponse.
	OnBindingEvent(ctx context.Context, in *BindingEventRequest, opts ...grpc.CallOption) (*BindingEventResponse, error)
}

type appCallbackClient struct {
	cc grpc.ClientConnInterface
}

func NewAppCallbackClient(cc grpc.ClientConnInterface) AppCallbackClient {
	return &appCallbackClient{cc}
}

func (c *appCallbackClient) OnInvoke(ctx context.Context, in *v1.InvokeRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error) {
	out := new(v1.InvokeResponse)
	err := c.cc.Invoke(ctx, "/daprd.proto.runtime.v1.AppCallback/OnInvoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appCallbackClient) ListTopicSubscriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TopicEventResponse, error) {
	out := new(TopicEventResponse)
	err := c.cc.Invoke(ctx, "/daprd.proto.runtime.v1.AppCallback/ListTopicSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appCallbackClient) OnTopicEvent(ctx context.Context, in *TopicEventRequest, opts ...grpc.CallOption) (*TopicEventResponse, error) {
	out := new(TopicEventResponse)
	err := c.cc.Invoke(ctx, "/daprd.proto.runtime.v1.AppCallback/OnTopicEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appCallbackClient) ListInputBindings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListInputBindingsResponse, error) {
	out := new(ListInputBindingsResponse)
	err := c.cc.Invoke(ctx, "/daprd.proto.runtime.v1.AppCallback/ListInputBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appCallbackClient) OnBindingEvent(ctx context.Context, in *BindingEventRequest, opts ...grpc.CallOption) (*BindingEventResponse, error) {
	out := new(BindingEventResponse)
	err := c.cc.Invoke(ctx, "/daprd.proto.runtime.v1.AppCallback/OnBindingEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppCallbackServer is the server API for AppCallback service.
// All implementations should embed UnimplementedAppCallbackServer
// for forward compatibility
type AppCallbackServer interface {
	// Invokes service method with InvokeRequest
	OnInvoke(context.Context, *v1.InvokeRequest) (*v1.InvokeResponse, error)
	// Lists all topics subscribed by this app.
	ListTopicSubscriptions(context.Context, *emptypb.Empty) (*TopicEventResponse, error)
	// Subscribes events from Pubsub
	OnTopicEvent(context.Context, *TopicEventRequest) (*TopicEventResponse, error)
	// Lists all input bindings subscribed by this app.
	ListInputBindings(context.Context, *emptypb.Empty) (*ListInputBindingsResponse, error)
	// Listens to events from the input bindings
	//
	// User application can save the states or send the events to the output
	// bindings optionally by returning BindingEventResponse.
	OnBindingEvent(context.Context, *BindingEventRequest) (*BindingEventResponse, error)
}

// UnimplementedAppCallbackServer should be embedded to have forward compatible implementations.
type UnimplementedAppCallbackServer struct {
}

func (UnimplementedAppCallbackServer) OnInvoke(context.Context, *v1.InvokeRequest) (*v1.InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnInvoke not implemented")
}
func (UnimplementedAppCallbackServer) ListTopicSubscriptions(context.Context, *emptypb.Empty) (*TopicEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopicSubscriptions not implemented")
}
func (UnimplementedAppCallbackServer) OnTopicEvent(context.Context, *TopicEventRequest) (*TopicEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnTopicEvent not implemented")
}
func (UnimplementedAppCallbackServer) ListInputBindings(context.Context, *emptypb.Empty) (*ListInputBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInputBindings not implemented")
}
func (UnimplementedAppCallbackServer) OnBindingEvent(context.Context, *BindingEventRequest) (*BindingEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnBindingEvent not implemented")
}

// UnsafeAppCallbackServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppCallbackServer will
// result in compilation errors.
type UnsafeAppCallbackServer interface {
	mustEmbedUnimplementedAppCallbackServer()
}

func RegisterAppCallbackServer(s grpc.ServiceRegistrar, srv AppCallbackServer) {
	s.RegisterService(&AppCallback_ServiceDesc, srv)
}

func _AppCallback_OnInvoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCallbackServer).OnInvoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprd.proto.runtime.v1.AppCallback/OnInvoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCallbackServer).OnInvoke(ctx, req.(*v1.InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppCallback_ListTopicSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCallbackServer).ListTopicSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprd.proto.runtime.v1.AppCallback/ListTopicSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCallbackServer).ListTopicSubscriptions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppCallback_OnTopicEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCallbackServer).OnTopicEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprd.proto.runtime.v1.AppCallback/OnTopicEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCallbackServer).OnTopicEvent(ctx, req.(*TopicEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppCallback_ListInputBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCallbackServer).ListInputBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprd.proto.runtime.v1.AppCallback/ListInputBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCallbackServer).ListInputBindings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppCallback_OnBindingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindingEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCallbackServer).OnBindingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprd.proto.runtime.v1.AppCallback/OnBindingEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCallbackServer).OnBindingEvent(ctx, req.(*BindingEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppCallback_ServiceDesc is the grpc.ServiceDesc for AppCallback service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppCallback_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daprd.proto.runtime.v1.AppCallback",
	HandlerType: (*AppCallbackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnInvoke",
			Handler:    _AppCallback_OnInvoke_Handler,
		},
		{
			MethodName: "ListTopicSubscriptions",
			Handler:    _AppCallback_ListTopicSubscriptions_Handler,
		},
		{
			MethodName: "OnTopicEvent",
			Handler:    _AppCallback_OnTopicEvent_Handler,
		},
		{
			MethodName: "ListInputBindings",
			Handler:    _AppCallback_ListInputBindings_Handler,
		},
		{
			MethodName: "OnBindingEvent",
			Handler:    _AppCallback_OnBindingEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daprd/proto/runtime/v1/appcallback.proto",
}

// AppCallbackHealthCheckClient is the client API for AppCallbackHealthCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppCallbackHealthCheckClient interface {
	// Health check.
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type appCallbackHealthCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewAppCallbackHealthCheckClient(cc grpc.ClientConnInterface) AppCallbackHealthCheckClient {
	return &appCallbackHealthCheckClient{cc}
}

func (c *appCallbackHealthCheckClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/daprd.proto.runtime.v1.AppCallbackHealthCheck/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppCallbackHealthCheckServer is the server API for AppCallbackHealthCheck service.
// All implementations should embed UnimplementedAppCallbackHealthCheckServer
// for forward compatibility
type AppCallbackHealthCheckServer interface {
	// Health check.
	HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
}

// UnimplementedAppCallbackHealthCheckServer should be embedded to have forward compatible implementations.
type UnimplementedAppCallbackHealthCheckServer struct {
}

func (UnimplementedAppCallbackHealthCheckServer) HealthCheck(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}

// UnsafeAppCallbackHealthCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppCallbackHealthCheckServer will
// result in compilation errors.
type UnsafeAppCallbackHealthCheckServer interface {
	mustEmbedUnimplementedAppCallbackHealthCheckServer()
}

func RegisterAppCallbackHealthCheckServer(s grpc.ServiceRegistrar, srv AppCallbackHealthCheckServer) {
	s.RegisterService(&AppCallbackHealthCheck_ServiceDesc, srv)
}

func _AppCallbackHealthCheck_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppCallbackHealthCheckServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daprd.proto.runtime.v1.AppCallbackHealthCheck/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppCallbackHealthCheckServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AppCallbackHealthCheck_ServiceDesc is the grpc.ServiceDesc for AppCallbackHealthCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppCallbackHealthCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daprd.proto.runtime.v1.AppCallbackHealthCheck",
	HandlerType: (*AppCallbackHealthCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _AppCallbackHealthCheck_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daprd/proto/runtime/v1/appcallback.proto",
}
