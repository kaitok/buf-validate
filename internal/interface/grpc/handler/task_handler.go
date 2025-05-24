package handler

import (
	"buf-validate-example/internal/application/service"
	"buf-validate-example/internal/infrastructure/repository"
	taskpb "buf-validate-example/tools/grpc/v1/task"
	"context"
)

var getTaskUsecase = service.GetTaskUsecase{
	TaskRepo: &repository.TaskRepositoryImpl{},
}

type TaskServiceServer struct {
	taskpb.UnimplementedTaskServiceServer
}

func (s *TaskServiceServer) GetTaskByUser(ctx context.Context, req *taskpb.GetTasksByUserRequest) (*taskpb.GetTasksByUserResponse, error) {
	tasks := []*taskpb.Task{}
	task, err := getTaskUsecase.Execute(string(rune(req.GetUserId())))
	if err != nil {
		return nil, err
	}
	tasks = append(tasks, &taskpb.Task{
		Id:     int32(1),
		Title:  task.Title,
		UserId: task.UserID,
	})
	return &taskpb.GetTasksByUserResponse{Tasks: tasks}, nil
}
