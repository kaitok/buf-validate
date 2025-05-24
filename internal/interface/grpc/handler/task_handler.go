package handler

import (
	"context"
	"go-grpc-api/internal/application/service"
	"go-grpc-api/internal/infrastructure/repository"
	taskpb "go-grpc-api/tools/grpc/go_grpc_api/v1"

	"github.com/bufbuild/connect-go"
)

var getTaskUsecase = service.GetTaskUsecase{
	TaskRepo: &repository.TaskRepositoryImpl{},
}

type TaskServiceServer struct{}

func (s *TaskServiceServer) GetTaskByUser(ctx context.Context, req *connect.Request[taskpb.GetTaskByUserRequest]) (*connect.Response[taskpb.GetTaskByUserResponse], error) {
	tasks := []*taskpb.Task{}
	task, err := getTaskUsecase.Execute(string(rune(req.Msg.GetUserId())))
	if err != nil {
		return nil, err
	}
	tasks = append(tasks, &taskpb.Task{
		Id:     int32(1),
		Title:  task.Title,
		UserId: task.UserID,
	})
	return connect.NewResponse(&taskpb.GetTaskByUserResponse{Tasks: tasks}), nil
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *connect.Request[taskpb.CreateTaskRequest]) (*connect.Response[taskpb.CreateTaskResponse], error) {
	// 実装例: 空のレスポンスを返す
	return connect.NewResponse(&taskpb.CreateTaskResponse{}), nil
}
