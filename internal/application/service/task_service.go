package service

import (
	"buf-validate/internal/domain/entity"
	"buf-validate/internal/domain/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func (s *TaskService) GetTask(id string) (*entity.Task, error) {
	return s.Repo.FindByID(id)
}
