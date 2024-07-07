package data

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/biz"
	"partyaffairs/internal/biz/types"
)

type taskRepo struct {
}

func NewTaskRepo() biz.TaskRepo {
	return &taskRepo{}
}

func (u *taskRepo) Get(ctx kratosx.Context, id uint32) (*biz.Task, error) {
	nc := biz.Task{}
	return &nc, ctx.DB().First(&nc, "id=?", id).Error
}

func (u *taskRepo) Page(ctx kratosx.Context, req *types.PageTaskRequest) ([]*biz.Task, uint32, error) {
	var list []*biz.Task
	var total int64

	db := ctx.DB().Model(biz.Task{}).
		Select("task.id", "title", "desc", "is_update", "start", "end", "task.created_at", "task.updated_at")

	if req.UserID != nil {
		db = db.Preload("TaskValue", ctx.DB().Where("user_id=?", req.UserID))
		db = db.Joins("TaskValue", ctx.DB().Where("user_id=?", req.UserID))
		if req.NotFinish != nil && *req.NotFinish {
			db = db.Where("user_id is null")
		}
	}

	if req.Title != nil {
		db.Where("title like ?", *req.Title+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *taskRepo) Create(ctx kratosx.Context, task *biz.Task) (uint32, error) {
	return task.Id, ctx.DB().Create(task).Error
}

func (u *taskRepo) Update(ctx kratosx.Context, task *biz.Task) error {
	return ctx.DB().Where("id=?", task.Id).Updates(task).Error
}

func (u *taskRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Task{}, "id=?", id).Error
}

func (u *taskRepo) GetValue(ctx kratosx.Context, taskId, userId uint32) (*biz.TaskValue, error) {
	nc := biz.TaskValue{}
	return &nc, ctx.DB().First(&nc, "task_id=? and user_id=?", taskId, userId).Error
}

func (u *taskRepo) PageValue(ctx kratosx.Context, req *types.PageTaskValueRequest) ([]*biz.TaskValue, uint32, error) {
	var list []*biz.TaskValue
	var total int64

	db := ctx.DB().Model(biz.TaskValue{}).Where("task_id=?", req.TaskID)

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *taskRepo) AllValueByTaskId(ctx kratosx.Context, id uint32) ([]*biz.TaskValue, error) {
	var list []*biz.TaskValue
	return list, ctx.DB().Where("task_id=?", id).Find(&list).Error
}

func (u *taskRepo) CreateValue(ctx kratosx.Context, task *biz.TaskValue) (uint32, error) {
	return task.Id, ctx.DB().Create(task).Error
}

func (u *taskRepo) UpdateValue(ctx kratosx.Context, task *biz.TaskValue) error {
	return ctx.DB().Where("user_id=? and task_id=?", task.UserID, task.TaskID).Updates(task).Error
}

func (u *taskRepo) DeleteValue(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.TaskValue{}, "id=?", id).Error
}
