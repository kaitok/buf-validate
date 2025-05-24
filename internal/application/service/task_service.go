package service

import (
	"go-grpc-api/internal/domain/entity"
	"go-grpc-api/internal/domain/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func (s *TaskService) GetTask(id string) (*entity.Task, error) {
	return s.Repo.FindByID(id)
}
