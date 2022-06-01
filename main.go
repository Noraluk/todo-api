package main

import (
	"flag"
	"log"
	"todo-api/internal/configs"
	"todo-api/internal/routes"
	"todo-api/pkg/database"
)

func main() {
	configPath := flag.String("config_path", "configs", "set configs path")

	if err := configs.Conf.New(*configPath); err != nil {
		log.Fatal(err)
	}

	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	configs.SetDomain(db)

	app := routes.Init()
	app.Listen(":8080")
}
