package controllers

import (
	services "a2sv-backend-track/task-6/data"
	model "a2sv-backend-track/task-6/models"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	services services.TaskManager
}

func NewTaskController(s services.TaskManager) *TaskController {
	return &TaskController{services: s}
}

func (tc *TaskController) AddTask(c *gin.Context) {
	var raw json.RawMessage
	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Try single task
	var task model.Task
	if err := json.Unmarshal(raw, &task); err == nil && task.Title != nil {
		if err := tc.services.AddTask(task); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, task)
		return
	}

	// Try multiple tasks
	var tasks []model.Task
	if err := json.Unmarshal(raw, &tasks); err == nil {
		if err := tc.services.AddTask(tasks...); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, tasks)
		return
	}

	c.JSON(400, gin.H{"error": "Invalid JSON format"})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.services.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
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
