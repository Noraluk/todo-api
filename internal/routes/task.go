package routes

import (
	"todo-api/internal/domains"

	"github.com/gofiber/fiber/v2"
)

func newTaskRoutes(app *fiber.App) {
	ctrl := domains.Controller.Task

	taskGroup := app.Group("/tasks")
	taskGroup.Get("", ctrl.GetTasks)
	taskGroup.Post("", ctrl.CreateTask)
	taskGroup.Delete("/:id", ctrl.DeleteTask)
}
