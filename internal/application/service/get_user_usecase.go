package service

import (
	"buf-validate-example/internal/domain/entity"
	"buf-validate-example/internal/domain/repository"
)

type GetUserUsecase struct {
	UserRepo repository.UserRepository
}

func (u *GetUserUsecase) Execute(id string) (*entity.User, error) {
	return u.UserRepo.FindByID(id)
}
