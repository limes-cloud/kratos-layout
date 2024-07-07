package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	"partyaffairs/api/errors"
	"partyaffairs/internal/biz/types"
	"partyaffairs/internal/conf"
)

type ResourceClassify struct {
	Name   string `json:"name" gorm:"not null;size:128;comment:分类名称"`
	Weight uint32 `json:"weight" gorm:"comment:分类权重"`
	ktypes.BaseModel
}

type ResourceContent struct {
	Title            string           `json:"title" gorm:"not null;size:128;comment:文档标题"`
	Desc             string           `json:"desc" gorm:"not null;size:256;comment:文档描述"`
	URL              string           `json:"url" gorm:"not null;size:256;comment:文档url"`
	DownloadCount    uint32           `json:"download_count" gorm:"default 0;comment:下载次数"`
	ClassifyID       uint32           `json:"classify_id" gorm:"not null;comment:资源分类"`
	ResourceClassify ResourceClassify `json:"resource_classify" gorm:"foreignKey:classify_id"`
	ktypes.BaseModel
}

type ResourceRepo interface {
	AllClassify(ctx kratosx.Context) ([]*ResourceClassify, error)
	AddClassify(ctx kratosx.Context, c *ResourceClassify) (uint32, error)
	UpdateClassify(ctx kratosx.Context, c *ResourceClassify) error
	DeleteClassify(ctx kratosx.Context, uint322 uint32) error
	GetContent(ctx kratosx.Context, id uint32) (*ResourceContent, error)
	PageContent(ctx kratosx.Context, req *types.PageResourceContentRequest) ([]*ResourceContent, uint32, error)
	AddContent(ctx kratosx.Context, c *ResourceContent) (uint32, error)
	UpdateContent(ctx kratosx.Context, c *ResourceContent) error
	DeleteContent(ctx kratosx.Context, uint322 uint32) error
}

type ResourceUseCase struct {
	config *conf.Config
	repo   ResourceRepo
}

func NewResourceUseCase(config *conf.Config, repo ResourceRepo) *ResourceUseCase {
	return &ResourceUseCase{config: config, repo: repo}
}

// AllClassify 获取全部资源分类
func (u *ResourceUseCase) AllClassify(ctx kratosx.Context) ([]*ResourceClassify, error) {
	nc, err := u.repo.AllClassify(ctx)
	if err != nil {
		return nil, errors.DatabaseError()
	}
	return nc, nil
}

// AddClassify 添加资源分类信息
func (u *ResourceUseCase) AddClassify(ctx kratosx.Context, nc *ResourceClassify) (uint32, error) {
	id, err := u.repo.AddClassify(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// UpdateClassify 更新资源分类信息
func (u *ResourceUseCase) UpdateClassify(ctx kratosx.Context, nc *ResourceClassify) error {
	if err := u.repo.UpdateClassify(ctx, nc); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteClassify 删除资源分类信息
func (u *ResourceUseCase) DeleteClassify(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteClassify(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// GetContent 获取分页资源
func (u *ResourceUseCase) GetContent(ctx kratosx.Context, id uint32) (*ResourceContent, error) {
	nc, err := u.repo.GetContent(ctx, id)
	if err != nil {
		return nil, errors.DatabaseError()
	}
	return nc, nil
}

// PageContent 获取分页资源
func (u *ResourceUseCase) PageContent(ctx kratosx.Context, req *types.PageResourceContentRequest) ([]*ResourceContent, uint32, error) {
	nc, total, err := u.repo.PageContent(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseError()
	}
	return nc, total, nil
}

// AddContent 添加资源
func (u *ResourceUseCase) AddContent(ctx kratosx.Context, nc *ResourceContent) (uint32, error) {
	id, err := u.repo.AddContent(ctx, nc)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// UpdateContent 更新资源
func (u *ResourceUseCase) UpdateContent(ctx kratosx.Context, nc *ResourceContent) error {
	if err := u.repo.UpdateContent(ctx, nc); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteContent 删除资源
func (u *ResourceUseCase) DeleteContent(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteContent(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}
