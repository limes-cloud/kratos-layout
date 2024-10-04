package service

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/pool"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

type NoticeService struct {
	conf *conf.Config
	repo repository.NoticeRepository
	user repository.UserRepository
	file repository.FileRepository
}

func NewNoticeService(conf *conf.Config, repo repository.NoticeRepository, user repository.UserRepository) *NoticeService {
	return &NoticeService{conf: conf, repo: repo, user: user}
}

// GetNotice 获取指定的通知信息
func (srv *NoticeService) GetNotice(ctx kratosx.Context, id uint32) (*entity.Notice, error) {
	notice, err := srv.repo.GetNotice(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	// 获取通知内容，判断是否具有图片
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(notice.Content))
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		origin, _ := selection.Attr("data-origin")
		md5, _ := selection.Attr("data-md5")
		if origin == "resource" && md5 != "" {
			// 获取图片连接地址
			src := srv.file.GetFileURL(ctx, md5)
			if src == "" {
				selection.Remove()
			} else {
				selection.SetAttr("src", src)
			}
		}
	})

	out, err := doc.Find("body").Html()
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	notice.Content = out
	return notice, nil
}

// ListNotice 获取通知信息列表
func (srv *NoticeService) ListNotice(ctx kratosx.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
	list, total, err := srv.repo.ListNotice(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// ListCurrentNotReadNotice 获取通知信息列表
// func (srv *NoticeService) ListCurNotReadNotice(ctx kratosx.Context, req *types.ListNoticeRequest) ([]*entity.Notice, uint32, error) {
//	list, total, err := srv.repo.ListNotice(ctx, req)
//	if err != nil {
//		return nil, 0, errors.ListError(err.Error())
//	}
//	return list, total, nil
// }

// CreateNotice 创建通知信息
func (srv *NoticeService) CreateNotice(ctx kratosx.Context, req *entity.Notice) (uint32, error) {
	id, err := srv.repo.CreateNotice(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

func (srv *NoticeService) PushNotice(ctx kratosx.Context, id uint32) error {
	// todo 后续接入消息推送中心

	// 获取通知
	notice, err := srv.repo.GetNotice(ctx, id)
	if err != nil {
		return err
	}
	if notice.Status == nil || !*notice.Status {
		return errors.NoticeStatusError()
	}

	_ = ctx.Go(pool.AddRunner(func() error {
		cctx := ctx.Clone()
		// 获取用户列表
		var page uint32 = 1
		var pageSize uint32 = 50
		for {
			defer func() {
				if r := recover(); r != nil {
					cctx.Logger().Errorf("推送通知异常: %v", r)
				}
			}()
			list, _, err := srv.user.ListUser(cctx, &types.ListUserRequest{
				Page:     pageSize,
				PageSize: pageSize,
			})
			if err != nil {
				return errors.ListUserError()
			}

			// 发送邮件
			for _, item := range list {
				if item.Email != nil {
					template := ctx.Email().Template("notice")
					return template.Send(*item.Email, *item.RealName, map[string]any{
						"title": notice.Title,
						"desc":  notice.Description,
					})
				}
			}
			if len(list) < int(pageSize) {
				break
			}
			page++
		}
		return nil
	}))

	return nil
}

// UpdateNotice 更新通知信息
func (srv *NoticeService) UpdateNotice(ctx kratosx.Context, ent *entity.Notice) error {
	if err := srv.repo.UpdateNotice(ctx, ent); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotice 删除通知信息
func (srv *NoticeService) DeleteNotice(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteNotice(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
