package controllers

import (
	"net/http"
	"strconv"
	"todo-api/internal/domains"
	"todo-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

type taskController struct {
	domains.TaskService
}

func NewTaskController(taskService domains.TaskService) domains.TaskControler {
	return &taskController{taskService}
}

func (ct *taskController) GetTasks(c *fiber.Ctx) error {
	tasks, err := ct.TaskService.GetTasks()
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

	task, err := ct.TaskService.CreateTask(req)
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

	err = ct.TaskService.DeleteTask(id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "internal server error")
	}

	return c.Status(http.StatusOK).SendString("success")
}
