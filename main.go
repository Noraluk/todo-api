package main

import (
	"todo-api/internal/routes"
)

func main() {
	// if err := configs.Conf.New(*configPath); err != nil {
	// 	log.Fatal(err)
	// }

	// database.Init()

	// configs.SetDomain(db)

	app := routes.Init()
	app.Listen(":8080")
}
