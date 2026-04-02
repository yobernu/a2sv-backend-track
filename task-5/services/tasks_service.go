package services

import (
	model "a2sv-backend-track/task-5/models"
	"errors"

	"github.com/google/uuid"
)

type TaskManager interface {
	AddTask(task model.Task) error
	UpdateTask(taskID string, updatedTask model.Task) error
	DeleteTask(taskID string) error
	GetTask(taskID string) (model.Task, error)
	ListTasks() []model.Task
}

type TaskService struct {
	tasks []model.Task
}

func NewTaskService() *TaskService {
	return &TaskService{tasks: []model.Task{}}
}

func (s *TaskService) AddTask(task model.Task) error {
	task.ID = uuid.New().String()
	s.tasks = append(s.tasks, task)
	return nil
}

func (s *TaskService) UpdateTask(taskID string, updatedTask model.Task) error {
	for i, t := range s.tasks {
		if t.ID == taskID {
			// Only update fields that were provided (non-nil)
			if updatedTask.Title != nil {
				t.Title = updatedTask.Title
			}
			if updatedTask.Description != nil {
				t.Description = updatedTask.Description
			}
			if updatedTask.Status != nil {
				t.Status = updatedTask.Status
			}

			// Save the merged task back into the slice
			s.tasks[i] = t
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) DeleteTask(taskID string) error {
	for i, t := range s.tasks {
		if t.ID == taskID {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) GetTask(taskID string) (model.Task, error) {
	for _, t := range s.tasks {
		if t.ID == taskID {
			return t, nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func (s *TaskService) ListTasks() []model.Task {
	return s.tasks
}
