package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"todoApp/middleware"
	"todoApp/model"
	"todoApp/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection URI
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://mongo:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Initialize model
	model.Initialize(client)

	// Set up router
	r := router.SetupRouter()
	handler := middleware.LoggingMiddleware(middleware.RecoverMiddleware(r))

	// Start the server
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
