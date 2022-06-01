package routes

import "github.com/gofiber/fiber/v2"

func Init() *fiber.App {
	app := fiber.New()

	newTaskRoutes(app)

	return app
}
