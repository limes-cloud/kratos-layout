// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/poverty/resource/poverty_resource_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationResourceCreateResource = "/poverty.api.poverty.resource.v1.Resource/CreateResource"
const OperationResourceDeleteResource = "/poverty.api.poverty.resource.v1.Resource/DeleteResource"
const OperationResourceGetResource = "/poverty.api.poverty.resource.v1.Resource/GetResource"
const OperationResourceListResource = "/poverty.api.poverty.resource.v1.Resource/ListResource"
const OperationResourceUpdateResource = "/poverty.api.poverty.resource.v1.Resource/UpdateResource"

type ResourceHTTPServer interface {
	// CreateResource CreateResource 创建文件信息
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceReply, error)
	// DeleteResource DeleteResource 删除文件信息
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceReply, error)
	// GetResource GetResource 获取指定的文件信息
	GetResource(context.Context, *GetResourceRequest) (*GetResourceReply, error)
	// ListResource ListResource 获取文件信息列表
	ListResource(context.Context, *ListResourceRequest) (*ListResourceReply, error)
	// UpdateResource UpdateResource 更新文件信息
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceReply, error)
}

func RegisterResourceHTTPServer(s *http.Server, srv ResourceHTTPServer) {
	r := s.Route("/")
	r.GET("/poverty/api/v1/resource", _Resource_GetResource0_HTTP_Handler(srv))
	r.GET("/poverty/client/v1/resources", _Resource_ListResource0_HTTP_Handler(srv))
	r.GET("/poverty/api/v1/resources", _Resource_ListResource1_HTTP_Handler(srv))
	r.POST("/poverty/api/v1/resource", _Resource_CreateResource0_HTTP_Handler(srv))
	r.PUT("/poverty/api/v1/resource", _Resource_UpdateResource0_HTTP_Handler(srv))
	r.DELETE("/poverty/api/v1/resource", _Resource_DeleteResource0_HTTP_Handler(srv))
}

func _Resource_GetResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceGetResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetResource(ctx, req.(*GetResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_ListResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceListResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListResource(ctx, req.(*ListResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_ListResource1_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceListResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListResource(ctx, req.(*ListResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_CreateResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceCreateResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateResource(ctx, req.(*CreateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_UpdateResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceUpdateResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateResource(ctx, req.(*UpdateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_DeleteResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceDeleteResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteResource(ctx, req.(*DeleteResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResourceReply)
		return ctx.Result(200, reply)
	}
}

type ResourceHTTPClient interface {
	CreateResource(ctx context.Context, req *CreateResourceRequest, opts ...http.CallOption) (rsp *CreateResourceReply, err error)
	DeleteResource(ctx context.Context, req *DeleteResourceRequest, opts ...http.CallOption) (rsp *DeleteResourceReply, err error)
	GetResource(ctx context.Context, req *GetResourceRequest, opts ...http.CallOption) (rsp *GetResourceReply, err error)
	ListResource(ctx context.Context, req *ListResourceRequest, opts ...http.CallOption) (rsp *ListResourceReply, err error)
	UpdateResource(ctx context.Context, req *UpdateResourceRequest, opts ...http.CallOption) (rsp *UpdateResourceReply, err error)
}

type ResourceHTTPClientImpl struct {
	cc *http.Client
}

func NewResourceHTTPClient(client *http.Client) ResourceHTTPClient {
	return &ResourceHTTPClientImpl{client}
}

func (c *ResourceHTTPClientImpl) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...http.CallOption) (*CreateResourceReply, error) {
	var out CreateResourceReply
	pattern := "/poverty/api/v1/resource"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationResourceCreateResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...http.CallOption) (*DeleteResourceReply, error) {
	var out DeleteResourceReply
	pattern := "/poverty/api/v1/resource"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceDeleteResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) GetResource(ctx context.Context, in *GetResourceRequest, opts ...http.CallOption) (*GetResourceReply, error) {
	var out GetResourceReply
	pattern := "/poverty/api/v1/resource"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceGetResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) ListResource(ctx context.Context, in *ListResourceRequest, opts ...http.CallOption) (*ListResourceReply, error) {
	var out ListResourceReply
	pattern := "/poverty/api/v1/resources"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceListResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...http.CallOption) (*UpdateResourceReply, error) {
	var out UpdateResourceReply
	pattern := "/poverty/api/v1/resource"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationResourceUpdateResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}