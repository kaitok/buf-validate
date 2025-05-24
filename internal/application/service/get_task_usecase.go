package service

import (
	"buf-validate/internal/domain/entity"
	"buf-validate/internal/domain/repository"
)

type GetTaskUsecase struct {
	TaskRepo repository.TaskRepository
}

func (u *GetTaskUsecase) Execute(id string) (*entity.Task, error) {
	return u.TaskRepo.FindByID(id)
}
