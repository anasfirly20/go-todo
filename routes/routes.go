package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

type TaskRouteController struct {
	taskController controllers.TaskController
}

func NewTaskRouteController(taskController controllers.TaskController) TaskRouteController {
	return TaskRouteController{taskController}
}

func (rc *TaskRouteController) TaskRoute(rg *gin.RouterGroup) {
	router := rg.Group("/task")

	router.GET("/", rc.taskController.GetTasks)
	router.GET("/:id", rc.taskController.GetTaskById)
	router.POST("/", rc.taskController.CreateTask)
	router.PATCH("/:id", rc.taskController.UpdateTask)
	router.DELETE("/:id", rc.taskController.DeleteTask)
}