package service

import (
	"encoding/json"
	"fmt"

	"github.com/limes-cloud/application/api/application/auth"
	"github.com/limes-cloud/kratosx"
	mauth "github.com/limes-cloud/manager/api/manager/auth"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/repository"
	"partyaffairs/internal/types"
)

const (
	exportScene = "party-affairs-task"
)

type TaskService struct {
	conf *conf.Config
	repo repository.TaskRepository
	file repository.FileRepository
	user repository.UserRepository
}

func NewTaskService(conf *conf.Config, repo repository.TaskRepository, file repository.FileRepository, user repository.UserRepository) *TaskService {
	return &TaskService{conf: conf, repo: repo, file: file, user: user}
}

// GetTask 获取指定的公告
func (u *TaskService) GetTask(ctx kratosx.Context, id uint32) (*entity.Task, error) {
	task, err := u.repo.GetTask(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return task, nil
}

// ListTask 获取分页任务
func (u *TaskService) ListTask(ctx kratosx.Context, req *types.ListTaskRequest) ([]*entity.Task, uint32, error) {
	task, total, err := u.repo.ListTask(ctx, req)
	if err != nil {
		return nil, total, errors.DatabaseError()
	}
	return task, total, nil
}

// ListClientTask 获取分页任务
func (u *TaskService) ListClientTask(ctx kratosx.Context, req *types.ListTaskRequest) ([]*entity.Task, uint32, error) {
	md, err := auth.Get(ctx)
	if err != nil {
		return nil, 0, errors.SystemError()
	}
	req.UserId = &md.UserId
	task, total, err := u.repo.ListTask(ctx, req)
	if err != nil {
		return nil, total, errors.DatabaseError()
	}
	return task, total, nil
}

// CreateTask 添加任务信息
func (u *TaskService) CreateTask(ctx kratosx.Context, task *entity.Task) (uint32, error) {
	id, err := u.repo.CreateTask(ctx, task)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateTask 更新任务信息
func (u *TaskService) UpdateTask(ctx kratosx.Context, task *entity.Task) error {
	if err := u.repo.UpdateTask(ctx, task); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteTask 删除任务信息
func (u *TaskService) DeleteTask(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteTask(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// GetTaskValue 获取指定的值
func (u *TaskService) GetTaskValue(ctx kratosx.Context, taskId, userId uint32) (*entity.TaskValue, error) {
	task, _ := u.repo.GetTaskValue(ctx, taskId, userId)
	return task, nil
}

// GetCurTaskValue 获取指定的值
func (u *TaskService) GetCurTaskValue(ctx kratosx.Context, taskId uint32) (*entity.TaskValue, error) {
	md, err := auth.Get(ctx)
	if err != nil {
		return nil, errors.SystemError()
	}
	task, _ := u.repo.GetTaskValue(ctx, taskId, md.UserId)
	return task, nil
}

// ListTaskValue 获取分页任务
func (u *TaskService) ListTaskValue(ctx kratosx.Context, req *types.ListTaskValueRequest) ([]*entity.TaskValue, uint32, error) {
	if req.Finish != nil && !*req.Finish {
		// 获取已经填写完成的用户id列表
		ids, err := u.repo.FinishTaskValueUsers(ctx, req.TaskId)
		if err != nil {
			return nil, 0, errors.ListUserError(err.Error())
		}

		users, total, err := u.user.ListUser(ctx, &types.ListUserRequest{
			Page:     req.Page,
			PageSize: req.PageSize,
			NotIn:    ids,
		})
		if err != nil {
			return nil, 0, errors.ListError(err.Error())
		}

		var values []*entity.TaskValue
		for _, user := range users {
			values = append(values, &entity.TaskValue{
				TaskId: req.TaskId,
				UserId: user.Id,
				User:   user,
			})
		}
		return values, total, nil
	}

	list, total, err := u.repo.ListTaskValue(ctx, req)
	if err != nil {
		return nil, total, errors.DatabaseError()
	}
	for ind, item := range list {
		user, _ := u.user.GetUser(ctx, item.UserId)
		list[ind].User = user
	}
	return list, total, nil
}

// CreateTaskValue 添加任务信息
func (u *TaskService) CreateTaskValue(ctx kratosx.Context, task *entity.TaskValue) (uint32, error) {
	md, err := auth.Get(ctx)
	if err != nil {
		return 0, errors.SystemError()
	}
	task.UserId = md.UserId

	id, err := u.repo.CreateTaskValue(ctx, task)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}
	return id, nil
}

// ExportValue 导出任务信息
func (u *TaskService) ExportValue(ctx kratosx.Context, id uint32) (uint32, error) {
	// 获取当前用户信息
	info, err := mauth.GetAuthInfo(ctx)
	if err != nil {
		return 0, errors.SystemError()
	}

	task, err := u.repo.GetTask(ctx, id)
	if err != nil {
		return 0, errors.DatabaseError(err.Error())
	}

	cfg := []struct {
		Type   string `json:"type"`
		Field  string `json:"field"`
		Config struct {
			Label string `json:"label"`
		} `json:"config"`
	}{{}}

	// 查询表头
	if err := json.Unmarshal([]byte(task.Config), &cfg); err != nil {
		return 0, errors.TransformError()
	}

	var (
		tps     = map[string]string{}
		headers = []string{"用户ID", "用户名"}
		dirs    = map[string]string{}
		files   []*types.ExportFileItem
		rows    [][]*types.ExportExcelCol
	)
	for _, item := range cfg {
		tp := "string"
		if item.Type == "upload" {
			tp = "file"
			dirs[item.Field] = item.Config.Label
		} else {
			headers = append(headers, item.Config.Label)
		}
		tps[item.Field] = tp
	}

	// 获取数据
	list, err := u.repo.AllTaskValueByTaskId(ctx, id)
	for _, item := range list {
		var (
			value = make(map[string]string)
			cols  []*types.ExportExcelCol
		)
		if err := json.Unmarshal([]byte(item.Value), &value); err != nil {
			continue
		}

		user, err := u.user.GetUser(ctx, item.UserId)
		if err != nil {
			return 0, err
		}

		// 写入用户信息
		cols = append(cols, []*types.ExportExcelCol{
			{
				Type:  "string",
				Value: fmt.Sprint(user.Id),
			},
			{
				Type:  "string",
				Value: user.GetName(),
			},
		}...)

		// 写入表格数据
		for _, ite := range cfg {
			if tps[ite.Field] == "file" {
				files = append(files, &types.ExportFileItem{
					Rename: fmt.Sprintf("%s/%s-%d", dirs[ite.Field], user.GetName(), user.Id),
					Value:  value[ite.Field],
				})
			} else {
				cols = append(cols, &types.ExportExcelCol{
					Type:  tps[ite.Field],
					Value: value[ite.Field],
				})
			}
		}

		// 追加进入行数据
		rows = append(rows, cols)
	}

	id, err = u.file.ExportExcel(ctx, &types.ExportExcelRequest{
		UserId:       info.UserId,
		DepartmentId: info.DepartmentId,
		Scene:        exportScene,
		Name:         task.Title,
		Rows:         rows,
		Files:        files,
		Headers:      headers,
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateTaskValue 更新任务信息
func (u *TaskService) UpdateTaskValue(ctx kratosx.Context, task *entity.TaskValue) error {
	md, err := auth.Get(ctx)
	if err != nil {
		return errors.SystemError()
	}

	task.UserId = md.UserId
	if err := u.repo.UpdateTaskValue(ctx, task); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteValue 删除任务信息
func (u *TaskService) DeleteValue(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteTaskValue(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
