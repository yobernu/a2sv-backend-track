package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	// Insert sample trainers
	trainers := []interface{}{
		Trainer{"Alice", 25, "New York"},
		Trainer{"Bob", 30, "London"},
		Trainer{"Charlie", 35, "Paris"},
	}

	_, err = collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted sample trainers!")

	// Query trainers whose name is Alice or Bob
	filter := bson.D{
		{"name", bson.D{
			{"$in", bson.A{"Alice", "Bob"}},
		}},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result Trainer
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found trainer: %+v\n", result)
	}
}
