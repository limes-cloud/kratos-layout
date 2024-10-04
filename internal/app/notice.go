package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "partyaffairs/api/partyaffairs/notice/v1"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Notice struct {
	pb.UnimplementedNoticeServer
	srv *service.NoticeService
}

func NewNotice(conf *conf.Config) *Notice {
	return &Notice{
		srv: service.NewNoticeService(conf, dbs.NewNotice(), rpc.NewUser()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server) {
		srv := NewNotice(c)
		pb.RegisterNoticeHTTPServer(hs, srv)
	})
}

// GetNotice 获取指定的通知信息
func (s *Notice) GetNotice(c context.Context, req *pb.GetNoticeRequest) (*pb.GetNoticeReply, error) {
	notice, err := s.srv.GetNotice(kratosx.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetNoticeReply{
		Id:          notice.Id,
		Title:       notice.Title,
		Description: notice.Description,
		Unit:        notice.Unit,
		Content:     notice.Content,
		IsTop:       notice.IsTop,
		Status:      notice.Status,
		CreatedAt:   uint32(notice.CreatedAt),
		UpdatedAt:   uint32(notice.UpdatedAt),
	}, nil
}

// ListNotice 获取通知信息列表
func (s *Notice) ListNotice(c context.Context, req *pb.ListNoticeRequest) (*pb.ListNoticeReply, error) {
	list, total, err := s.srv.ListNotice(kratosx.MustContext(c), &types.ListNoticeRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Title:    req.Title,
		IsTop:    req.IsTop,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListNoticeReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListNoticeReply_Notice{
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
func (s *Notice) CreateNotice(c context.Context, req *pb.CreateNoticeRequest) (*pb.CreateNoticeReply, error) {
	id, err := s.srv.CreateNotice(kratosx.MustContext(c), &entity.Notice{
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
	return &pb.CreateNoticeReply{Id: id}, nil
}

// PushNotice 推送通知信息
func (s *Notice) PushNotice(c context.Context, req *pb.PushNoticeRequest) (*pb.PushNoticeReply, error) {
	return &pb.PushNoticeReply{}, s.srv.PushNotice(kratosx.MustContext(c), req.Id)
}

// UpdateNotice 更新通知信息
func (s *Notice) UpdateNotice(c context.Context, req *pb.UpdateNoticeRequest) (*pb.UpdateNoticeReply, error) {
	if err := s.srv.UpdateNotice(kratosx.MustContext(c), &entity.Notice{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Title:       req.Title,
		Description: req.Description,
		Unit:        req.Unit,
		Content:     req.Content,
		IsTop:       req.IsTop,
		Status:      req.Status,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateNoticeReply{}, nil
}

// DeleteNotice 删除通知信息
func (s *Notice) DeleteNotice(c context.Context, req *pb.DeleteNoticeRequest) (*pb.DeleteNoticeReply, error) {
	return &pb.DeleteNoticeReply{}, s.srv.DeleteNotice(kratosx.MustContext(c), req.Id)
}
