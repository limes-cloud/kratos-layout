package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "layout/api/layout/dictionary/v1"
	"layout/api/layout/errors"
	"layout/internal/conf"
	"layout/internal/domain/entity"
	"layout/internal/domain/service"
	"layout/internal/infra/dbs"
	"layout/internal/types"
)

type Dictionary struct {
	pb.UnimplementedDictionaryServer
	srv *service.Dictionary
}

func NewDictionary(conf *conf.Config) *Dictionary {
	return &Dictionary{
		srv: service.NewDictionary(conf, dbs.NewDictionary()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		app := NewDictionary(c)
		pb.RegisterDictionaryHTTPServer(hs, app)
		pb.RegisterDictionaryServer(gs, app)
	})
}

// GetDictionary 获取指定的字典目录
func (s *Dictionary) GetDictionary(c context.Context, req *pb.GetDictionaryRequest) (*pb.GetDictionaryReply, error) {
	ctx := kratosx.MustContext(c)
	result, err := s.srv.GetDictionary(ctx, &types.GetDictionaryRequest{
		Id:      req.Id,
		Keyword: req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.GetDictionaryReply{}
	if err = valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// ListDictionary 获取字典目录列表
func (s *Dictionary) ListDictionary(c context.Context, req *pb.ListDictionaryRequest) (*pb.ListDictionaryReply, error) {
	ctx := kratosx.MustContext(c)

	result, total, err := s.srv.ListDictionary(ctx, &types.ListDictionaryRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListDictionaryReply{Total: total}
	if err = valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDictionary 创建字典目录
func (s *Dictionary) CreateDictionary(c context.Context, req *pb.CreateDictionaryRequest) (*pb.CreateDictionaryReply, error) {
	var (
		ctx = kratosx.MustContext(c)
		ent = entity.Dictionary{}
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.srv.CreateDictionary(ctx, &ent)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDictionaryReply{Id: id}, nil
}

// UpdateDictionary 更新字典目录
func (s *Dictionary) UpdateDictionary(c context.Context, req *pb.UpdateDictionaryRequest) (*pb.UpdateDictionaryReply, error) {
	var (
		ctx = kratosx.MustContext(c)
		ent = entity.Dictionary{}
	)

	if err := valx.Transform(req, &ent); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.srv.UpdateDictionary(ctx, &ent); err != nil {
		return nil, err
	}

	return &pb.UpdateDictionaryReply{}, nil
}

// DeleteDictionary 删除字典目录
func (s *Dictionary) DeleteDictionary(c context.Context, req *pb.DeleteDictionaryRequest) (*pb.DeleteDictionaryReply, error) {
	err := s.srv.DeleteDictionary(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDictionaryReply{}, nil
}
