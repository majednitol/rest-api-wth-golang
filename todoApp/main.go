package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"todoApp/model"
	"todoApp/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI") // Use environment variable for MongoDB URI
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
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

	model.Initialize(client)

	r := router.SetupRouter()

	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
