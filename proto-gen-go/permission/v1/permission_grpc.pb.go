// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package permissionv1

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	// Find all Permissions
	FindAllPermissions(ctx context.Context, in *FindAllPermissionsRequest, opts ...grpc.CallOption) (*FindAllPermissionsResponse, error)
	// Find Permission by ID
	FindPermissionByID(ctx context.Context, in *CommonUUIDRequest, opts ...grpc.CallOption) (*Permission, error)
	// Create new Permission
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// Update Permission by ID
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// Delete Permission
	DeletePermission(ctx context.Context, in *CommonUUIDRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) FindAllPermissions(ctx context.Context, in *FindAllPermissionsRequest, opts ...grpc.CallOption) (*FindAllPermissionsResponse, error) {
	out := new(FindAllPermissionsResponse)
	err := c.cc.Invoke(ctx, "/permission.v1.PermissionService/FindAllPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) FindPermissionByID(ctx context.Context, in *CommonUUIDRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := c.cc.Invoke(ctx, "/permission.v1.PermissionService/FindPermissionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/permission.v1.PermissionService/CreatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/permission.v1.PermissionService/UpdatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) DeletePermission(ctx context.Context, in *CommonUUIDRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/permission.v1.PermissionService/DeletePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations should embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	// Find all Permissions
	FindAllPermissions(context.Context, *FindAllPermissionsRequest) (*FindAllPermissionsResponse, error)
	// Find Permission by ID
	FindPermissionByID(context.Context, *CommonUUIDRequest) (*Permission, error)
	// Create new Permission
	CreatePermission(context.Context, *CreatePermissionRequest) (*CommonResponse, error)
	// Update Permission by ID
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*CommonResponse, error)
	// Delete Permission
	DeletePermission(context.Context, *CommonUUIDRequest) (*CommonResponse, error)
}

// UnimplementedPermissionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) FindAllPermissions(context.Context, *FindAllPermissionsRequest) (*FindAllPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) FindPermissionByID(context.Context, *CommonUUIDRequest) (*Permission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPermissionByID not implemented")
}
func (UnimplementedPermissionServiceServer) CreatePermission(context.Context, *CreatePermissionRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedPermissionServiceServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedPermissionServiceServer) DeletePermission(context.Context, *CommonUUIDRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_FindAllPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).FindAllPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.v1.PermissionService/FindAllPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).FindAllPermissions(ctx, req.(*FindAllPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_FindPermissionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).FindPermissionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.v1.PermissionService/FindPermissionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).FindPermissionByID(ctx, req.(*CommonUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.v1.PermissionService/CreatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.v1.PermissionService/UpdatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permission.v1.PermissionService/DeletePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).DeletePermission(ctx, req.(*CommonUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "permission.v1.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllPermissions",
			Handler:    _PermissionService_FindAllPermissions_Handler,
		},
		{
			MethodName: "FindPermissionByID",
			Handler:    _PermissionService_FindPermissionByID_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _PermissionService_CreatePermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _PermissionService_UpdatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _PermissionService_DeletePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission/v1/permission.proto",
}