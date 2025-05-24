package repository

import "buf-validate-example/internal/domain/entity"

type UserRepository interface {
	FindByID(id string) (*entity.User, error)
}
