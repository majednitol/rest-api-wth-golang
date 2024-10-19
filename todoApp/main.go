package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"todoApp/model"
	"todoApp/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection setup
	mongoURI := "mongodb://mongo:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Initialize the task collection
	model.Initialize(client)

	// Setup router
	r := router.SetupRouter()

	// Print message to indicate server is running
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
