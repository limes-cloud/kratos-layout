package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	userV1 "github.com/limes-cloud/usercenter/api/usercenter/user/v1"
	"google.golang.org/protobuf/proto"

	"partyaffairs/api/errors"
	v1 "partyaffairs/api/partyaffairs/v1"
	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
	"partyaffairs/internal/consts"
	"partyaffairs/internal/pkg/service"
)

func (s *Service) PageNotice(ctx context.Context, in *v1.PageNoticeRequest) (*v1.PageNoticeReply, error) {
	list, total, err := s.notice.Page(kratosx.MustContext(ctx), &types.PageNoticeRequest{
		Page:     in.Page,
		PageSize: in.PageSize,
		Title:    in.Title,
		NotRead:  in.NotRead,
	})
	if err != nil {
		return nil, err
	}

	reply := v1.PageNoticeReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformError()
	}
	for ind, item := range list {
		reply.List[ind].IsRead = item.NoticeUser != nil
	}

	return &reply, nil
}

func (s *Service) PageNoticeUser(ctx context.Context, in *v1.PageNoticeUserRequest) (*v1.PageNoticeUserReply, error) {
	ids, err := s.notice.ReadUserIds(kratosx.MustContext(ctx), in.NoticeId)
	if err != nil {
		return nil, err
	}

	client, err := service.NewUser(ctx)
	if err != nil {
		return nil, errors.UserCenterError()
	}
	req := &userV1.ListUserRequest{
		App:      proto.String(consts.ClientApp),
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	if in.IsRead {
		req.InIds = ids
	} else {
		req.NotInIds = ids
	}
	userReply, err := client.ListUser(ctx, req)
	if err != nil {
		return nil, err
	}

	reply := &v1.PageNoticeUserReply{Total: userReply.Total}
	for _, item := range userReply.List {
		reply.List = append(reply.List, &v1.PageNoticeUserReply_User{
			Id:        item.Id,
			NickName:  item.NickName,
			RealName:  item.RealName,
			AvatarUrl: item.AvatarUrl,
			Gender:    item.Gender,
		})
	}
	return reply, nil
}

func (s *Service) GetNotice(ctx context.Context, in *v1.GetNoticeRequest) (*v1.Notice, error) {
	nc, err := s.notice.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.Notice{}
	if err := copier.Copy(&reply, nc); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Service) ReadNotice(ctx context.Context, in *v1.ReadNoticeRequest) (*empty.Empty, error) {
	return nil, s.notice.Read(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) AddNotice(ctx context.Context, in *v1.AddNoticeRequest) (*empty.Empty, error) {
	var notice biz.Notice
	kCtx := kratosx.MustContext(ctx)
	if err := copier.Copy(&notice, in); err != nil {
		return nil, errors.TransformError()
	}

	if _, err := s.notice.Add(kCtx, &notice); err != nil {
		return nil, err
	}
	go func() {
		if err := s.notice.SendNoticeEmail(kratosx.MustContext(context.Background()), notice.Id); err != nil {
			kCtx.Logger().Errorf("发送邮件失败：%s", err.Error())
		}
	}()
	return nil, nil
}

func (s *Service) UpdateNotice(ctx context.Context, in *v1.UpdateNoticeRequest) (*empty.Empty, error) {
	var notice biz.Notice
	if err := copier.Copy(&notice, in); err != nil {
		return nil, errors.TransformError()
	}

	kCtx := kratosx.MustContext(ctx)
	if err := s.notice.Update(kCtx, &notice); err != nil {
		return nil, nil
	}
	go func() {
		if err := s.notice.SendNoticeEmail(kratosx.MustContext(context.Background()), notice.Id); err != nil {
			kCtx.Logger().Errorf("发送邮件失败：%s", err.Error())
		}
	}()
	return nil, s.notice.Update(kratosx.MustContext(ctx), &notice)
}

func (s *Service) DeleteNotice(ctx context.Context, in *v1.DeleteNoticeRequest) (*empty.Empty, error) {
	return nil, s.notice.Delete(kratosx.MustContext(ctx), in.Id)
}
