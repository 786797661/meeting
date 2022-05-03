// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: meeting/v1/meeting.proto

package v1

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

// MeetingClient is the client API for Meeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeetingClient interface {
	Create(ctx context.Context, in *MeetingRequest, opts ...grpc.CallOption) (*MeetingReploy, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReploy, error)
}

type meetingClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingClient(cc grpc.ClientConnInterface) MeetingClient {
	return &meetingClient{cc}
}

func (c *meetingClient) Create(ctx context.Context, in *MeetingRequest, opts ...grpc.CallOption) (*MeetingReploy, error) {
	out := new(MeetingReploy)
	err := c.cc.Invoke(ctx, "/meeting.v1.Meeting/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReploy, error) {
	out := new(RegisterReploy)
	err := c.cc.Invoke(ctx, "/meeting.v1.Meeting/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingServer is the server API for Meeting service.
// All implementations must embed UnimplementedMeetingServer
// for forward compatibility
type MeetingServer interface {
	Create(context.Context, *MeetingRequest) (*MeetingReploy, error)
	Register(context.Context, *RegisterRequest) (*RegisterReploy, error)
	mustEmbedUnimplementedMeetingServer()
}

// UnimplementedMeetingServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingServer struct {
}

func (UnimplementedMeetingServer) Create(context.Context, *MeetingRequest) (*MeetingReploy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMeetingServer) Register(context.Context, *RegisterRequest) (*RegisterReploy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMeetingServer) mustEmbedUnimplementedMeetingServer() {}

// UnsafeMeetingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingServer will
// result in compilation errors.
type UnsafeMeetingServer interface {
	mustEmbedUnimplementedMeetingServer()
}

func RegisterMeetingServer(s grpc.ServiceRegistrar, srv MeetingServer) {
	s.RegisterService(&Meeting_ServiceDesc, srv)
}

func _Meeting_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meeting.v1.Meeting/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServer).Create(ctx, req.(*MeetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meeting_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meeting.v1.Meeting/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Meeting_ServiceDesc is the grpc.ServiceDesc for Meeting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Meeting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meeting.v1.Meeting",
	HandlerType: (*MeetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Meeting_Create_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Meeting_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meeting/v1/meeting.proto",
}
