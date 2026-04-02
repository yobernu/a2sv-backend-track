package main

import (
	"a2sv-backend-track/task-5/controllers"
	"a2sv-backend-track/task-5/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	service := &services.TaskService{}
	controller := controllers.NewTaskController(service)
	router.POST("/task", controller.AddTask)
	router.GET("/tasks", controller.GetTasks)
	router.GET("/task/:id", controller.GetTask)
	router.PUT("/task/:id", controller.UpdateTask)
	router.DELETE("/task/:id", controller.DeleteTask)
	router.Run(":8080")
}
