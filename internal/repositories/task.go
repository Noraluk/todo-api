package repositories

import (
	"todo-api/internal/domains"
	"todo-api/internal/models"

	"gorm.io/gorm"
)

type todoRepository struct {
	*gorm.DB
}

func NewTaskRepository(db *gorm.DB) domains.TaskRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetTasks() ([]models.Task, error) {
	todos := []models.Task{}
	err := r.Order("id").Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) CreateTask(todo *models.Task) (*models.Task, error) {
	err := r.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) DeleteTask(id int) error {
	return r.Delete(&models.Task{}, id).Error
}
