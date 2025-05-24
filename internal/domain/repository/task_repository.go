package repository

import "buf-validate-example/internal/domain/entity"

type TaskRepository interface {
	FindByID(id string) (*entity.Task, error)
}
