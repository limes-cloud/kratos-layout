package app

import (
	"context"

	"github.com/limes-cloud/kratosx/model"
	"partyaffairs/api/task"
	"partyaffairs/internal/core"

	"github.com/go-kratos/kratos/v2/transport/http"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Task struct {
	task.UnimplementedTaskServer
	srv *service.TaskService
}

func NewTask() *Task {
	return &Task{
		srv: service.NewTaskService(dbs.NewTask(), rpc.NewFile(), rpc.NewUser()),
	}
}

func init() {
	register(func(hs *http.Server) {
		srv := NewTask()
		task.RegisterTaskHTTPServer(hs, srv)
	})
}

// GetUserPoints 获取任务列表
func (s *Task) GetUserPoints(ctx context.Context, req *task.GetUserPointsRequest) (*task.GetUserPointsResponse, error) {
	total, err := s.srv.GetTaskPoints(core.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	reply := task.GetUserPointsResponse{Points: total}
	return &reply, nil
}

// ListTask 获取任务列表
func (s *Task) ListTask(ctx context.Context, req *task.ListTaskRequest) (*task.ListTaskReply, error) {
	list, total, err := s.srv.ListTask(core.MustContext(ctx), &types.ListTaskRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Title:    req.Title,
	})
	if err != nil {
		return nil, err
	}
	reply := task.ListTaskReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &task.ListTaskReply_Task{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			IsUpdate:    item.IsUpdate,
			Start:       item.Start,
			End:         item.End,
			Config:      item.Config,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// ListClientTask 获取当前用户未完成的任务列表
func (s *Task) ListClientTask(ctx context.Context, req *task.ListClientTaskRequest) (*task.ListClientTaskReply, error) {
	list, total, err := s.srv.ListClientTask(core.MustContext(ctx), &types.ListTaskRequest{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Title:     req.Title,
		NotFinish: req.NotFinish,
	})
	if err != nil {
		return nil, err
	}
	reply := task.ListClientTaskReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &task.ListClientTaskReply_Task{
			Id:          item.Id,
			Title:       item.Title,
			Description: item.Description,
			IsUpdate:    item.IsUpdate,
			Start:       item.Start,
			End:         item.End,
			Config:      item.Config,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// GetTask 获取指定任务
func (s *Task) GetTask(ctx context.Context, in *task.GetTaskRequest) (*task.GetTaskReply, error) {
	res, err := s.srv.GetTask(core.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	return &task.GetTaskReply{
		Id:          res.Id,
		Points:      res.Points,
		Title:       res.Title,
		Description: res.Description,
		IsUpdate:    res.IsUpdate,
		Start:       res.Start,
		End:         res.End,
		Config:      res.Config,
		CreatedAt:   uint32(res.CreatedAt),
		UpdatedAt:   uint32(res.UpdatedAt),
	}, nil
}

func (s *Task) CreateTask(ctx context.Context, req *task.CreateTaskRequest) (*task.CreateTaskReply, error) {
	id, err := s.srv.CreateTask(core.MustContext(ctx), &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Points:      req.Points,
		IsUpdate:    req.IsUpdate,
		Start:       req.Start,
		End:         req.End,
		Config:      req.Config,
	})
	if err != nil {
		return nil, err
	}
	return &task.CreateTaskReply{Id: id}, err
}

func (s *Task) UpdateTask(ctx context.Context, req *task.UpdateTaskRequest) (*task.UpdateTaskReply, error) {
	err := s.srv.UpdateTask(core.MustContext(ctx), &entity.Task{
		BaseTenantModel: model.BaseTenantModel{Id: req.Id},
		Title:           req.Title,
		Points:          req.Points,
		Description:     req.Description,
		IsUpdate:        req.IsUpdate,
		Start:           req.Start,
		End:             req.End,
		Config:          req.Config,
	})
	if err != nil {
		return nil, err
	}
	return &task.UpdateTaskReply{}, err
}

func (s *Task) DeleteTask(ctx context.Context, req *task.DeleteTaskRequest) (*task.DeleteTaskReply, error) {
	return nil, s.srv.DeleteTask(core.MustContext(ctx), req.Id)
}

func (s *Task) ListTaskValue(ctx context.Context, req *task.ListTaskValueRequest) (*task.ListTaskValueReply, error) {
	list, total, err := s.srv.ListTaskValue(core.MustContext(ctx), &types.ListTaskValueRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		TaskId:   req.TaskId,
		Finish:   req.Finish,
	})
	if err != nil {
		return nil, err
	}
	reply := task.ListTaskValueReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &task.ListTaskValueReply_Value{
			Id:        item.Id,
			TaskId:    item.TaskId,
			UserId:    item.UserId,
			Value:     item.Value,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
			User: &task.ListTaskValueReply_Value_User{
				Id:       item.User.Id,
				NickName: item.User.Nickname,
				Username: &item.User.Username,
				Avatar:   &item.User.Avatar,
			},
		})
	}
	return &reply, nil
}

func (s *Task) GetTaskValue(ctx context.Context, in *task.GetTaskValueRequest) (*task.GetTaskValueReply, error) {
	res, err := s.srv.GetTaskValue(core.MustContext(ctx), in.TaskId, in.UserId)
	if err != nil {
		return nil, err
	}
	reply := task.GetTaskValueReply{
		Id:        res.Id,
		TaskId:    res.TaskId,
		UserId:    res.UserId,
		Value:     res.Value,
		CreatedAt: uint32(res.CreatedAt),
		UpdatedAt: uint32(res.UpdatedAt),
		User: &task.GetTaskValueReply_User{
			Id:       res.User.Id,
			Username: res.User.Username,
			Avatar:   &res.User.Avatar,
		},
	}
	return &reply, nil
}

func (s *Task) GetCurTaskValue(ctx context.Context, in *task.GetCurTaskValueRequest) (*task.GetCurTaskValueReply, error) {
	res, err := s.srv.GetCurTaskValue(core.MustContext(ctx), in.TaskId)
	if err != nil {
		return nil, err
	}
	return &task.GetCurTaskValueReply{
		Id:        res.Id,
		TaskId:    res.TaskId,
		UserId:    res.UserId,
		Value:     res.Value,
		CreatedAt: uint32(res.CreatedAt),
		UpdatedAt: uint32(res.UpdatedAt),
	}, nil
}

func (s *Task) CreateTaskValue(ctx context.Context, req *task.CreateTaskValueRequest) (*task.CreateTaskValueReply, error) {
	id, err := s.srv.CreateTaskValue(core.MustContext(ctx), &entity.TaskValue{
		TaskId: req.TaskId,
		Value:  req.Value,
	})
	return &task.CreateTaskValueReply{Id: id}, err
}

func (s *Task) ExportTaskValue(ctx context.Context, in *task.ExportTaskValueRequest) (*task.ExportTaskValueReply, error) {
	id, err := s.srv.ExportValue(core.MustContext(ctx), in.TaskId)
	if err != nil {
		return &task.ExportTaskValueReply{}, err
	}
	return &task.ExportTaskValueReply{Id: id}, nil
}

func (s *Task) UpdateTaskValue(ctx context.Context, req *task.UpdateTaskValueRequest) (*task.UpdateTaskValueReply, error) {
	return &task.UpdateTaskValueReply{}, s.srv.UpdateTaskValue(core.MustContext(ctx), &entity.TaskValue{
		TaskId: req.TaskId,
		Value:  req.Value,
	})
}

func (s *Task) DeleteTaskValue(ctx context.Context, in *task.DeleteTaskValueRequest) (*task.DeleteTaskValueReply, error) {
	return &task.DeleteTaskValueReply{}, s.srv.DeleteValue(core.MustContext(ctx), in.Id)
}
