package repository

import "buf-validate/internal/domain/entity"

type UserRepository interface {
	FindByID(id string) (*entity.User, error)
}
