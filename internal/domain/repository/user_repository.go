package repository

import "go-grpc-api/internal/domain/entity"

type UserRepository interface {
	FindByID(id string) (*entity.User, error)
}
