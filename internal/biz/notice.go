package biz

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/pool"
	ktypes "github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/usercenter/api/usercenter/auth"
	userV1 "github.com/limes-cloud/usercenter/api/usercenter/user/v1"
	"google.golang.org/protobuf/proto"

	"partyaffairs/api/errors"
	"partyaffairs/internal/biz/types"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/consts"
	"partyaffairs/internal/pkg/service"
)

type Notice struct {
	Title      string      `json:"title" gorm:"not null;size:128;comment:公告标题"`
	Desc       string      `json:"desc" gorm:"not null;size:512;comment:公告描述"`
	Unit       string      `json:"unit" gorm:"not null;size:128;comment:发布单位"`
	Content    string      `json:"content" gorm:"not null;type:text;comment:轮播权重"`
	IsTop      *bool       `json:"is_top" gorm:"default false;comment:是否置顶"`
	NoticeUser *NoticeUser `json:"notice_user" gorm:"constraint:onDelete:cascade"`
	ktypes.BaseModel
}

type NoticeUser struct {
	UserID   uint32 `json:"user_id" gorm:"uniqueIndex:un;not null;comment:用户id"`
	NoticeID uint32 `json:"notice_id" gorm:"uniqueIndex:un;not null;comment:通知id"`
	ktypes.CreateModel
}

type NoticeRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Notice, error)
	Page(ctx kratosx.Context, req *types.PageNoticeRequest) ([]*Notice, uint32, error)
	Create(ctx kratosx.Context, c *Notice) (uint32, error)
	ReadUserIds(ctx kratosx.Context, nid uint32) ([]uint32, error)
	Update(ctx kratosx.Context, c *Notice) error
	Delete(ctx kratosx.Context, id uint32) error
	AddNoticeUser(ctx kratosx.Context, nu *NoticeUser) error
}

type NoticeUseCase struct {
	config *conf.Config
	repo   NoticeRepo
}

func NewNoticeUseCase(config *conf.Config, repo NoticeRepo) *NoticeUseCase {
	return &NoticeUseCase{config: config, repo: repo}
}

// Get 获取指定的公告
func (u *NoticeUseCase) Get(ctx kratosx.Context, id uint32) (*Notice, error) {
	notice, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.DatabaseError()
	}
	return notice, nil
}

// ReadUserIds 获取全部已读用户id
func (u *NoticeUseCase) ReadUserIds(ctx kratosx.Context, nid uint32) ([]uint32, error) {
	ids, err := u.repo.ReadUserIds(ctx, nid)
	if err != nil {
		return nil, errors.DatabaseError()
	}
	return ids, nil
}

// Page 获取全部公告
func (u *NoticeUseCase) Page(ctx kratosx.Context, req *types.PageNoticeRequest) ([]*Notice, uint32, error) {
	md, err := auth.Get(ctx)
	if err == nil && md.UserId != 0 {
		req.UserID = &md.UserId
	}
	notice, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, total, errors.DatabaseError()
	}
	return notice, total, nil
}

// Add 添加公告信息
func (u *NoticeUseCase) Add(ctx kratosx.Context, notice *Notice) (uint32, error) {
	id, err := u.repo.Create(ctx, notice)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

func (u *NoticeUseCase) SendNoticeEmail(ctx kratosx.Context, nid uint32) error {
	notice, err := u.repo.Get(ctx, nid)
	if err != nil {
		return err
	}

	// 获取全部应用用户
	client, err := service.NewUser(ctx)
	if err != nil {
		return errors.UserCenterError()
	}

	var page uint32 = 1
	var pageSize uint32 = 50
	for {
		req := &userV1.ListUserRequest{
			App:      proto.String(consts.ClientApp),
			Page:     page,
			PageSize: pageSize,
		}
		resp, err := client.ListUser(context.Background(), req)
		if err != nil {
			return errors.UserCenterError()
		}

		// 发送邮件
		for _, item := range resp.List {
			if item.Email != nil {
				user := item
				_ = ctx.Go(pool.AddRunner(func() error {
					template := ctx.Email().Template("notice")
					return template.Send(*user.Email, *user.RealName, map[string]any{
						"title": notice.Title,
						"desc":  notice.Desc,
					})
				}))
			}
		}

		if len(resp.List) < int(pageSize) {
			break
		}
		page++
	}

	return nil
}

// Read 阅读公告信息
func (u *NoticeUseCase) Read(ctx kratosx.Context, id uint32) error {
	md, err := auth.Get(ctx)
	if err != nil {
		return errors.AuthInfoError()
	}
	_ = u.repo.AddNoticeUser(ctx, &NoticeUser{
		UserID:   md.UserId,
		NoticeID: id,
	})
	return nil
}

// Update 更新公告信息
func (u *NoticeUseCase) Update(ctx kratosx.Context, notice *Notice) error {
	if err := u.repo.Update(ctx, notice); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// Delete 删除公告信息
func (u *NoticeUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}
