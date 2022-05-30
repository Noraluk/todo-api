package main

import (
	"todo-api/routes"
)

func main() {
	app := routes.Init()
	app.Listen(":8080")
}
