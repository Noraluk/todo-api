package configs

import (
	"todo-api/internal/controllers"
	"todo-api/internal/domains"
	"todo-api/internal/repositories"
	"todo-api/internal/services"

	"gorm.io/gorm"
)

type (
	controller struct {
		Task domains.TaskControler
	}

	service struct {
		Task domains.TaskService
	}

	repository struct {
		Task domains.TaskRepository
	}
)

var Controller controller
var Service service
var Repository repository

func SetDomain(db *gorm.DB) {
	setRepository(db)
	setService()
	setController()
}

func setController() {
	Controller = controller{
		Task: controllers.NewTaskController(Service.Task),
	}
}

func setService() {
	Service = service{
		Task: services.NewTaskService(Repository.Task),
	}
}

func setRepository(db *gorm.DB) {
	Repository = repository{
		Task: repositories.NewTaskRepository(db),
	}
}
