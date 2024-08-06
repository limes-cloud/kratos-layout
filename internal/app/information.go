package app

import (
	"context"
	"poverty/internal/infra/rpc"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"poverty/api/poverty/errors"
	pb "poverty/api/poverty/information/v1"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/service"
	"poverty/internal/infra/dbs"
	"poverty/internal/types"
)

type InformationApp struct {
	pb.UnimplementedInformationServer
	srv *service.InformationService
}

func NewInformationApp(conf *conf.Config) *InformationApp {
	return &InformationApp{
		srv: service.NewInformationService(conf, dbs.NewInformationInfra(), rpc.NewFileInfra()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewInformationApp(c)
		pb.RegisterInformationHTTPServer(hs, srv)
		pb.RegisterInformationServer(gs, srv)
	})
}

// GetInformation 获取指定的资讯信息
func (s *InformationApp) GetInformation(c context.Context, req *pb.GetInformationRequest) (*pb.GetInformationReply, error) {
	var ctx = kratosx.MustContext(c)

	ent, err := s.srv.GetInformation(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	reply := pb.GetInformationReply{}
	if err := valx.Transform(ent, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListInformation 获取资讯信息列表
func (s *InformationApp) ListInformation(c context.Context, req *pb.ListInformationRequest) (*pb.ListInformationReply, error) {
	var ctx = kratosx.MustContext(c)
	result, total, err := s.srv.ListInformation(ctx, &types.ListInformationRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Title:    req.Title,
		IsTop:    req.IsTop,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListInformationReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateInformation 创建资讯信息
func (s *InformationApp) CreateInformation(c context.Context, req *pb.CreateInformationRequest) (*pb.CreateInformationReply, error) {
	var (
		ent = entity.Information{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateInformation(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateInformationReply{Id: id}, nil
}

// UpdateInformation 更新资讯信息
func (s *InformationApp) UpdateInformation(c context.Context, req *pb.UpdateInformationRequest) (*pb.UpdateInformationReply, error) {
	var (
		ent = entity.Information{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateInformation(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.UpdateInformationReply{}, nil
}

// DeleteInformation 删除资讯信息
func (s *InformationApp) DeleteInformation(c context.Context, req *pb.DeleteInformationRequest) (*pb.DeleteInformationReply, error) {
	return &pb.DeleteInformationReply{}, s.srv.DeleteInformation(kratosx.MustContext(c), req.Id)
}
