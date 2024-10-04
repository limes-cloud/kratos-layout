// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/partyaffairs/resource/partyaffairs_resource_service.proto

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

const (
	Resource_ListResourceClassify_FullMethodName   = "/partyaffairs.api.partyaffairs.resource.v1.Resource/ListResourceClassify"
	Resource_CreateResourceClassify_FullMethodName = "/partyaffairs.api.partyaffairs.resource.v1.Resource/CreateResourceClassify"
	Resource_UpdateResourceClassify_FullMethodName = "/partyaffairs.api.partyaffairs.resource.v1.Resource/UpdateResourceClassify"
	Resource_DeleteResourceClassify_FullMethodName = "/partyaffairs.api.partyaffairs.resource.v1.Resource/DeleteResourceClassify"
	Resource_GetResource_FullMethodName            = "/partyaffairs.api.partyaffairs.resource.v1.Resource/GetResource"
	Resource_ListResource_FullMethodName           = "/partyaffairs.api.partyaffairs.resource.v1.Resource/ListResource"
	Resource_CreateResource_FullMethodName         = "/partyaffairs.api.partyaffairs.resource.v1.Resource/CreateResource"
	Resource_UpdateResource_FullMethodName         = "/partyaffairs.api.partyaffairs.resource.v1.Resource/UpdateResource"
	Resource_DeleteResource_FullMethodName         = "/partyaffairs.api.partyaffairs.resource.v1.Resource/DeleteResource"
)

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceClient interface {
	// ListResourceClassify 获取资料分组列表
	ListResourceClassify(ctx context.Context, in *ListResourceClassifyRequest, opts ...grpc.CallOption) (*ListResourceClassifyReply, error)
	// CreateResourceClassify 创建资料分组
	CreateResourceClassify(ctx context.Context, in *CreateResourceClassifyRequest, opts ...grpc.CallOption) (*CreateResourceClassifyReply, error)
	// UpdateResourceClassify 更新资料分组
	UpdateResourceClassify(ctx context.Context, in *UpdateResourceClassifyRequest, opts ...grpc.CallOption) (*UpdateResourceClassifyReply, error)
	// DeleteResourceClassify 删除资料分组
	DeleteResourceClassify(ctx context.Context, in *DeleteResourceClassifyRequest, opts ...grpc.CallOption) (*DeleteResourceClassifyReply, error)
	// GetResource 获取指定的咨询信息
	GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceReply, error)
	// ListResource 获取咨询信息列表
	ListResource(ctx context.Context, in *ListResourceRequest, opts ...grpc.CallOption) (*ListResourceReply, error)
	// CreateResource 创建咨询信息
	CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceReply, error)
	// UpdateResource 更新咨询信息
	UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceReply, error)
	// DeleteResource 删除咨询信息
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceReply, error)
}

type resourceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceClient(cc grpc.ClientConnInterface) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) ListResourceClassify(ctx context.Context, in *ListResourceClassifyRequest, opts ...grpc.CallOption) (*ListResourceClassifyReply, error) {
	out := new(ListResourceClassifyReply)
	err := c.cc.Invoke(ctx, Resource_ListResourceClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) CreateResourceClassify(ctx context.Context, in *CreateResourceClassifyRequest, opts ...grpc.CallOption) (*CreateResourceClassifyReply, error) {
	out := new(CreateResourceClassifyReply)
	err := c.cc.Invoke(ctx, Resource_CreateResourceClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) UpdateResourceClassify(ctx context.Context, in *UpdateResourceClassifyRequest, opts ...grpc.CallOption) (*UpdateResourceClassifyReply, error) {
	out := new(UpdateResourceClassifyReply)
	err := c.cc.Invoke(ctx, Resource_UpdateResourceClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteResourceClassify(ctx context.Context, in *DeleteResourceClassifyRequest, opts ...grpc.CallOption) (*DeleteResourceClassifyReply, error) {
	out := new(DeleteResourceClassifyReply)
	err := c.cc.Invoke(ctx, Resource_DeleteResourceClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) GetResource(ctx context.Context, in *GetResourceRequest, opts ...grpc.CallOption) (*GetResourceReply, error) {
	out := new(GetResourceReply)
	err := c.cc.Invoke(ctx, Resource_GetResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) ListResource(ctx context.Context, in *ListResourceRequest, opts ...grpc.CallOption) (*ListResourceReply, error) {
	out := new(ListResourceReply)
	err := c.cc.Invoke(ctx, Resource_ListResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceReply, error) {
	out := new(CreateResourceReply)
	err := c.cc.Invoke(ctx, Resource_CreateResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceReply, error) {
	out := new(UpdateResourceReply)
	err := c.cc.Invoke(ctx, Resource_UpdateResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceReply, error) {
	out := new(DeleteResourceReply)
	err := c.cc.Invoke(ctx, Resource_DeleteResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServer is the server API for Resource service.
// All implementations must embed UnimplementedResourceServer
// for forward compatibility
type ResourceServer interface {
	// ListResourceClassify 获取资料分组列表
	ListResourceClassify(context.Context, *ListResourceClassifyRequest) (*ListResourceClassifyReply, error)
	// CreateResourceClassify 创建资料分组
	CreateResourceClassify(context.Context, *CreateResourceClassifyRequest) (*CreateResourceClassifyReply, error)
	// UpdateResourceClassify 更新资料分组
	UpdateResourceClassify(context.Context, *UpdateResourceClassifyRequest) (*UpdateResourceClassifyReply, error)
	// DeleteResourceClassify 删除资料分组
	DeleteResourceClassify(context.Context, *DeleteResourceClassifyRequest) (*DeleteResourceClassifyReply, error)
	// GetResource 获取指定的咨询信息
	GetResource(context.Context, *GetResourceRequest) (*GetResourceReply, error)
	// ListResource 获取咨询信息列表
	ListResource(context.Context, *ListResourceRequest) (*ListResourceReply, error)
	// CreateResource 创建咨询信息
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceReply, error)
	// UpdateResource 更新咨询信息
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceReply, error)
	// DeleteResource 删除咨询信息
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceReply, error)
	mustEmbedUnimplementedResourceServer()
}

// UnimplementedResourceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (UnimplementedResourceServer) ListResourceClassify(context.Context, *ListResourceClassifyRequest) (*ListResourceClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResourceClassify not implemented")
}
func (UnimplementedResourceServer) CreateResourceClassify(context.Context, *CreateResourceClassifyRequest) (*CreateResourceClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResourceClassify not implemented")
}
func (UnimplementedResourceServer) UpdateResourceClassify(context.Context, *UpdateResourceClassifyRequest) (*UpdateResourceClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResourceClassify not implemented")
}
func (UnimplementedResourceServer) DeleteResourceClassify(context.Context, *DeleteResourceClassifyRequest) (*DeleteResourceClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResourceClassify not implemented")
}
func (UnimplementedResourceServer) GetResource(context.Context, *GetResourceRequest) (*GetResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResource not implemented")
}
func (UnimplementedResourceServer) ListResource(context.Context, *ListResourceRequest) (*ListResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResource not implemented")
}
func (UnimplementedResourceServer) CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedResourceServer) UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedResourceServer) DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedResourceServer) mustEmbedUnimplementedResourceServer() {}

// UnsafeResourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServer will
// result in compilation errors.
type UnsafeResourceServer interface {
	mustEmbedUnimplementedResourceServer()
}

func RegisterResourceServer(s grpc.ServiceRegistrar, srv ResourceServer) {
	s.RegisterService(&Resource_ServiceDesc, srv)
}

func _Resource_ListResourceClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourceClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).ListResourceClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_ListResourceClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).ListResourceClassify(ctx, req.(*ListResourceClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_CreateResourceClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).CreateResourceClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_CreateResourceClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).CreateResourceClassify(ctx, req.(*CreateResourceClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_UpdateResourceClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).UpdateResourceClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_UpdateResourceClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).UpdateResourceClassify(ctx, req.(*UpdateResourceClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteResourceClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteResourceClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_DeleteResourceClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteResourceClassify(ctx, req.(*DeleteResourceClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_GetResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).GetResource(ctx, req.(*GetResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_ListResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).ListResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_ListResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).ListResource(ctx, req.(*ListResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_CreateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).CreateResource(ctx, req.(*CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_UpdateResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).UpdateResource(ctx, req.(*UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Resource_DeleteResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resource_ServiceDesc is the grpc.ServiceDesc for Resource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partyaffairs.api.partyaffairs.resource.v1.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListResourceClassify",
			Handler:    _Resource_ListResourceClassify_Handler,
		},
		{
			MethodName: "CreateResourceClassify",
			Handler:    _Resource_CreateResourceClassify_Handler,
		},
		{
			MethodName: "UpdateResourceClassify",
			Handler:    _Resource_UpdateResourceClassify_Handler,
		},
		{
			MethodName: "DeleteResourceClassify",
			Handler:    _Resource_DeleteResourceClassify_Handler,
		},
		{
			MethodName: "GetResource",
			Handler:    _Resource_GetResource_Handler,
		},
		{
			MethodName: "ListResource",
			Handler:    _Resource_ListResource_Handler,
		},
		{
			MethodName: "CreateResource",
			Handler:    _Resource_CreateResource_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _Resource_UpdateResource_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _Resource_DeleteResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/partyaffairs/resource/partyaffairs_resource_service.proto",
}
