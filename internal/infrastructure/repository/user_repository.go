package repository

import "buf-validate-example/internal/domain/entity"

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) FindByID(id string) (*entity.User, error) {
	return &entity.User{ID: id, Name: "dummy"}, nil
}
