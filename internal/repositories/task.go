package repositories

import (
	"todo-api/internal/models"
	"todo-api/pkg/base"
)

var TaskRepository base.Repository[models.Task] = base.Repository[models.Task]{}

func init() {
	TaskRepository.Init([]base.Association{})
}
