// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.14.0
// source: go-usermgmt-grpc/usermgmt/usermgmt.proto

package go_usermgmt_grpc

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

// UserManagementClient is the client API for UserManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementClient interface {
	CreateNewUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DelUser, opts ...grpc.CallOption) (*UserID, error)
	GetListUsers(ctx context.Context, in *GetUsers, opts ...grpc.CallOption) (*ListUsers, error)
}

type userManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementClient(cc grpc.ClientConnInterface) UserManagementClient {
	return &userManagementClient{cc}
}

func (c *userManagementClient) CreateNewUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagement/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) DeleteUser(ctx context.Context, in *DelUser, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagement/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetListUsers(ctx context.Context, in *GetUsers, opts ...grpc.CallOption) (*ListUsers, error) {
	out := new(ListUsers)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagement/GetListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServer is the server API for UserManagement service.
// All implementations must embed UnimplementedUserManagementServer
// for forward compatibility
type UserManagementServer interface {
	CreateNewUser(context.Context, *NewUser) (*User, error)
	DeleteUser(context.Context, *DelUser) (*UserID, error)
	GetListUsers(context.Context, *GetUsers) (*ListUsers, error)
	mustEmbedUnimplementedUserManagementServer()
}

// UnimplementedUserManagementServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServer struct {
}

func (UnimplementedUserManagementServer) CreateNewUser(context.Context, *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (UnimplementedUserManagementServer) DeleteUser(context.Context, *DelUser) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagementServer) GetListUsers(context.Context, *GetUsers) (*ListUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUsers not implemented")
}
func (UnimplementedUserManagementServer) mustEmbedUnimplementedUserManagementServer() {}

// UnsafeUserManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServer will
// result in compilation errors.
type UnsafeUserManagementServer interface {
	mustEmbedUnimplementedUserManagementServer()
}

func RegisterUserManagementServer(s grpc.ServiceRegistrar, srv UserManagementServer) {
	s.RegisterService(&UserManagement_ServiceDesc, srv)
}

func _UserManagement_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagement/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CreateNewUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagement/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).DeleteUser(ctx, req.(*DelUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagement/GetListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetListUsers(ctx, req.(*GetUsers))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagement_ServiceDesc is the grpc.ServiceDesc for UserManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermgmt.UserManagement",
	HandlerType: (*UserManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewUser",
			Handler:    _UserManagement_CreateNewUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManagement_DeleteUser_Handler,
		},
		{
			MethodName: "GetListUsers",
			Handler:    _UserManagement_GetListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-usermgmt-grpc/usermgmt/usermgmt.proto",
}
