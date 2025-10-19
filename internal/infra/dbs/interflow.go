package dbs

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Interflow struct {
}

func NewInterflow() *Interflow {
	return &Interflow{}
}

// ListInterflowPerson 获取通知信息列表
func (t *Interflow) ListInterflowPerson(ctx core.Context) ([]*entity.InterflowPerson, error) {
	var list []*entity.InterflowPerson
	return list, ctx.DB().Model(entity.InterflowPerson{}).Find(&list).Error
}

// CreateInterflowPerson 创建通知信息
func (t *Interflow) CreateInterflowPerson(ctx core.Context, req *entity.InterflowPerson) (uint32, error) {
	return req.Id, ctx.DB().Create(req).Error
}

// DeleteInterflowPerson 删除通知信息
func (t *Interflow) DeleteInterflowPerson(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.InterflowPerson{}).Error
}

// ListInterflow 获取通知信息列表
func (r *Interflow) ListInterflow(ctx core.Context, req *types.ListInterflowRequest) ([]*entity.Interflow, uint32, error) {
	var (
		list  []*entity.Interflow
		total int64
	)

	ids := []uint32{req.FromUserId, req.ToUserId}
	db := ctx.DB().Model(entity.Interflow{}).
		Where("from_user_id in ?", ids).
		Where("to_user_id in ?", ids)

	db = db.Order("id desc")
	//if req.Search != nil {
	//	req.Order = value.Pointer("desc")
	//	req.OrderBy = value.Pointer("id")
	//	if err := db.Count(&total).Error; err != nil {
	//		return nil, 0, err
	//	}
	//
	//	db = page.SearchScopes(db, req.Search)
	//}

	return list, uint32(total), db.Find(&list).Error
}

// CreateInterflow 创建通知信息数据
func (r *Interflow) CreateInterflow(ctx core.Context, ent *entity.Interflow) (uint32, error) {
	return ent.Id, ctx.DB().Create(ent).Error
}

// UpdateInterflow 更新通知信息数据
func (r *Interflow) UpdateInterflow(ctx core.Context, ent *entity.Interflow) error {
	return ctx.DB().Updates(ent).Error
}

// DeleteInterflow 删除通知信息数据
func (r *Interflow) DeleteInterflow(ctx core.Context, id uint32) error {
	return ctx.DB().Delete(&entity.Interflow{}, id).Error
}
