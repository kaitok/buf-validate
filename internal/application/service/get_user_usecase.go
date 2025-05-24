package service

import (
	"buf-validate/internal/domain/entity"
	"buf-validate/internal/domain/repository"
)

type GetUserUsecase struct {
	UserRepo repository.UserRepository
}

func (u *GetUserUsecase) Execute(id string) (*entity.User, error) {
	return u.UserRepo.FindByID(id)
}
