package dbs

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type ResourceInfra struct {
}

func NewResourceInfra() *ResourceInfra {
	return &ResourceInfra{}
}

// GetResource 获取指定的文件信息数据
func (r *ResourceInfra) GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error) {
	var (
		ent = entity.Resource{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListResource 获取文件信息列表
func (r *ResourceInfra) ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	var (
		list  []*entity.Resource
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Resource{})
	db = db.Select(fs)

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}

	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateResource 创建文件信息数据
func (r *ResourceInfra) CreateResource(ctx kratosx.Context, ent *entity.Resource) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateResource 更新文件信息数据
func (r *ResourceInfra) UpdateResource(ctx kratosx.Context, ent *entity.Resource) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteResource 删除文件信息数据
func (r *ResourceInfra) DeleteResource(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Resource{}, id).Error
}
