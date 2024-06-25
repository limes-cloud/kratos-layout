package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	pb "layout/api/layout/dictionary/v1"
	"layout/api/layout/errors"
	"layout/internal/biz/dictionary"
	"layout/internal/conf"
	"layout/internal/data"
)

type DictionaryService struct {
	pb.UnimplementedDictionaryServer
	uc *dictionary.UseCase
}

func NewDictionaryService(conf *conf.Config) *DictionaryService {
	return &DictionaryService{
		uc: dictionary.NewUseCase(conf, data.NewDictionaryRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewDictionaryService(c)
		pb.RegisterDictionaryHTTPServer(hs, srv)
		pb.RegisterDictionaryServer(gs, srv)
	})
}

// GetDictionary 获取指定的字典目录
func (s *DictionaryService) GetDictionary(c context.Context, req *pb.GetDictionaryRequest) (*pb.GetDictionaryReply, error) {
	var (
		in  = dictionary.GetDictionaryRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetDictionary(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetDictionaryReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListDictionary 获取字典目录列表
func (s *DictionaryService) ListDictionary(c context.Context, req *pb.ListDictionaryRequest) (*pb.ListDictionaryReply, error) {
	var (
		in  = dictionary.ListDictionaryRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListDictionary(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListDictionaryReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateDictionary 创建字典目录
func (s *DictionaryService) CreateDictionary(c context.Context, req *pb.CreateDictionaryRequest) (*pb.CreateDictionaryReply, error) {
	var (
		in  = dictionary.Dictionary{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateDictionary(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDictionaryReply{Id: id}, nil
}

// UpdateDictionary 更新字典目录
func (s *DictionaryService) UpdateDictionary(c context.Context, req *pb.UpdateDictionaryRequest) (*pb.UpdateDictionaryReply, error) {
	var (
		in  = dictionary.Dictionary{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateDictionary(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateDictionaryReply{}, nil
}

// DeleteDictionary 删除字典目录
func (s *DictionaryService) DeleteDictionary(c context.Context, req *pb.DeleteDictionaryRequest) (*pb.DeleteDictionaryReply, error) {
	total, err := s.uc.DeleteDictionary(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDictionaryReply{Total: total}, nil
}
