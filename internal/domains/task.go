package domains

import (
	"todo-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

type TaskControler interface {
	GetTasks(c *fiber.Ctx) error
	CreateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}

type TaskService interface {
	GetTasks() ([]models.TaskResponse, error)
	CreateTask(req *models.TaskRequest) (*models.TaskResponse, error)
	DeleteTask(id int) error
}

type TaskRepository interface {
	GetTasks() ([]models.Task, error)
	CreateTask(todo *models.Task) (*models.Task, error)
	DeleteTask(id int) error
}
