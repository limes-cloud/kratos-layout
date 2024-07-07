package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/usercenter/api/usercenter/auth"

	"partyaffairs/api/errors"
	"partyaffairs/internal/biz/types"
	"partyaffairs/internal/conf"
)

type NewsClassify struct {
	Name   string `json:"name" gorm:"not null;size:128;comment:分类名称"`
	Weight uint32 `json:"weight" gorm:"comment:分类权重"`
	ktypes.BaseModel
}

type NewsContent struct {
	Title      string         `json:"title" gorm:"not null;size:128;comment:新闻标题"`
	Unit       string         `json:"unit" gorm:"not null;size:128;comment:发布单位"`
	Cover      string         `json:"cover" gorm:"not null;size:128;comment:封面图片"`
	Desc       string         `json:"desc" gorm:"not null;size:128;comment:新闻描述"`
	Content    string         `json:"content" gorm:"not null;type:blob;comment:新闻内容"`
	Read       uint32         `json:"read" gorm:"default 0;comment:阅读人数"`
	IsTop      *bool          `json:"is_top" gorm:"default false;comment:是否置顶"`
	ClassifyID uint32         `json:"classify_id" gorm:"not null;comment:新闻分类"`
	Classify   NewsClassify   `json:"classify" gorm:"foreignKey:classify_id"`
	Comments   []*NewsComment `json:"comments" gorm:"foreignKey:content_id"`
	ktypes.BaseModel
}

type NewsComment struct {
	ContentID uint32       `json:"content_id" gorm:"not null;comment:新闻ID"`
	FromUid   uint32       `json:"from_uid" gorm:"not null;comment:发布者ID"`
	ReplyUid  uint32       `json:"reply_uid" gorm:"comment:回复者ID"`
	Text      string       `json:"text" gorm:"size:1024;comment:回复文本"`
	Content   *NewsContent `json:"content" gorm:"foreignKey:content_id;constraint:onDelete:cascade"`
	ktypes.CreateModel
}

type NewsRepo interface {
	AllClassify(ctx kratosx.Context) ([]*NewsClassify, error)
	AddClassify(ctx kratosx.Context, c *NewsClassify) (uint32, error)
	UpdateClassify(ctx kratosx.Context, c *NewsClassify) error
	DeleteClassify(ctx kratosx.Context, uint322 uint32) error
	GetContent(ctx kratosx.Context, id uint32) (*NewsContent, error)
	PageContent(ctx kratosx.Context, req *types.PageNewsContentRequest) ([]*NewsContent, uint32, error)
	AddContent(ctx kratosx.Context, c *NewsContent) (uint32, error)
	UpdateContent(ctx kratosx.Context, c *NewsContent) error
	AddContentReadCount(ctx kratosx.Context, id uint32) error
	DeleteContent(ctx kratosx.Context, uint322 uint32) error
	GetComment(ctx kratosx.Context, uid uint32) (*NewsComment, error)
	AddComment(ctx kratosx.Context, comment *NewsComment) (uint32, error)
	DeleteComment(ctx kratosx.Context, id uint32) error
	PageComment(ctx kratosx.Context, req *types.PageNewsCommentRequest) ([]*NewsComment, uint32, error)
}

type NewsUseCase struct {
	config *conf.Config
	repo   NewsRepo
}

func NewNewsUseCase(config *conf.Config, repo NewsRepo) *NewsUseCase {
	return &NewsUseCase{config: config, repo: repo}
}

// AllClassify 获取全部新闻分类
func (u *NewsUseCase) AllClassify(ctx kratosx.Context) ([]*NewsClassify, error) {
	nc, err := u.repo.AllClassify(ctx)
	if err != nil {
		return nil, errors.DatabaseError()
	}
	return nc, nil
}

// AddClassify 添加新闻分类信息
func (u *NewsUseCase) AddClassify(ctx kratosx.Context, nc *NewsClassify) (uint32, error) {
	id, err := u.repo.AddClassify(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// UpdateClassify 更新新闻分类信息
func (u *NewsUseCase) UpdateClassify(ctx kratosx.Context, nc *NewsClassify) error {
	if err := u.repo.UpdateClassify(ctx, nc); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteClassify 删除新闻分类信息
func (u *NewsUseCase) DeleteClassify(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteClassify(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// GetContent 获取分页新闻
func (u *NewsUseCase) GetContent(ctx kratosx.Context, id uint32) (*NewsContent, error) {
	nc, err := u.repo.GetContent(ctx, id)
	if err != nil {
		return nil, errors.DatabaseError()
	}

	if err = u.repo.AddContentReadCount(ctx, nc.Id); err != nil {
		return nil, errors.DatabaseError()
	}
	return nc, nil
}

// PageContent 获取分页新闻
func (u *NewsUseCase) PageContent(ctx kratosx.Context, req *types.PageNewsContentRequest) ([]*NewsContent, uint32, error) {
	nc, total, err := u.repo.PageContent(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseError()
	}
	return nc, total, nil
}

// AddContent 添加新闻
func (u *NewsUseCase) AddContent(ctx kratosx.Context, nc *NewsContent) (uint32, error) {
	id, err := u.repo.AddContent(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// UpdateContent 更新新闻
func (u *NewsUseCase) UpdateContent(ctx kratosx.Context, nc *NewsContent) error {
	if err := u.repo.UpdateContent(ctx, nc); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteContent 删除新闻
func (u *NewsUseCase) DeleteContent(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteContent(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// PageComment 分页查看评论
func (u *NewsUseCase) PageComment(ctx kratosx.Context, req *types.PageNewsCommentRequest) ([]*NewsComment, uint32, error) {
	list, total, err := u.repo.PageComment(ctx, req)
	if err != nil {
		return nil, total, errors.DatabaseError(err.Error())
	}
	return list, total, nil
}

// AddComment 添加评论
func (u *NewsUseCase) AddComment(ctx kratosx.Context, nc *NewsComment) (uint32, error) {
	md, err := auth.Get(ctx)
	if err != nil {
		return 0, errors.AuthInfoError()
	}
	nc.FromUid = md.UserId
	id, err := u.repo.AddComment(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// DeleteComment 删除评论
func (u *NewsUseCase) DeleteComment(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteComment(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteCurComment 删除用户当前评论
func (u *NewsUseCase) DeleteCurComment(ctx kratosx.Context, id uint32) error {
	md, err := auth.Get(ctx)
	if err != nil {
		return errors.AuthInfoError()
	}

	comment, err := u.repo.GetComment(ctx, id)
	if err != nil {
		return errors.DatabaseError()
	}

	if md.UserId != comment.FromUid {
		return errors.DatabaseError()
	}

	if err := u.repo.DeleteComment(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}
