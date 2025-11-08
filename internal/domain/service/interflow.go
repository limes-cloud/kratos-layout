package service

import (
	"encoding/json"
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
func (srv *InterflowService) ListInterflowHistory(ctx core.Context) ([]*entity.User, error) {

	ids, err := srv.repo.ListInterflowFormUserIds(ctx, ctx.Auth().UserId)
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	if len(ids) == 0 {
		return nil, nil
	}

	bucket, _, err := srv.user.ListUser(ctx, &types.ListUserRequest{
		Page:     1,
		PageSize: uint32(len(ids)),
		In:       ids,
	})
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return bucket, nil
}

// ListInterflowClassify 获取资讯分组列表
func (srv *InterflowService) ListInterflowClassify(ctx core.Context) ([]*entity.InterflowClassify, error) {
	list, err := srv.repo.ListInterflowClassify(ctx)

	for i, v := range list {
		var ids []uint32
		_ = json.Unmarshal([]byte(v.Person), &ids)
		if len(ids) == 0 {
			continue
		}
		bucket, _, err := srv.user.ListUser(ctx, &types.ListUserRequest{
			Page:     1,
			PageSize: uint32(len(ids)),
			In:       ids,
		})
		if err != nil {
			return nil, errors.ListError(err.Error())
		}
		list[i].Users = bucket

		//for _, bucketValue := range bucket {
		//	v.Users = append(v.Users, bu)
		//}

	}

	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// CreateInterflowClassify 创建资讯分组
func (srv *InterflowService) CreateInterflowClassify(ctx core.Context, tg *entity.InterflowClassify) (uint32, error) {
	id, err := srv.repo.CreateInterflowClassify(ctx, tg)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateInterflowClassify 更新资讯分组
func (srv *InterflowService) UpdateInterflowClassify(ctx core.Context, tg *entity.InterflowClassify) error {
	if err := srv.repo.UpdateInterflowClassify(ctx, tg); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteInterflowClassify 删除资讯分组
func (srv *InterflowService) DeleteInterflowClassify(ctx core.Context, id uint32) error {
	err := srv.repo.DeleteInterflowClassify(ctx, id)
	if err != nil {
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
