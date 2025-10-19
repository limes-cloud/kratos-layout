package app

import (
	"context"
	"github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/model/page"
	"github.com/limes-cloud/kratosx/pkg/value"
	"partyaffairs/api/notice"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Notice struct {
	notice.UnimplementedNoticeServer
	srv *service.NoticeService
}

func NewNotice() *Notice {
	return &Notice{
		srv: service.NewNoticeService(dbs.NewNotice(), rpc.NewUser()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewNotice()
		notice.RegisterNoticeHTTPServer(hs, srv)
	})
}

// GetNotice 获取指定的通知信息
func (s *Notice) GetNotice(c context.Context, req *notice.GetNoticeRequest) (*notice.GetNoticeReply, error) {
	res, err := s.srv.GetNotice(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &notice.GetNoticeReply{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
		Unit:        res.Unit,
		Content:     res.Content,
		IsTop:       res.IsTop,
		Status:      res.Status,
		CreatedAt:   uint32(res.CreatedAt),
		UpdatedAt:   uint32(res.UpdatedAt),
	}, nil
}

// ListNotice 获取通知信息列表
func (s *Notice) ListNotice(c context.Context, req *notice.ListNoticeRequest) (*notice.ListNoticeReply, error) {
	list, total, err := s.srv.ListNotice(core.MustContext(c), &types.ListNoticeRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Title:  req.Title,
		IsTop:  req.IsTop,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := notice.ListNoticeReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &notice.ListNoticeReply_Notice{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			Unit:        item.Unit,
			IsTop:       item.IsTop,
			Status:      item.Status,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// ListVisibleNotice 获取可见通知信息列表
func (s *Notice) ListVisibleNotice(c context.Context, req *notice.ListNoticeRequest) (*notice.ListNoticeReply, error) {
	list, total, err := s.srv.ListNotice(core.MustContext(c), &types.ListNoticeRequest{
		Search: &page.Search{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Status:  value.Pointer(true),
		NotRead: req.NotRead,
	})
	if err != nil {
		return nil, err
	}

	reply := notice.ListNoticeReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &notice.ListNoticeReply_Notice{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			Unit:        item.Unit,
			IsTop:       item.IsTop,
			Status:      item.Status,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateNotice 创建通知信息
func (s *Notice) CreateNotice(c context.Context, req *notice.CreateNoticeRequest) (*notice.CreateNoticeReply, error) {
	id, err := s.srv.CreateNotice(core.MustContext(c), &entity.Notice{
		Title:       req.Title,
		Description: req.Description,
		Unit:        req.Unit,
		Content:     req.Content,
		IsTop:       req.IsTop,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &notice.CreateNoticeReply{Id: id}, nil
}

// PushNotice 推送通知信息
func (s *Notice) PushNotice(c context.Context, req *notice.PushNoticeRequest) (*notice.PushNoticeReply, error) {
	return &notice.PushNoticeReply{}, s.srv.PushNotice(core.MustContext(c), req.Id)
}

// UpdateNotice 更新通知信息
func (s *Notice) UpdateNotice(c context.Context, req *notice.UpdateNoticeRequest) (*notice.UpdateNoticeReply, error) {
	if err := s.srv.UpdateNotice(core.MustContext(c), &entity.Notice{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Title:           req.Title,
		Description:     req.Description,
		Unit:            req.Unit,
		Content:         req.Content,
		IsTop:           req.IsTop,
		Status:          req.Status,
	}); err != nil {
		return nil, err
	}
	return &notice.UpdateNoticeReply{}, nil
}

// DeleteNotice 删除通知信息
func (s *Notice) DeleteNotice(c context.Context, req *notice.DeleteNoticeRequest) (*notice.DeleteNoticeReply, error) {
	return &notice.DeleteNoticeReply{}, s.srv.DeleteNotice(core.MustContext(c), req.Id)
}
