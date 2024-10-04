// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/partyaffairs/information/partyaffairs_information_service.proto

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
	Information_ListInformationClassify_FullMethodName   = "/partyaffairs.api.partyaffairs.information.v1.Information/ListInformationClassify"
	Information_CreateInformationClassify_FullMethodName = "/partyaffairs.api.partyaffairs.information.v1.Information/CreateInformationClassify"
	Information_UpdateInformationClassify_FullMethodName = "/partyaffairs.api.partyaffairs.information.v1.Information/UpdateInformationClassify"
	Information_DeleteInformationClassify_FullMethodName = "/partyaffairs.api.partyaffairs.information.v1.Information/DeleteInformationClassify"
	Information_GetInformation_FullMethodName            = "/partyaffairs.api.partyaffairs.information.v1.Information/GetInformation"
	Information_ListInformation_FullMethodName           = "/partyaffairs.api.partyaffairs.information.v1.Information/ListInformation"
	Information_CreateInformation_FullMethodName         = "/partyaffairs.api.partyaffairs.information.v1.Information/CreateInformation"
	Information_UpdateInformation_FullMethodName         = "/partyaffairs.api.partyaffairs.information.v1.Information/UpdateInformation"
	Information_DeleteInformation_FullMethodName         = "/partyaffairs.api.partyaffairs.information.v1.Information/DeleteInformation"
)

// InformationClient is the client API for Information service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformationClient interface {
	// ListInformationClassify 获取资讯分组列表
	ListInformationClassify(ctx context.Context, in *ListInformationClassifyRequest, opts ...grpc.CallOption) (*ListInformationClassifyReply, error)
	// CreateInformationClassify 创建资讯分组
	CreateInformationClassify(ctx context.Context, in *CreateInformationClassifyRequest, opts ...grpc.CallOption) (*CreateInformationClassifyReply, error)
	// UpdateInformationClassify 更新资讯分组
	UpdateInformationClassify(ctx context.Context, in *UpdateInformationClassifyRequest, opts ...grpc.CallOption) (*UpdateInformationClassifyReply, error)
	// DeleteInformationClassify 删除资讯分组
	DeleteInformationClassify(ctx context.Context, in *DeleteInformationClassifyRequest, opts ...grpc.CallOption) (*DeleteInformationClassifyReply, error)
	// GetInformation 获取指定的咨询信息
	GetInformation(ctx context.Context, in *GetInformationRequest, opts ...grpc.CallOption) (*GetInformationReply, error)
	// ListInformation 获取咨询信息列表
	ListInformation(ctx context.Context, in *ListInformationRequest, opts ...grpc.CallOption) (*ListInformationReply, error)
	// CreateInformation 创建咨询信息
	CreateInformation(ctx context.Context, in *CreateInformationRequest, opts ...grpc.CallOption) (*CreateInformationReply, error)
	// UpdateInformation 更新咨询信息
	UpdateInformation(ctx context.Context, in *UpdateInformationRequest, opts ...grpc.CallOption) (*UpdateInformationReply, error)
	// DeleteInformation 删除咨询信息
	DeleteInformation(ctx context.Context, in *DeleteInformationRequest, opts ...grpc.CallOption) (*DeleteInformationReply, error)
}

type informationClient struct {
	cc grpc.ClientConnInterface
}

func NewInformationClient(cc grpc.ClientConnInterface) InformationClient {
	return &informationClient{cc}
}

