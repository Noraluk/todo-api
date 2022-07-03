package services

import (
	"todo-api/internal/models"
	"todo-api/internal/repositories"
	"todo-api/pkg/base"
)

type TaskService interface {
	GetTasks(searchModel models.TaskSearchModel) ([]models.Task, error)
	CreateTask(req *models.TaskRequest) (*models.Task, error)
	DeleteTask(id int) error
}

type todoService struct {
	taskRepository base.Repository[models.Task]
}

func NewTaskService() TaskService {
	return &todoService{
		taskRepository: repositories.TaskRepository,
	}
}

func (s *todoService) GetTasks(searchModel models.TaskSearchModel) ([]models.Task, error) {
	todos, db := s.taskRepository.FindPaginate(map[string]any{}, int(searchModel.Page), int(searchModel.Limit))
	return todos, db.Error
}

func (s *todoService) CreateTask(req *models.TaskRequest) (*models.Task, error) {
	todo, db := s.taskRepository.Insert(&models.Task{Name: req.Name})
	return todo, db.Error
}

func (s *todoService) DeleteTask(id int) error {
	_, db := s.taskRepository.Delete(&models.Task{Id: id})
	return db.Error
}
