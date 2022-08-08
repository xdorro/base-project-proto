// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: permission/v1/permission.proto

package permissionv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/xdorro/proto-base-project/proto-gen-go/permission/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// PermissionServiceName is the fully-qualified name of the PermissionService service.
	PermissionServiceName = "permission.v1.PermissionService"
)

// PermissionServiceClient is a client for the permission.v1.PermissionService service.
type PermissionServiceClient interface {
	// Find all Permissions
	FindAllPermissions(context.Context, *connect_go.Request[v1.FindAllPermissionsRequest]) (*connect_go.Response[v1.FindAllPermissionsResponse], error)
	// Find Permission by ID
	FindPermissionByID(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.Permission], error)
	// Create new Permission
	CreatePermission(context.Context, *connect_go.Request[v1.CreatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Update Permission by ID
	UpdatePermission(context.Context, *connect_go.Request[v1.UpdatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Delete Permission
	DeletePermission(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error)
}

// NewPermissionServiceClient constructs a client for the permission.v1.PermissionService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPermissionServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PermissionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &permissionServiceClient{
		findAllPermissions: connect_go.NewClient[v1.FindAllPermissionsRequest, v1.FindAllPermissionsResponse](
			httpClient,
			baseURL+"/permission.v1.PermissionService/FindAllPermissions",
			opts...,
		),
		findPermissionByID: connect_go.NewClient[v1.CommonUUIDRequest, v1.Permission](
			httpClient,
			baseURL+"/permission.v1.PermissionService/FindPermissionByID",
			opts...,
		),
		createPermission: connect_go.NewClient[v1.CreatePermissionRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/permission.v1.PermissionService/CreatePermission",
			opts...,
		),
		updatePermission: connect_go.NewClient[v1.UpdatePermissionRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/permission.v1.PermissionService/UpdatePermission",
			opts...,
		),
		deletePermission: connect_go.NewClient[v1.CommonUUIDRequest, v1.CommonResponse](
			httpClient,
			baseURL+"/permission.v1.PermissionService/DeletePermission",
			opts...,
		),
	}
}

// permissionServiceClient implements PermissionServiceClient.
type permissionServiceClient struct {
	findAllPermissions *connect_go.Client[v1.FindAllPermissionsRequest, v1.FindAllPermissionsResponse]
	findPermissionByID *connect_go.Client[v1.CommonUUIDRequest, v1.Permission]
	createPermission   *connect_go.Client[v1.CreatePermissionRequest, v1.CommonResponse]
	updatePermission   *connect_go.Client[v1.UpdatePermissionRequest, v1.CommonResponse]
	deletePermission   *connect_go.Client[v1.CommonUUIDRequest, v1.CommonResponse]
}

// FindAllPermissions calls permission.v1.PermissionService.FindAllPermissions.
func (c *permissionServiceClient) FindAllPermissions(ctx context.Context, req *connect_go.Request[v1.FindAllPermissionsRequest]) (*connect_go.Response[v1.FindAllPermissionsResponse], error) {
	return c.findAllPermissions.CallUnary(ctx, req)
}

// FindPermissionByID calls permission.v1.PermissionService.FindPermissionByID.
func (c *permissionServiceClient) FindPermissionByID(ctx context.Context, req *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.Permission], error) {
	return c.findPermissionByID.CallUnary(ctx, req)
}

// CreatePermission calls permission.v1.PermissionService.CreatePermission.
func (c *permissionServiceClient) CreatePermission(ctx context.Context, req *connect_go.Request[v1.CreatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.createPermission.CallUnary(ctx, req)
}

// UpdatePermission calls permission.v1.PermissionService.UpdatePermission.
func (c *permissionServiceClient) UpdatePermission(ctx context.Context, req *connect_go.Request[v1.UpdatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.updatePermission.CallUnary(ctx, req)
}

// DeletePermission calls permission.v1.PermissionService.DeletePermission.
func (c *permissionServiceClient) DeletePermission(ctx context.Context, req *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return c.deletePermission.CallUnary(ctx, req)
}

// PermissionServiceHandler is an implementation of the permission.v1.PermissionService service.
type PermissionServiceHandler interface {
	// Find all Permissions
	FindAllPermissions(context.Context, *connect_go.Request[v1.FindAllPermissionsRequest]) (*connect_go.Response[v1.FindAllPermissionsResponse], error)
	// Find Permission by ID
	FindPermissionByID(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.Permission], error)
	// Create new Permission
	CreatePermission(context.Context, *connect_go.Request[v1.CreatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Update Permission by ID
	UpdatePermission(context.Context, *connect_go.Request[v1.UpdatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error)
	// Delete Permission
	DeletePermission(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error)
}

// NewPermissionServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPermissionServiceHandler(svc PermissionServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/permission.v1.PermissionService/FindAllPermissions", connect_go.NewUnaryHandler(
		"/permission.v1.PermissionService/FindAllPermissions",
		svc.FindAllPermissions,
		opts...,
	))
	mux.Handle("/permission.v1.PermissionService/FindPermissionByID", connect_go.NewUnaryHandler(
		"/permission.v1.PermissionService/FindPermissionByID",
		svc.FindPermissionByID,
		opts...,
	))
	mux.Handle("/permission.v1.PermissionService/CreatePermission", connect_go.NewUnaryHandler(
		"/permission.v1.PermissionService/CreatePermission",
		svc.CreatePermission,
		opts...,
	))
	mux.Handle("/permission.v1.PermissionService/UpdatePermission", connect_go.NewUnaryHandler(
		"/permission.v1.PermissionService/UpdatePermission",
		svc.UpdatePermission,
		opts...,
	))
	mux.Handle("/permission.v1.PermissionService/DeletePermission", connect_go.NewUnaryHandler(
		"/permission.v1.PermissionService/DeletePermission",
		svc.DeletePermission,
		opts...,
	))
	return "/permission.v1.PermissionService/", mux
}

// UnimplementedPermissionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPermissionServiceHandler struct{}

func (UnimplementedPermissionServiceHandler) FindAllPermissions(context.Context, *connect_go.Request[v1.FindAllPermissionsRequest]) (*connect_go.Response[v1.FindAllPermissionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("permission.v1.PermissionService.FindAllPermissions is not implemented"))
}

func (UnimplementedPermissionServiceHandler) FindPermissionByID(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.Permission], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("permission.v1.PermissionService.FindPermissionByID is not implemented"))
}

func (UnimplementedPermissionServiceHandler) CreatePermission(context.Context, *connect_go.Request[v1.CreatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("permission.v1.PermissionService.CreatePermission is not implemented"))
}

func (UnimplementedPermissionServiceHandler) UpdatePermission(context.Context, *connect_go.Request[v1.UpdatePermissionRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("permission.v1.PermissionService.UpdatePermission is not implemented"))
}

func (UnimplementedPermissionServiceHandler) DeletePermission(context.Context, *connect_go.Request[v1.CommonUUIDRequest]) (*connect_go.Response[v1.CommonResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("permission.v1.PermissionService.DeletePermission is not implemented"))
}