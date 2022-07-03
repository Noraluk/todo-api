package controllers

import (
	"net/http"
	"strconv"
	"todo-api/internal/models"
	"todo-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TaskControler interface {
	GetTasks(c *fiber.Ctx) error
	CreateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}

type taskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) TaskControler {
	return &taskController{taskService: taskService}
}

func (ct *taskController) GetTasks(c *fiber.Ctx) error {
	var searchModel models.TaskSearchModel
	err := c.QueryParser(&searchModel)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	tasks, err := ct.taskService.GetTasks(searchModel)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "internal server error")
	}

	return c.Status(http.StatusOK).JSON(tasks)
}

func (ct *taskController) CreateTask(c *fiber.Ctx) error {
	req := new(models.TaskRequest)
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	task, err := ct.taskService.CreateTask(req)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "internal server error")
	}

	return c.Status(http.StatusOK).JSON(task)
}

func (ct *taskController) DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	err = ct.taskService.DeleteTask(id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "internal server error")
	}

	return c.Status(http.StatusOK).SendString("success")
}
