package data

import (
	"fmt"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	biz "layout/internal/biz/dictionary"
	"layout/internal/data/model"
)

type dictionaryRepo struct {
}

func NewDictionaryRepo() biz.Repo {
	return &dictionaryRepo{}
}

// ToDictionaryEntity model转entity
func (r dictionaryRepo) ToDictionaryEntity(m *model.Dictionary) *biz.Dictionary {
	e := &biz.Dictionary{}
	_ = valx.Transform(m, e)
	return e
}

// ToDictionaryModel entity转model
func (r dictionaryRepo) ToDictionaryModel(e *biz.Dictionary) *model.Dictionary {
	m := &model.Dictionary{}
	_ = valx.Transform(e, m)
	return m
}

// GetDictionaryByKeyword 获取指定数据
func (r dictionaryRepo) GetDictionaryByKeyword(ctx kratosx.Context, keyword string) (*biz.Dictionary, error) {
	var (
		m  = model.Dictionary{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToDictionaryEntity(&m), nil
}

// GetDictionary 获取指定的数据
func (r dictionaryRepo) GetDictionary(ctx kratosx.Context, id uint32) (*biz.Dictionary, error) {
	var (
		m  = model.Dictionary{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToDictionaryEntity(&m), nil
}

// ListDictionary 获取列表
func (r dictionaryRepo) ListDictionary(ctx kratosx.Context, req *biz.ListDictionaryRequest) ([]*biz.Dictionary, uint32, error) {
	var (
		bs    []*biz.Dictionary
		ms    []*model.Dictionary
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Dictionary{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToDictionaryEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateDictionary 创建数据
func (r dictionaryRepo) CreateDictionary(ctx kratosx.Context, req *biz.Dictionary) (uint32, error) {
	m := r.ToDictionaryModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateDictionary 更新数据
func (r dictionaryRepo) UpdateDictionary(ctx kratosx.Context, req *biz.Dictionary) error {
	return ctx.DB().Updates(r.ToDictionaryModel(req)).Error
}

// DeleteDictionary 删除数据
func (r dictionaryRepo) DeleteDictionary(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Dictionary{})
	return uint32(db.RowsAffected), db.Error
}
