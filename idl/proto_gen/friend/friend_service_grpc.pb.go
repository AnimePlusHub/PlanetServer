// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: friend_service.proto

package friend_service

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

const (
	FriendService_AddGroup_FullMethodName    = "/friend_idl.FriendService/AddGroup"
	FriendService_SearchUsers_FullMethodName = "/friend_idl.FriendService/SearchUsers"
)

// FriendServiceClient is the client API for FriendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendServiceClient interface {
	// 添加分组
	AddGroup(ctx context.Context, in *GroupReq, opts ...grpc.CallOption) (*MsgRsp, error)
	// 查找用户（account或昵称），返回多个用户
	SearchUsers(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (FriendService_SearchUsersClient, error)
}

type friendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendServiceClient(cc grpc.ClientConnInterface) FriendServiceClient {
	return &friendServiceClient{cc}
}

func (c *friendServiceClient) AddGroup(ctx context.Context, in *GroupReq, opts ...grpc.CallOption) (*MsgRsp, error) {
	out := new(MsgRsp)
	err := c.cc.Invoke(ctx, FriendService_AddGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) SearchUsers(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (FriendService_SearchUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &FriendService_ServiceDesc.Streams[0], FriendService_SearchUsers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &friendServiceSearchUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FriendService_SearchUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type friendServiceSearchUsersClient struct {
	grpc.ClientStream
}

func (x *friendServiceSearchUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FriendServiceServer is the server API for FriendService service.
// All implementations must embed UnimplementedFriendServiceServer
// for forward compatibility
type FriendServiceServer interface {
	// 添加分组
	AddGroup(context.Context, *GroupReq) (*MsgRsp, error)
	// 查找用户（account或昵称），返回多个用户
	SearchUsers(*NameReq, FriendService_SearchUsersServer) error
	mustEmbedUnimplementedFriendServiceServer()
}

// UnimplementedFriendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServiceServer struct {
}

func (UnimplementedFriendServiceServer) AddGroup(context.Context, *GroupReq) (*MsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedFriendServiceServer) SearchUsers(*NameReq, FriendService_SearchUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedFriendServiceServer) mustEmbedUnimplementedFriendServiceServer() {}

// UnsafeFriendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServiceServer will
// result in compilation errors.
type UnsafeFriendServiceServer interface {
	mustEmbedUnimplementedFriendServiceServer()
}

func RegisterFriendServiceServer(s grpc.ServiceRegistrar, srv FriendServiceServer) {
	s.RegisterService(&FriendService_ServiceDesc, srv)
}

func _FriendService_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendService_AddGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).AddGroup(ctx, req.(*GroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_SearchUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NameReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FriendServiceServer).SearchUsers(m, &friendServiceSearchUsersServer{stream})
}

type FriendService_SearchUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type friendServiceSearchUsersServer struct {
	grpc.ServerStream
}

func (x *friendServiceSearchUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

// FriendService_ServiceDesc is the grpc.ServiceDesc for FriendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "friend_idl.FriendService",
	HandlerType: (*FriendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGroup",
			Handler:    _FriendService_AddGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchUsers",
			Handler:       _FriendService_SearchUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "friend_service.proto",
}
