package service

import (
	"go-grpc-api/internal/domain/entity"
	"go-grpc-api/internal/domain/repository"
)

type GetTaskUsecase struct {
	TaskRepo repository.TaskRepository
}

func (u *GetTaskUsecase) Execute(id string) (*entity.Task, error) {
	return u.TaskRepo.FindByID(id)
}
