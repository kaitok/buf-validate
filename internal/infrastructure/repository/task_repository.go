package repository

import "go-grpc-api/internal/domain/entity"

type TaskRepositoryImpl struct{}

func (r *TaskRepositoryImpl) FindByID(id string) (*entity.Task, error) {
	return &entity.Task{ID: id, Title: "dummy", UserID: "1"}, nil
}
