package service

import (
	"github.com/limes-cloud/kratosx"

	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/repository"
	"poverty/internal/types"
)

type CommentService struct {
	conf *conf.Config
	repo repository.CommentRepository
	user repository.UserRepository
}

func NewCommentService(conf *conf.Config, repo repository.CommentRepository, user repository.UserRepository) *CommentService {
	return &CommentService{conf: conf, repo: repo, user: user}
}

// GetComment 获取指定的留言信息
func (srv *CommentService) GetComment(ctx kratosx.Context, id uint32) (*entity.Comment, error) {
	ent, err := srv.repo.GetComment(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return ent, nil
}

// ListComment 获取留言信息列表
func (srv *CommentService) ListComment(ctx kratosx.Context, req *types.ListCommentRequest) ([]*entity.Comment, uint32, error) {
	if req.UserName != nil {
		users, err := srv.user.ListUserByName(ctx, *req.UserName)
		if err != nil {
			return nil, 0, errors.ListError(err.Error())
		}
		req.UserIds = make([]uint32, 0)
		for _, item := range users {
			req.UserIds = append(req.UserIds, item.Id)
		}
	}
	list, total, err := srv.repo.ListComment(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}

	for ind, item := range list {
		user, err := srv.user.GetUser(ctx, item.UserId)
		if err != nil {
			continue
		}
		if user.RealName != nil {
			list[ind].UserName = *user.RealName
		}
		if user.AvatarUrl != nil {
			list[ind].UserAvatar = *user.AvatarUrl
		}
	}

	return list, total, nil
}

// CreateComment 创建留言信息
func (srv *CommentService) CreateComment(ctx kratosx.Context, req *entity.Comment) (uint32, error) {
	id, err := srv.repo.CreateComment(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateComment 更新留言信息
func (srv *CommentService) UpdateComment(ctx kratosx.Context, ent *entity.Comment) error {
	if err := srv.repo.UpdateComment(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteComment 删除留言信息
func (srv *CommentService) DeleteComment(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteComment(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
