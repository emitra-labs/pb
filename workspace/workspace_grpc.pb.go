// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.4
// source: workspace/workspace.proto

package workspace

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Workspace_Authorize_FullMethodName = "/workspace.Workspace/Authorize"
)

// WorkspaceClient is the client API for Workspace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

type workspaceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceClient(cc grpc.ClientConnInterface) WorkspaceClient {
	return &workspaceClient{cc}
}

func (c *workspaceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, Workspace_Authorize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServer is the server API for Workspace service.
// All implementations must embed UnimplementedWorkspaceServer
// for forward compatibility.
type WorkspaceServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	mustEmbedUnimplementedWorkspaceServer()
}

// UnimplementedWorkspaceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkspaceServer struct{}

func (UnimplementedWorkspaceServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedWorkspaceServer) mustEmbedUnimplementedWorkspaceServer() {}
func (UnimplementedWorkspaceServer) testEmbeddedByValue()                   {}

// UnsafeWorkspaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServer will
// result in compilation errors.
type UnsafeWorkspaceServer interface {
	mustEmbedUnimplementedWorkspaceServer()
}

func RegisterWorkspaceServer(s grpc.ServiceRegistrar, srv WorkspaceServer) {
	// If the following call pancis, it indicates UnimplementedWorkspaceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Workspace_ServiceDesc, srv)
}

func _Workspace_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Workspace_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Workspace_ServiceDesc is the grpc.ServiceDesc for Workspace service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workspace_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workspace.Workspace",
	HandlerType: (*WorkspaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _Workspace_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace/workspace.proto",
}
