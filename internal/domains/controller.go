package domains

import "todo-api/internal/controllers"

type controller struct {
	Task controllers.TaskControler
}

var Controller controller = controller{
	Task: controllers.NewTaskController(Service.Task),
}
