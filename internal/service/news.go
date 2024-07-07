package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/resource/file/v1"
	userV1 "github.com/limes-cloud/usercenter/api/usercenter/user/v1"

	"partyaffairs/api/errors"
	v1 "partyaffairs/api/partyaffairs/v1"
	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
	"partyaffairs/internal/pkg/service"
)

func (s *Service) AllNewsClassify(ctx context.Context, _ *empty.Empty) (*v1.AllNewsClassifyReply, error) {
	list, err := s.news.AllClassify(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	var reply v1.AllNewsClassifyReply
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddNewsClassify(ctx context.Context, in *v1.AddNewsClassifyRequest) (*empty.Empty, error) {
	var nc biz.NewsClassify
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.TransformError()
	}
	_, err := s.news.AddClassify(kratosx.MustContext(ctx), &nc)
	return nil, err
}

func (s *Service) UpdateNewsClassify(ctx context.Context, in *v1.UpdateNewsClassifyRequest) (*empty.Empty, error) {
	var nc biz.NewsClassify
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.TransformError()
	}
	return nil, s.news.UpdateClassify(kratosx.MustContext(ctx), &nc)
}

func (s *Service) DeleteNewsClassify(ctx context.Context, in *v1.DeleteNewsClassifyRequest) (*empty.Empty, error) {
	return nil, s.news.DeleteClassify(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) PageNewsContent(ctx context.Context, in *v1.PageNewsContentRequest) (*v1.PageNewsContentReply, error) {
	list, total, err := s.news.PageContent(kratosx.MustContext(ctx), &types.PageNewsContentRequest{
		Page:       in.Page,
		PageSize:   in.PageSize,
		ClassifyID: in.ClassifyId,
		Title:      in.Title,
	})
	if err != nil {
		return nil, err
	}

	reply := v1.PageNewsContentReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformError()
	}

	resource, err := service.NewResourceFile(ctx)
	if err == nil {
		for ind, item := range reply.List {
			file, err := resource.GetFile(ctx, &resourceV1.GetFileRequest{Sha: &item.Cover})
			if err != nil {
				continue
			}
			reply.List[ind].Resource = &v1.NewsContent_File{
				Name: file.Name,
				Sha:  file.Sha,
				Url:  file.Url,
			}
		}
	}

	return &reply, nil
}

func (s *Service) GetNewsContent(ctx context.Context, in *v1.GetNewsContentRequest) (*v1.NewsContent, error) {
	nc, err := s.news.GetContent(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.NewsContent{}
	if err := copier.Copy(&reply, nc); err != nil {
		return nil, errors.TransformError()
	}

	// 转换资源
	resource, err := service.NewResourceFile(ctx)
	if err == nil {
		if file, err := resource.GetFile(ctx, &resourceV1.GetFileRequest{Sha: &reply.Cover}); err == nil {
			reply.Resource = &v1.NewsContent_File{
				Name: file.Name,
				Sha:  file.Sha,
				Url:  file.Url,
			}
		}
	}

	// 获取用户
	user, err := service.NewUser(ctx)
	if err == nil {
		for ind, item := range nc.Comments {
			if item.FromUid != 0 {
				userReply, err := user.GetUser(ctx, &userV1.GetUserRequest{Id: &item.FromUid})
				if err != nil {
					continue
				}
				reply.Comments[ind].From = &v1.NewsComment_User{
					Id:        userReply.Id,
					NickName:  userReply.NickName,
					RealName:  userReply.RealName,
					AvatarUrl: userReply.AvatarUrl,
					Gender:    userReply.Gender,
				}
			}
			if item.ReplyUid != 0 {
				userReply, err := user.GetUser(ctx, &userV1.GetUserRequest{Id: &item.ReplyUid})
				if err != nil {
					continue
				}
				reply.Comments[ind].Reply = &v1.NewsComment_User{
					Id:        userReply.Id,
					NickName:  userReply.NickName,
					RealName:  userReply.RealName,
					AvatarUrl: userReply.AvatarUrl,
					Gender:    userReply.Gender,
				}
			}
		}
	}

	return &reply, nil
}

func (s *Service) AddNewsContent(ctx context.Context, in *v1.AddNewsContentRequest) (*empty.Empty, error) {
	var nc biz.NewsContent
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.TransformError()
	}
	_, err := s.news.AddContent(kratosx.MustContext(ctx), &nc)
	return nil, err
}

func (s *Service) UpdateNewsContent(ctx context.Context, in *v1.UpdateNewsContentRequest) (*empty.Empty, error) {
	var nc biz.NewsContent
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.TransformError()
	}
	return nil, s.news.UpdateContent(kratosx.MustContext(ctx), &nc)
}

func (s *Service) DeleteNewsContent(ctx context.Context, in *v1.DeleteNewsContentRequest) (*empty.Empty, error) {
	return nil, s.news.DeleteContent(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) AddNewsComment(ctx context.Context, in *v1.AddNewsCommentRequest) (*v1.AddNewsCommentReply, error) {
	var nc biz.NewsComment
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.TransformError()
	}
	id, err := s.news.AddComment(kratosx.MustContext(ctx), &nc)
	return &v1.AddNewsCommentReply{Id: id}, err
}

func (s *Service) DeleteNewsComment(ctx context.Context, in *v1.DeleteNewsCommentRequest) (*empty.Empty, error) {
	return nil, s.news.DeleteComment(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) DeleteCurNewsComment(ctx context.Context, in *v1.DeleteNewsCommentRequest) (*empty.Empty, error) {
	return nil, s.news.DeleteCurComment(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) PageNewsComment(ctx context.Context, in *v1.PageNewsCommentRequest) (*v1.PageNewsCommentReply, error) {
	req := types.PageNewsCommentRequest{}
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.TransformError()
	}

	list, total, err := s.news.PageComment(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageNewsCommentReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformError()
	}

	// 获取用户
	user, err := service.NewUser(ctx)
	if err == nil {
		for ind, item := range list {
			if item.FromUid != 0 {
				userReply, err := user.GetUser(ctx, &userV1.GetUserRequest{Id: &item.FromUid})
				if err != nil {
					continue
				}
				reply.List[ind].From = &v1.NewsComment_User{
					Id:        userReply.Id,
					NickName:  userReply.NickName,
					RealName:  userReply.RealName,
					AvatarUrl: userReply.AvatarUrl,
					Gender:    userReply.Gender,
				}
			}
			if item.ReplyUid != 0 {
				userReply, err := user.GetUser(ctx, &userV1.GetUserRequest{Id: &item.ReplyUid})
				if err != nil {
					continue
				}
				reply.List[ind].Reply = &v1.NewsComment_User{
					Id:        userReply.Id,
					NickName:  userReply.NickName,
					RealName:  userReply.RealName,
					AvatarUrl: userReply.AvatarUrl,
					Gender:    userReply.Gender,
				}
			}
		}
	}

	return &reply, nil
}
