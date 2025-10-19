package service

import (
	"github.com/limes-cloud/kratosx/library/pool"
	"partyaffairs/internal/core"

	"partyaffairs/api/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type NoticeService struct {
	repo repository.NoticeRepository
	user repository.UserRepository
}

func NewNoticeService(repo repository.NoticeRepository, user repository.UserRepository) *NoticeService {
	return &NoticeService{repo: repo, user: user}
}

// GetNotice 获取指定的通知信息
func (srv *NoticeService) GetNotice(ctx core.Context, id uint32) (*entity.Notice, error) {
	notice, err := srv.repo.GetNotice(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	if err := srv.repo.ReadUserNotice(ctx, ctx.Auth().UserId, id); err != nil {
		return nil, errors.GetError(err.Error())
	}
	return notice, nil
}

// ListNotice 获取通知信息列表
func (srv *NoticeService) ListNotice(ctx core.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	list, total, err := srv.repo.ListNotice(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateNotice 创建通知信息
func (srv *NoticeService) CreateNotice(ctx core.Context, req *entity.Notice) (uint32, error) {
	id, err := srv.repo.CreateNotice(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

func (srv *NoticeService) PushNotice(ctx core.Context, id uint32) error {
	// todo 后续接入消息推送中心

	// 获取通知
	notice, err := srv.repo.GetNotice(ctx, id)
	if err != nil {
		return err
	}
	if notice.Status == nil || !*notice.Status {
		return errors.NoticeStatusError()
	}

	_ = ctx.Pool().Go(pool.AddRunner(ctx.Clone(), func() {
		cctx := ctx.Clone()
		// 获取用户列表
		var page uint32 = 1
		var pageSize uint32 = 50
		for {
			list, _, err := srv.user.ListUser(cctx, &types.ListUserRequest{
				Page:     pageSize,
				PageSize: pageSize,
			})
			if err != nil {
				panic(err)
			}

			// 发送邮件
			//for _, item := range list {
			//	//if item.Email != nil {
			//	//	template := ctx.Email().Template("notice")
			//	//	err := template.Send(*item.Email, email.WithTemplateVariable(map[string]any{
			//	//		"title": notice.Title,
			//	//		"desc":  notice.Description,
			//	//	}))
			//	//	if err != nil {
			//	//		panic(err)
			//	//	}
			//	//}
			//}
			if len(list) < int(pageSize) {
				break
			}
			page++
		}
	}))

	return nil
}

// UpdateNotice 更新通知信息
func (srv *NoticeService) UpdateNotice(ctx core.Context, ent *entity.Notice) error {
	if err := srv.repo.UpdateNotice(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotice 删除通知信息
func (srv *NoticeService) DeleteNotice(ctx core.Context, id uint32) error {
	if err := srv.repo.DeleteNotice(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
