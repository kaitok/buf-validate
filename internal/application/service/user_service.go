package service

import (
	"go-grpc-api/internal/domain/entity"
	"go-grpc-api/internal/domain/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) GetUser(id string) (*entity.User, error) {
	return s.Repo.FindByID(id)
}
