package repository

import "go-grpc-api/internal/domain/entity"

type TaskRepository interface {
	FindByID(id string) (*entity.Task, error)
}
