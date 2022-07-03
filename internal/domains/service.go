package domains

import "todo-api/internal/services"

type service struct {
	Task services.TaskService
}

var Service service = service{
	Task: services.NewTaskService(),
}
