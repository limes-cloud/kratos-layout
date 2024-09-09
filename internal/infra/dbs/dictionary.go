package dbs

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"

	"layout/internal/domain/entity"
	"layout/internal/types"
)

type Dictionary struct {
	// db 基础设施中间件
}

func NewDictionary() *Dictionary {
	return &Dictionary{}
}

// ListDictionary 获取列表
func (r Dictionary) ListDictionary(ctx kratosx.Context, req *types.ListDictionaryRequest) ([]*entity.Dictionary, uint32, error) {
	var (
		es    []*entity.Dictionary
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

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

	return es, uint32(total), db.Find(&es).Error
}

// CreateDictionary 创建数据
func (r Dictionary) CreateDictionary(ctx kratosx.Context, ent *entity.Dictionary) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

func (r Dictionary) GetDictionaryByKeyword(ctx kratosx.Context, keyword string) (*entity.Dictionary, error) {
	var (
		m  = entity.Dictionary{}
		fs = []string{"*"}
	)

	return &m, ctx.DB().Select(fs).First(&m, "keyword=?", keyword).Error
}

// GetDictionary 获取指定的数据
func (r Dictionary) GetDictionary(ctx kratosx.Context, id uint32) (*entity.Dictionary, error) {
	var (
		m  = entity.Dictionary{}
		fs = []string{"*"}
	)

	return &m, ctx.DB().Select(fs).First(&m, id).Error
}

// UpdateDictionary 更新数据
func (r Dictionary) UpdateDictionary(ctx kratosx.Context, req *entity.Dictionary) error {
	return ctx.DB().Updates(req).Error
}

// DeleteDictionary 删除数据
func (r Dictionary) DeleteDictionary(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id=?", id).Delete(&entity.Dictionary{}).Error
}
