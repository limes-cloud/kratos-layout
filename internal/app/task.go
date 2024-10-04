package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	pb "partyaffairs/api/partyaffairs/task/v1"
	"partyaffairs/internal/conf"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/domain/service"
	"partyaffairs/internal/infra/dbs"
	"partyaffairs/internal/infra/rpc"
	"partyaffairs/internal/types"
)

type Task struct {
	pb.UnimplementedTaskServer
	srv *service.TaskService
}

func NewTask(conf *conf.Config) *Task {
	return &Task{
		srv: service.NewTaskService(conf, dbs.NewTask(), rpc.NewFile(), rpc.NewUser()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server) {
		srv := NewTask(c)
		pb.RegisterTaskHTTPServer(hs, srv)
	})
}

// ListTask 获取任务列表
func (s *Task) ListTask(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	list, total, err := s.srv.ListTask(kratosx.MustContext(ctx), &types.ListTaskRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Title:    req.Title,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListTaskReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListTaskReply_Task{
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

// ListCurNotFinishTask 获取当前用户未完成的任务列表
func (s *Task) ListCurNotFinishTask(ctx context.Context, req *pb.ListCurNotFinishTaskRequest) (*pb.ListCurNotFinishTaskReply, error) {
	list, total, err := s.srv.ListCurNotFinishTask(kratosx.MustContext(ctx), &types.ListTaskRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Title:    req.Title,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListCurNotFinishTaskReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListCurNotFinishTaskReply_Task{
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
func (s *Task) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	task, err := s.srv.GetTask(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskReply{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		IsUpdate:    task.IsUpdate,
		Start:       task.Start,
		End:         task.End,
		Config:      task.Config,
		CreatedAt:   uint32(task.CreatedAt),
		UpdatedAt:   uint32(task.UpdatedAt),
	}, nil
}

func (s *Task) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskReply, error) {
	id, err := s.srv.CreateTask(kratosx.MustContext(ctx), &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		IsUpdate:    req.IsUpdate,
		Start:       req.Start,
		End:         req.End,
		Config:      req.Config,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskReply{Id: id}, err
}

func (s *Task) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskReply, error) {
	err := s.srv.UpdateTask(kratosx.MustContext(ctx), &entity.Task{
		BaseModel:   ktypes.BaseModel{Id: req.Id},
		Title:       req.Title,
		Description: req.Description,
		IsUpdate:    req.IsUpdate,
		Start:       req.Start,
		End:         req.End,
		Config:      req.Config,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTaskReply{}, err
}

func (s *Task) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskReply, error) {
	return nil, s.srv.DeleteTask(kratosx.MustContext(ctx), req.Id)
}

func (s *Task) ListTaskValue(ctx context.Context, req *pb.ListTaskValueRequest) (*pb.ListTaskValueReply, error) {
	list, total, err := s.srv.ListTaskValue(kratosx.MustContext(ctx), &types.ListTaskValueRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		TaskId:   req.TaskId,
		Finish:   req.Finish,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.ListTaskValueReply{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListTaskValueReply_Value{
			Id:        item.Id,
			TaskId:    item.TaskId,
			UserId:    item.UserId,
			Value:     item.Value,
			CreatedAt: uint32(item.CreatedAt),
			UpdatedAt: uint32(item.UpdatedAt),
			User: &pb.ListTaskValueReply_Value_User{
				Id:        item.User.Id,
				NickName:  item.User.NickName,
				RealName:  item.User.RealName,
				AvatarUrl: item.User.AvatarUrl,
				Phone:     item.User.Phone,
				Email:     item.User.Email,
			},
		})
	}
	return &reply, nil
}

func (s *Task) GetTaskValue(ctx context.Context, in *pb.GetTaskValueRequest) (*pb.GetTaskValueReply, error) {
	task, err := s.srv.GetTaskValue(kratosx.MustContext(ctx), in.TaskId, in.UserId)
	if err != nil {
		return nil, err
	}
	reply := pb.GetTaskValueReply{
		Id:        task.Id,
		TaskId:    task.TaskId,
		UserId:    task.UserId,
		Value:     task.Value,
		CreatedAt: uint32(task.CreatedAt),
		UpdatedAt: uint32(task.UpdatedAt),
		User: &pb.GetTaskValueReply_User{
			Id:        task.User.Id,
			NickName:  task.User.NickName,
			RealName:  task.User.RealName,
			AvatarUrl: task.User.AvatarUrl,
			Phone:     task.User.Phone,
			Email:     task.User.Email,
		},
	}
	return &reply, nil
}

func (s *Task) GetCurTaskValue(ctx context.Context, in *pb.GetCurTaskValueRequest) (*pb.GetCurTaskValueReply, error) {
	task, err := s.srv.GetCurTaskValue(kratosx.MustContext(ctx), in.TaskId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCurTaskValueReply{
		Id:        task.Id,
		TaskId:    task.TaskId,
		UserId:    task.UserId,
		Value:     task.Value,
		CreatedAt: uint32(task.CreatedAt),
		UpdatedAt: uint32(task.UpdatedAt),
	}, nil
}

func (s *Task) CreateTaskValue(ctx context.Context, req *pb.CreateTaskValueRequest) (*pb.CreateTaskValueReply, error) {
	id, err := s.srv.CreateTaskValue(kratosx.MustContext(ctx), &entity.TaskValue{
		TaskId: req.TaskId,
		Value:  req.Value,
	})
	return &pb.CreateTaskValueReply{Id: id}, err
}

func (s *Task) ExportTaskValue(ctx context.Context, in *pb.ExportTaskValueRequest) (*pb.ExportTaskValueReply, error) {
	return &pb.ExportTaskValueReply{}, s.srv.ExportValue(kratosx.MustContext(ctx), in.TaskId)
}

func (s *Task) UpdateTaskValue(ctx context.Context, req *pb.UpdateTaskValueRequest) (*pb.UpdateTaskValueReply, error) {
	return &pb.UpdateTaskValueReply{}, s.srv.UpdateTaskValue(kratosx.MustContext(ctx), &entity.TaskValue{
		TaskId: req.TaskId,
		Value:  req.Value,
	})
}

func (s *Task) DeleteTaskValue(ctx context.Context, in *pb.DeleteTaskValueRequest) (*pb.DeleteTaskValueReply, error) {
	return &pb.DeleteTaskValueReply{}, s.srv.DeleteValue(kratosx.MustContext(ctx), in.Id)
}
