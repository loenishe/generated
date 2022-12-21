// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: com/daml/ledger/api/v1/admin/user_management_service.proto

package admin

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

// UserManagementServiceClient is the client API for UserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementServiceClient interface {
	// Create a new user.
	// Errors:
	// - “ALREADY_EXISTS“: if the user already exists
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Get the user data of a specific user or the authenticated user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Delete an existing user and all its rights.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	// List all existing users.
	// Errors:
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	// Grant rights to a user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	GrantUserRights(ctx context.Context, in *GrantUserRightsRequest, opts ...grpc.CallOption) (*GrantUserRightsResponse, error)
	// Revoke rights from a user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	RevokeUserRights(ctx context.Context, in *RevokeUserRightsRequest, opts ...grpc.CallOption) (*RevokeUserRightsResponse, error)
	// List the set of all rights granted to a user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	ListUserRights(ctx context.Context, in *ListUserRightsRequest, opts ...grpc.CallOption) (*ListUserRightsResponse, error)
}

type userManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementServiceClient(cc grpc.ClientConnInterface) UserManagementServiceClient {
	return &userManagementServiceClient{cc}
}

func (c *userManagementServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GrantUserRights(ctx context.Context, in *GrantUserRightsRequest, opts ...grpc.CallOption) (*GrantUserRightsResponse, error) {
	out := new(GrantUserRightsResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/GrantUserRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) RevokeUserRights(ctx context.Context, in *RevokeUserRightsRequest, opts ...grpc.CallOption) (*RevokeUserRightsResponse, error) {
	out := new(RevokeUserRightsResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/RevokeUserRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) ListUserRights(ctx context.Context, in *ListUserRightsRequest, opts ...grpc.CallOption) (*ListUserRightsResponse, error) {
	out := new(ListUserRightsResponse)
	err := c.cc.Invoke(ctx, "/com.daml.ledger.api.v1.admin.UserManagementService/ListUserRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServiceServer is the server API for UserManagementService service.
// All implementations must embed UnimplementedUserManagementServiceServer
// for forward compatibility
type UserManagementServiceServer interface {
	// Create a new user.
	// Errors:
	// - “ALREADY_EXISTS“: if the user already exists
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Get the user data of a specific user or the authenticated user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Delete an existing user and all its rights.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	// List all existing users.
	// Errors:
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// Grant rights to a user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	GrantUserRights(context.Context, *GrantUserRightsRequest) (*GrantUserRightsResponse, error)
	// Revoke rights from a user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	RevokeUserRights(context.Context, *RevokeUserRightsRequest) (*RevokeUserRightsResponse, error)
	// List the set of all rights granted to a user.
	// Errors:
	// - “NOT_FOUND“: if the user doesn't exist
	// - “UNAUTHENTICATED“: if the request does not include a valid access token
	// - “PERMISSION_DENIED“: if the claims in the token are insufficient to perform a given operation
	// - “INVALID_ARGUMENT“: if the payload is malformed or is missing required fields
	ListUserRights(context.Context, *ListUserRightsRequest) (*ListUserRightsResponse, error)
	mustEmbedUnimplementedUserManagementServiceServer()
}

// UnimplementedUserManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServiceServer struct {
}

func (UnimplementedUserManagementServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserManagementServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserManagementServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagementServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserManagementServiceServer) GrantUserRights(context.Context, *GrantUserRightsRequest) (*GrantUserRightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantUserRights not implemented")
}
func (UnimplementedUserManagementServiceServer) RevokeUserRights(context.Context, *RevokeUserRightsRequest) (*RevokeUserRightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUserRights not implemented")
}
func (UnimplementedUserManagementServiceServer) ListUserRights(context.Context, *ListUserRightsRequest) (*ListUserRightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRights not implemented")
}
func (UnimplementedUserManagementServiceServer) mustEmbedUnimplementedUserManagementServiceServer() {}

// UnsafeUserManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServiceServer will
// result in compilation errors.
type UnsafeUserManagementServiceServer interface {
	mustEmbedUnimplementedUserManagementServiceServer()
}

func RegisterUserManagementServiceServer(s grpc.ServiceRegistrar, srv UserManagementServiceServer) {
	s.RegisterService(&UserManagementService_ServiceDesc, srv)
}

func _UserManagementService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GrantUserRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantUserRightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GrantUserRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/GrantUserRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GrantUserRights(ctx, req.(*GrantUserRightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_RevokeUserRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeUserRightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).RevokeUserRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/RevokeUserRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).RevokeUserRights(ctx, req.(*RevokeUserRightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_ListUserRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).ListUserRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.daml.ledger.api.v1.admin.UserManagementService/ListUserRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).ListUserRights(ctx, req.(*ListUserRightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagementService_ServiceDesc is the grpc.ServiceDesc for UserManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.daml.ledger.api.v1.admin.UserManagementService",
	HandlerType: (*UserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserManagementService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserManagementService_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManagementService_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserManagementService_ListUsers_Handler,
		},
		{
			MethodName: "GrantUserRights",
			Handler:    _UserManagementService_GrantUserRights_Handler,
		},
		{
			MethodName: "RevokeUserRights",
			Handler:    _UserManagementService_RevokeUserRights_Handler,
		},
		{
			MethodName: "ListUserRights",
			Handler:    _UserManagementService_ListUserRights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/daml/ledger/api/v1/admin/user_management_service.proto",
}
