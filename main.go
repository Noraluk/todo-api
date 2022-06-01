package main

import (
	"log"
	"todo-api/internal/configs"
	"todo-api/internal/routes"
	"todo-api/pkg/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	configs.SetDomain(db)

	app := routes.Init()
	app.Listen(":8080")
}
