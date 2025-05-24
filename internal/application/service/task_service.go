package service

import (
	"buf-validate-example/internal/domain/entity"
	"buf-validate-example/internal/domain/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func (s *TaskService) GetTask(id string) (*entity.Task, error) {
	return s.Repo.FindByID(id)
}
