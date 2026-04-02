package controllers

import (
	model "a2sv-backend-track/task-5/models"
	"a2sv-backend-track/task-5/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	services services.TaskManager
}

func NewTaskController(s services.TaskManager) *TaskController {
	return &TaskController{services: s}
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := tc.services.AddTask(task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, task)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks := tc.services.ListTasks()
	c.JSON(200, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task, err := tc.services.GetTask(taskID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := tc.services.UpdateTask(taskID, updatedTask); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if err := tc.services.DeleteTask(taskID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
