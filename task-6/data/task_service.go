package services

import (
	model "a2sv-backend-track/task-6/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManager interface {
	AddTask(tasks ...model.Task) error
	UpdateTask(taskID string, updatedTask model.Task) error
	DeleteTask(taskID string) error
	GetTask(taskID string) (model.Task, error)
	ListTasks() ([]model.Task, error)
}

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{collection: collection}
}

func (s *TaskService) AddTask(tasks ...model.Task) error {
	var docs []interface{}
	for i := range tasks {
		tasks[i].ID = primitive.NewObjectID().Hex()
		docs = append(docs, tasks[i])
	}

	if len(docs) == 1 {
		_, err := s.collection.InsertOne(context.TODO(), docs[0])
		return err
	}
	_, err := s.collection.InsertMany(context.TODO(), docs)
	return err
}

func (s *TaskService) UpdateTask(taskID string, updatedTask model.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{}
	if updatedTask.Title != nil {
		update["title"] = updatedTask.Title
	}
	if updatedTask.Description != nil {
		update["description"] = updatedTask.Description
	}
	if updatedTask.Status != nil {
		update["status"] = updatedTask.Status
	}

	res, err := s.collection.UpdateOne(ctx,
		bson.M{"id": taskID},
		bson.M{"$set": update},
	)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *TaskService) DeleteTask(taskID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.collection.DeleteOne(ctx, bson.M{"id": taskID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *TaskService) GetTask(taskID string) (model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task model.Task
	err := s.collection.FindOne(ctx, bson.M{"id": taskID}).Decode(&task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskService) ListTasks() ([]model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []model.Task
	for cursor.Next(ctx) {
		var task model.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
