package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type TaskRepository interface {
	// GetTask 获取指定的任务
	GetTask(ctx core.Context, id uint32) (*entity.Task, error)

	// ListTask 获取任务列表
	ListTask(ctx core.Context, in *types.ListTaskRequest) ([]*entity.Task, uint32, error)

	// CreateTask 创建定时任务
	CreateTask(ctx core.Context, c *entity.Task) (uint32, error)

	// UpdateTask 更新定时任务
	UpdateTask(ctx core.Context, c *entity.Task) error

	// DeleteTask 删除定时任务
	DeleteTask(ctx core.Context, uint322 uint32) error

	// GetTaskValue 获取定时任务值
	GetTaskValue(ctx core.Context, taskId, userId uint32) (*entity.TaskValue, error)

	// AllTaskValueByTaskId 获取指定任务的所有值
	AllTaskValueByTaskId(ctx core.Context, id uint32) ([]*entity.TaskValue, error)

	// ListTaskValue 获取任务值列表
	ListTaskValue(ctx core.Context, in *types.ListTaskValueRequest) ([]*entity.TaskValue, uint32, error)

	// CreateTaskValue 创建任务值列表
	CreateTaskValue(ctx core.Context, c *entity.TaskValue) (uint32, error)

	// UpdateTaskValue 更新值列表
	UpdateTaskValue(ctx core.Context, c *entity.TaskValue) error

	// DeleteTaskValue 删除值列表
	DeleteTaskValue(ctx core.Context, id uint32) error

	// FinishTaskValueUsers 获取已经完成的用户的id列表
	FinishTaskValueUsers(ctx core.Context, id uint32) ([]uint32, error)
}
