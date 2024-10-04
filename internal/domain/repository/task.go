package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type TaskRepository interface {
	// GetTask 获取指定的任务
	GetTask(ctx kratosx.Context, id uint32) (*entity.Task, error)

	// ListTask 获取任务列表
	ListTask(ctx kratosx.Context, in *types.ListTaskRequest) ([]*entity.Task, uint32, error)

	// CreateTask 创建定时任务
	CreateTask(ctx kratosx.Context, c *entity.Task) (uint32, error)

	// UpdateTask 更新定时任务
	UpdateTask(ctx kratosx.Context, c *entity.Task) error

	// DeleteTask 删除定时任务
	DeleteTask(ctx kratosx.Context, uint322 uint32) error

	// GetTaskValue 获取定时任务值
	GetTaskValue(ctx kratosx.Context, taskId, userId uint32) (*entity.TaskValue, error)

	// AllTaskValueByTaskId 获取指定任务的所有值
	AllTaskValueByTaskId(ctx kratosx.Context, id uint32) ([]*entity.TaskValue, error)

	// ListTaskValue 获取任务值列表
	ListTaskValue(ctx kratosx.Context, in *types.ListTaskValueRequest) ([]*entity.TaskValue, uint32, error)

	// CreateTaskValue 创建任务值列表
	CreateTaskValue(ctx kratosx.Context, c *entity.TaskValue) (uint32, error)

	// UpdateTaskValue 更新值列表
	UpdateTaskValue(ctx kratosx.Context, c *entity.TaskValue) error

	// DeleteTaskValue 删除值列表
	DeleteTaskValue(ctx kratosx.Context, id uint32) error

	// FinishTaskValueUsers 获取已经完成的用户的id列表
	FinishTaskValueUsers(ctx kratosx.Context, id uint32) ([]uint32, error)
}
