// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.4
// source: kratos_layout_notice_service.proto

package noticepb

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationServiceAddNotice = "/noticepb.Service/AddNotice"
const OperationServiceDeleteNotice = "/noticepb.Service/DeleteNotice"
const OperationServiceGetNotice = "/noticepb.Service/GetNotice"
const OperationServicePageNotice = "/noticepb.Service/PageNotice"
const OperationServiceUpdateNotice = "/noticepb.Service/UpdateNotice"

type ServiceHTTPServer interface {
	AddNotice(context.Context, *AddNoticeRequest) (*AddNoticeReply, error)
	DeleteNotice(context.Context, *DeleteNoticeRequest) (*emptypb.Empty, error)
	// GetNotice GetNotice 获取用户信息
	GetNotice(context.Context, *GetNoticeRequest) (*Notice, error)
	PageNotice(context.Context, *PageNoticeRequest) (*PageNoticeReply, error)
	UpdateNotice(context.Context, *UpdateNoticeRequest) (*emptypb.Empty, error)
}

func RegisterServiceHTTPServer(s *http.Server, srv ServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/kratos-layout/notice/api/v1/notice", _Service_GetNotice0_HTTP_Handler(srv))
	r.GET("/kratos-layout/notice/api/v1/notices", _Service_PageNotice0_HTTP_Handler(srv))
	r.POST("/kratos-layout/notice/api/v1/notice", _Service_AddNotice0_HTTP_Handler(srv))
	r.PUT("/kratos-layout/notice/api/v1/notice", _Service_UpdateNotice0_HTTP_Handler(srv))
	r.PUT("/kratos-layout/notice/api/v1/notice", _Service_DeleteNotice0_HTTP_Handler(srv))
}

func _Service_GetNotice0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetNoticeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetNotice)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetNotice(ctx, req.(*GetNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Notice)
		return ctx.Result(200, reply)
	}
}

func _Service_PageNotice0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageNoticeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServicePageNotice)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.PageNotice(ctx, req.(*PageNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageNoticeReply)
		return ctx.Result(200, reply)
	}
}

func _Service_AddNotice0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddNoticeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddNotice)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.AddNotice(ctx, req.(*AddNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddNoticeReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateNotice0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateNoticeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateNotice)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateNotice(ctx, req.(*UpdateNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteNotice0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteNoticeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteNotice)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteNotice(ctx, req.(*DeleteNoticeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type ServiceHTTPClient interface {
	AddNotice(ctx context.Context, req *AddNoticeRequest, opts ...http.CallOption) (rsp *AddNoticeReply, err error)
	DeleteNotice(ctx context.Context, req *DeleteNoticeRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetNotice(ctx context.Context, req *GetNoticeRequest, opts ...http.CallOption) (rsp *Notice, err error)
	PageNotice(ctx context.Context, req *PageNoticeRequest, opts ...http.CallOption) (rsp *PageNoticeReply, err error)
	UpdateNotice(ctx context.Context, req *UpdateNoticeRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type ServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceHTTPClient(client *http.Client) ServiceHTTPClient {
	return &ServiceHTTPClientImpl{client}
}

func (c *ServiceHTTPClientImpl) AddNotice(ctx context.Context, in *AddNoticeRequest, opts ...http.CallOption) (*AddNoticeReply, error) {
	var out AddNoticeReply
	pattern := "/kratos-layout/notice/api/v1/notice"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteNotice(ctx context.Context, in *DeleteNoticeRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/kratos-layout/notice/api/v1/notice"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceDeleteNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetNotice(ctx context.Context, in *GetNoticeRequest, opts ...http.CallOption) (*Notice, error) {
	var out Notice
	pattern := "/kratos-layout/notice/api/v1/notice"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) PageNotice(ctx context.Context, in *PageNoticeRequest, opts ...http.CallOption) (*PageNoticeReply, error) {
	var out PageNoticeReply
	pattern := "/kratos-layout/notice/api/v1/notices"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServicePageNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateNotice(ctx context.Context, in *UpdateNoticeRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/kratos-layout/notice/api/v1/notice"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateNotice))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
