package service

import (
	"buf-validate/internal/domain/entity"
	"buf-validate/internal/domain/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	return s.Repo.FindByID(id)
}
