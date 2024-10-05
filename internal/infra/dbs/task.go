package dbs

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

func (u *Task) GetTask(ctx kratosx.Context, id uint32) (*entity.Task, error) {
	nc := entity.Task{}
	return &nc, ctx.DB().First(&nc, "id=?", id).Error
}

func (u *Task) ListTask(ctx kratosx.Context, req *types.ListTaskRequest) ([]*entity.Task, uint32, error) {
	var list []*entity.Task
	var total int64

	db := ctx.DB().Model(entity.Task{}).
		Select([]string{
			"task.id", "title", "description", "is_update", "start", "end", "task.created_at", "task.updated_at",
		})

	if req.Title != nil {
		db.Where("title like ?", *req.Title+"%")
	}
	if req.UserId != nil && req.NotFinish != nil && *req.NotFinish {
		db = db.Joins("TaskValue", ctx.DB().Where("user_id=?", req.UserId)).Where("user_id is null")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *Task) CreateTask(ctx kratosx.Context, task *entity.Task) (uint32, error) {
	return task.Id, ctx.DB().Create(task).Error
}

func (u *Task) UpdateTask(ctx kratosx.Context, task *entity.Task) error {
	return ctx.DB().Updates(task).Error
}

func (u *Task) DeleteTask(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(entity.Task{}, "id=?", id).Error
}

func (u *Task) GetTaskValue(ctx kratosx.Context, taskId, userId uint32) (*entity.TaskValue, error) {
	nc := entity.TaskValue{}
	return &nc, ctx.DB().First(&nc, "task_id=? and user_id=?", taskId, userId).Error
}

func (u *Task) ListTaskValue(ctx kratosx.Context, req *types.ListTaskValueRequest) ([]*entity.TaskValue, uint32, error) {
	var list []*entity.TaskValue
	var total int64

	db := ctx.DB().Model(entity.TaskValue{}).Where("task_id=?", req.TaskId)

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (u *Task) AllTaskValueByTaskId(ctx kratosx.Context, id uint32) ([]*entity.TaskValue, error) {
	var list []*entity.TaskValue
	return list, ctx.DB().Where("task_id=?", id).Find(&list).Error
}

func (u *Task) CreateTaskValue(ctx kratosx.Context, task *entity.TaskValue) (uint32, error) {
	return task.Id, ctx.DB().Create(task).Error
}

func (u *Task) UpdateTaskValue(ctx kratosx.Context, task *entity.TaskValue) error {
	return ctx.DB().Where("user_id=? and task_id=?", task.UserId, task.TaskId).Updates(task).Error
}

func (u *Task) DeleteTaskValue(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(entity.TaskValue{}, "id=?", id).Error
}

func (u *Task) FinishTaskValueUsers(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(entity.TaskValue{}).Select("user_id").Where("task_id=?", id).Scan(&ids).Error
}