func (c *informationClient) ListInformationClassify(ctx context.Context, in *ListInformationClassifyRequest, opts ...grpc.CallOption) (*ListInformationClassifyReply, error) {
	out := new(ListInformationClassifyReply)
	err := c.cc.Invoke(ctx, Information_ListInformationClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) CreateInformationClassify(ctx context.Context, in *CreateInformationClassifyRequest, opts ...grpc.CallOption) (*CreateInformationClassifyReply, error) {
	out := new(CreateInformationClassifyReply)
	err := c.cc.Invoke(ctx, Information_CreateInformationClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) UpdateInformationClassify(ctx context.Context, in *UpdateInformationClassifyRequest, opts ...grpc.CallOption) (*UpdateInformationClassifyReply, error) {
	out := new(UpdateInformationClassifyReply)
	err := c.cc.Invoke(ctx, Information_UpdateInformationClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) DeleteInformationClassify(ctx context.Context, in *DeleteInformationClassifyRequest, opts ...grpc.CallOption) (*DeleteInformationClassifyReply, error) {
	out := new(DeleteInformationClassifyReply)
	err := c.cc.Invoke(ctx, Information_DeleteInformationClassify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) GetInformation(ctx context.Context, in *GetInformationRequest, opts ...grpc.CallOption) (*GetInformationReply, error) {
	out := new(GetInformationReply)
	err := c.cc.Invoke(ctx, Information_GetInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) ListInformation(ctx context.Context, in *ListInformationRequest, opts ...grpc.CallOption) (*ListInformationReply, error) {
	out := new(ListInformationReply)
	err := c.cc.Invoke(ctx, Information_ListInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) CreateInformation(ctx context.Context, in *CreateInformationRequest, opts ...grpc.CallOption) (*CreateInformationReply, error) {
	out := new(CreateInformationReply)
	err := c.cc.Invoke(ctx, Information_CreateInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) UpdateInformation(ctx context.Context, in *UpdateInformationRequest, opts ...grpc.CallOption) (*UpdateInformationReply, error) {
	out := new(UpdateInformationReply)
	err := c.cc.Invoke(ctx, Information_UpdateInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informationClient) DeleteInformation(ctx context.Context, in *DeleteInformationRequest, opts ...grpc.CallOption) (*DeleteInformationReply, error) {
	out := new(DeleteInformationReply)
	err := c.cc.Invoke(ctx, Information_DeleteInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformationServer is the server API for Information service.
// All implementations must embed UnimplementedInformationServer
// for forward compatibility
type InformationServer interface {
	// ListInformationClassify 获取资讯分组列表
	ListInformationClassify(context.Context, *ListInformationClassifyRequest) (*ListInformationClassifyReply, error)
	// CreateInformationClassify 创建资讯分组
	CreateInformationClassify(context.Context, *CreateInformationClassifyRequest) (*CreateInformationClassifyReply, error)
	// UpdateInformationClassify 更新资讯分组
	UpdateInformationClassify(context.Context, *UpdateInformationClassifyRequest) (*UpdateInformationClassifyReply, error)
	// DeleteInformationClassify 删除资讯分组
	DeleteInformationClassify(context.Context, *DeleteInformationClassifyRequest) (*DeleteInformationClassifyReply, error)
	// GetInformation 获取指定的咨询信息
	GetInformation(context.Context, *GetInformationRequest) (*GetInformationReply, error)
	// ListInformation 获取咨询信息列表
	ListInformation(context.Context, *ListInformationRequest) (*ListInformationReply, error)
	// CreateInformation 创建咨询信息
	CreateInformation(context.Context, *CreateInformationRequest) (*CreateInformationReply, error)
	// UpdateInformation 更新咨询信息
	UpdateInformation(context.Context, *UpdateInformationRequest) (*UpdateInformationReply, error)
	// DeleteInformation 删除咨询信息
	DeleteInformation(context.Context, *DeleteInformationRequest) (*DeleteInformationReply, error)
	mustEmbedUnimplementedInformationServer()
}

// UnimplementedInformationServer must be embedded to have forward compatible implementations.
type UnimplementedInformationServer struct {
}

func (UnimplementedInformationServer) ListInformationClassify(context.Context, *ListInformationClassifyRequest) (*ListInformationClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInformationClassify not implemented")
}
func (UnimplementedInformationServer) CreateInformationClassify(context.Context, *CreateInformationClassifyRequest) (*CreateInformationClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInformationClassify not implemented")
}
func (UnimplementedInformationServer) UpdateInformationClassify(context.Context, *UpdateInformationClassifyRequest) (*UpdateInformationClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInformationClassify not implemented")
}
func (UnimplementedInformationServer) DeleteInformationClassify(context.Context, *DeleteInformationClassifyRequest) (*DeleteInformationClassifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInformationClassify not implemented")
}
func (UnimplementedInformationServer) GetInformation(context.Context, *GetInformationRequest) (*GetInformationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInformation not implemented")
}
func (UnimplementedInformationServer) ListInformation(context.Context, *ListInformationRequest) (*ListInformationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInformation not implemented")
}
func (UnimplementedInformationServer) CreateInformation(context.Context, *CreateInformationRequest) (*CreateInformationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInformation not implemented")
}
func (UnimplementedInformationServer) UpdateInformation(context.Context, *UpdateInformationRequest) (*UpdateInformationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInformation not implemented")
}
func (UnimplementedInformationServer) DeleteInformation(context.Context, *DeleteInformationRequest) (*DeleteInformationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInformation not implemented")
}
func (UnimplementedInformationServer) mustEmbedUnimplementedInformationServer() {}

// UnsafeInformationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformationServer will
// result in compilation errors.
type UnsafeInformationServer interface {
	mustEmbedUnimplementedInformationServer()
}

func RegisterInformationServer(s grpc.ServiceRegistrar, srv InformationServer) {
	s.RegisterService(&Information_ServiceDesc, srv)
}

func _Information_ListInformationClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInformationClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).ListInformationClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_ListInformationClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).ListInformationClassify(ctx, req.(*ListInformationClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_CreateInformationClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInformationClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).CreateInformationClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_CreateInformationClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).CreateInformationClassify(ctx, req.(*CreateInformationClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_UpdateInformationClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInformationClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).UpdateInformationClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_UpdateInformationClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).UpdateInformationClassify(ctx, req.(*UpdateInformationClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_DeleteInformationClassify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInformationClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).DeleteInformationClassify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_DeleteInformationClassify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).DeleteInformationClassify(ctx, req.(*DeleteInformationClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_GetInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).GetInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_GetInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).GetInformation(ctx, req.(*GetInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_ListInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).ListInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_ListInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).ListInformation(ctx, req.(*ListInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_CreateInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).CreateInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_CreateInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).CreateInformation(ctx, req.(*CreateInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_UpdateInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).UpdateInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_UpdateInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).UpdateInformation(ctx, req.(*UpdateInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Information_DeleteInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).DeleteInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Information_DeleteInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).DeleteInformation(ctx, req.(*DeleteInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Information_ServiceDesc is the grpc.ServiceDesc for Information service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Information_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partyaffairs.api.partyaffairs.information.v1.Information",
	HandlerType: (*InformationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInformationClassify",
			Handler:    _Information_ListInformationClassify_Handler,
		},
		{
			MethodName: "CreateInformationClassify",
			Handler:    _Information_CreateInformationClassify_Handler,
		},
		{
			MethodName: "UpdateInformationClassify",
			Handler:    _Information_UpdateInformationClassify_Handler,
		},
		{
			MethodName: "DeleteInformationClassify",
			Handler:    _Information_DeleteInformationClassify_Handler,
		},
		{
			MethodName: "GetInformation",
			Handler:    _Information_GetInformation_Handler,
		},
		{
			MethodName: "ListInformation",
			Handler:    _Information_ListInformation_Handler,
		},
		{
			MethodName: "CreateInformation",
			Handler:    _Information_CreateInformation_Handler,
		},
		{
			MethodName: "UpdateInformation",
			Handler:    _Information_UpdateInformation_Handler,
		},
		{
			MethodName: "DeleteInformation",
			Handler:    _Information_DeleteInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/partyaffairs/information/partyaffairs_information_service.proto",
}
