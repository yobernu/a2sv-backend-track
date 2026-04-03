package main

import (
	"a2sv-backend-track/task-6/controllers"
	services "a2sv-backend-track/task-6/data"
	"a2sv-backend-track/task-6/router"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("test")
	taskService := services.NewTaskService(db.Collection("tasks"))
	taskControllers := controllers.NewTaskController(taskService)
	router := router.SetupRouter(taskControllers)
	router.Run(":8080")
}
