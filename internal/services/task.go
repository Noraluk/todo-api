package services

import (
	"todo-api/internal/domains"
	"todo-api/internal/models"
)

type todoService struct {
	domains.TaskRepository
}

func NewTaskService(todoRepository domains.TaskRepository) domains.TaskService {
	return &todoService{todoRepository}
}

func (s *todoService) GetTasks() ([]models.TaskResponse, error) {
	todos, err := s.TaskRepository.GetTasks()
	if err != nil {
		return nil, err
	}

	res := []models.TaskResponse{}
	for _, v := range todos {
		res = append(res, *v.ToTaskResponse())
	}

	return res, nil
}

func (s *todoService) CreateTask(req *models.TaskRequest) (*models.TaskResponse, error) {
	var err error
	var todo = &models.Task{
		Name: req.Name,
	}

	todo, err = s.TaskRepository.CreateTask(todo)
	if err != nil {
		return nil, err
	}

	return todo.ToTaskResponse(), nil
}

func (s *todoService) DeleteTask(id int) error {
	return s.TaskRepository.DeleteTask(id)
}
