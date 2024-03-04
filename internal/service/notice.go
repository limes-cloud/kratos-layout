package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"

	"github.com/go-kratos/kratos-layout/api/errors"
	"github.com/go-kratos/kratos-layout/api/noticepb"
	noticebiz "github.com/go-kratos/kratos-layout/internal/biz/notice"
	"github.com/go-kratos/kratos-layout/internal/conf"
	noticedata "github.com/go-kratos/kratos-layout/internal/data/notice"
)

type NoticeService struct {
	noticepb.UnimplementedServiceServer
	notice *noticebiz.UseCase
}

func NewNotice(conf *conf.Config) *NoticeService {
	return &NoticeService{
		notice: noticebiz.NewUseCase(conf, noticedata.NewRepo()),
	}
}

func (s *NoticeService) GetNotice(ctx context.Context, req *noticepb.GetNoticeRequest) (*noticepb.Notice, error) {
	user, err := s.notice.Get(kratosx.MustContext(ctx), req.Id)
	if err != nil {
		return nil, err
	}
	reply := noticepb.Notice{}
	if err := copier.Copy(&reply, user); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}
	return &reply, nil
}

func (s *NoticeService) PageNotice(ctx context.Context, req *noticepb.PageNoticeRequest) (*noticepb.PageNoticeReply, error) {
	bizReq := noticebiz.PageRequest{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	list, total, err := s.notice.Page(kratosx.MustContext(ctx), &bizReq)
	if err != nil {
		return nil, err
	}

	reply := noticepb.PageNoticeReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	return &reply, nil
}

func (s *NoticeService) AddNotice(ctx context.Context, req *noticepb.AddNoticeRequest) (*noticepb.AddNoticeReply, error) {
	bizReq := noticebiz.Notice{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	id, err := s.notice.Add(kratosx.MustContext(ctx), &bizReq)
	if err != nil {
		return nil, err
	}

	return &noticepb.AddNoticeReply{Id: id}, nil
}

func (s *NoticeService) UpdateNotice(ctx context.Context, req *noticepb.UpdateNoticeRequest) (*empty.Empty, error) {
	bizReq := noticebiz.Notice{}
	if err := copier.Copy(&bizReq, req); err != nil {
		return nil, errors.TransformFormat(err.Error())
	}

	if err := s.notice.Update(kratosx.MustContext(ctx), &bizReq); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *NoticeService) DeleteNotice(ctx context.Context, req *noticepb.DeleteNoticeRequest) (*empty.Empty, error) {
	return nil, s.notice.Delete(kratosx.MustContext(ctx), req.Id)
}
