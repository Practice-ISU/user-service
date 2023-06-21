// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: discovery-ping.proto

package ping

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

// DiscoveryPingClient is the client API for DiscoveryPing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryPingClient interface {
	// Method to check availability of a registered service.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type discoveryPingClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryPingClient(cc grpc.ClientConnInterface) DiscoveryPingClient {
	return &discoveryPingClient{cc}
}

func (c *discoveryPingClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/discovery_ping.DiscoveryPing/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryPingServer is the server API for DiscoveryPing service.
// All implementations must embed UnimplementedDiscoveryPingServer
// for forward compatibility
type DiscoveryPingServer interface {
	// Method to check availability of a registered service.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedDiscoveryPingServer()
}

// UnimplementedDiscoveryPingServer must be embedded to have forward compatible implementations.
type UnimplementedDiscoveryPingServer struct {
}

func (UnimplementedDiscoveryPingServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDiscoveryPingServer) mustEmbedUnimplementedDiscoveryPingServer() {}

// UnsafeDiscoveryPingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryPingServer will
// result in compilation errors.
type UnsafeDiscoveryPingServer interface {
	mustEmbedUnimplementedDiscoveryPingServer()
}

func RegisterDiscoveryPingServer(s grpc.ServiceRegistrar, srv DiscoveryPingServer) {
	s.RegisterService(&DiscoveryPing_ServiceDesc, srv)
}

func _DiscoveryPing_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryPingServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery_ping.DiscoveryPing/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryPingServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscoveryPing_ServiceDesc is the grpc.ServiceDesc for DiscoveryPing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveryPing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discovery_ping.DiscoveryPing",
	HandlerType: (*DiscoveryPingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DiscoveryPing_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery-ping.proto",
}
