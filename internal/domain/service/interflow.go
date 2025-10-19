package service

import (
	"partyaffairs/internal/core"

	"partyaffairs/api/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type InterflowService struct {
	repo repository.InterflowRepository
	user repository.UserRepository
}

func NewInterflowService(repo repository.InterflowRepository, user repository.UserRepository) *InterflowService {
	return &InterflowService{repo: repo, user: user}
}

// ListInterflow 获取通知信息列表
func (srv *InterflowService) ListInterflowPerson(ctx core.Context) ([]*entity.InterflowPerson, error) {
	list, err := srv.repo.ListInterflowPerson(ctx)
	if err != nil {
		return nil, errors.ListError(err.Error())
	}

	var ids []uint32
	for _, v := range list {
		ids = append(ids, v.UserID)
	}

	bucket, err := srv.user.ListUserMap(ctx, &types.ListUserRequest{
		Page:     1,
		PageSize: uint32(len(ids)),
		In:       ids,
	})
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	for i, v := range list {
		user := bucket[v.UserID]
		list[i].Username = user.Username
		list[i].Nickname = user.Nickname
		list[i].Avatar = user.Avatar
	}

	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreateInterflow 创建通知信息
func (srv *InterflowService) CreateInterflowPerson(ctx core.Context, req *entity.InterflowPerson) (uint32, error) {
	id, err := srv.repo.CreateInterflowPerson(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// DeleteInterflow 删除通知信息
func (srv *InterflowService) DeleteInterflowPerson(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteInterflowPerson(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListInterflow 获取通知信息列表
func (srv *InterflowService) ListInterflow(ctx core.Context, req *types.ListInterflowRequest) ([]*entity.Interflow, uint32, error) {
	if req.FromCur {
		req.FromUserId = ctx.Auth().UserId
	}
	list, total, err := srv.repo.ListInterflow(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateInterflow 创建通知信息
func (srv *InterflowService) CreateInterflow(ctx core.Context, req *entity.Interflow) (uint32, error) {
	req.FromUserID = ctx.Auth().UserId
	id, err := srv.repo.CreateInterflow(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateInterflow 更新通知信息
func (srv *InterflowService) UpdateInterflow(ctx core.Context, ent *entity.Interflow) error {
	if err := srv.repo.UpdateInterflow(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteInterflow 删除通知信息
func (srv *InterflowService) DeleteInterflow(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteInterflow(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
