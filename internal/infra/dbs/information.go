package dbs

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"poverty/internal/domain/entity"
	"poverty/internal/types"
)

type InformationInfra struct {
}

func NewInformationInfra() *InformationInfra {
	return &InformationInfra{}
}

// GetInformation 获取指定的资讯信息数据
func (r *InformationInfra) GetInformation(ctx kratosx.Context, id uint32) (*entity.Information, error) {
	var (
		ent = entity.Information{}
		fs  = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	return &ent, db.First(&ent, id).Error
}

// ListInformation 获取资讯信息列表
func (r *InformationInfra) ListInformation(ctx kratosx.Context, req *types.ListInformationRequest) ([]*entity.Information, uint32, error) {
	var (
		list  []*entity.Information
		fs    = []string{"*"}
		total int64
	)

	db := ctx.DB().Model(entity.Information{})
	db = db.Select(fs)

	if req.Title != nil {
		db = db.Where("title LIKE ?", *req.Title+"%")
	}
	if req.IsTop != nil {
		db = db.Where("is_top = ?", *req.IsTop)
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
		req.Order = proto.String("desc")
	}

	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id desc")
	}

	return list, uint32(total), db.Find(&list).Error
}

// CreateInformation 创建资讯信息数据
func (r *InformationInfra) CreateInformation(ctx kratosx.Context, ent *entity.Information) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateInformation 更新资讯信息数据
func (r *InformationInfra) UpdateInformation(ctx kratosx.Context, ent *entity.Information) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteInformation 删除资讯信息数据
func (r *InformationInfra) DeleteInformation(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Information{}, id).Error
}
