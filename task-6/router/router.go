package router

import (
	"a2sv-backend-track/task-6/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", taskController.AddTask)
	r.GET("/tasks", taskController.GetTasks)
	r.GET("/tasks/:id", taskController.GetTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}
