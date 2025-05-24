package service

import (
	"buf-validate-example/internal/domain/entity"
	"buf-validate-example/internal/domain/repository"
)

type GetTaskUsecase struct {
	TaskRepo repository.TaskRepository
}

func (u *GetTaskUsecase) Execute(id string) (*entity.Task, error) {
	return u.TaskRepo.FindByID(id)
}
