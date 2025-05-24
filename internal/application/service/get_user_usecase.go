package service

import (
	"go-grpc-api/internal/domain/entity"
	"go-grpc-api/internal/domain/repository"
)

type GetUserUsecase struct {
	UserRepo repository.UserRepository
}

func (u *GetUserUsecase) Execute(id string) (*entity.User, error) {
	return u.UserRepo.FindByID(id)
}
