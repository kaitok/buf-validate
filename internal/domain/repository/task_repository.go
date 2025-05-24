package repository

import "buf-validate/internal/domain/entity"

type TaskRepository interface {
	FindByID(id string) (*entity.Task, error)
}
