package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Task struct represents a task in the TODO list
type Task struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Detail string             `bson:"detail" json:"detail"`
}

var taskCollection *mongo.Collection

// Initialize the task collection
func Initialize(client *mongo.Client) {
	taskCollection = client.Database("tododb").Collection("tasks")
}

// CreateTask inserts a new task into the collection
func CreateTask(task Task) (*mongo.InsertOneResult, error) {
	task.ID = primitive.NewObjectID()
	return taskCollection.InsertOne(context.TODO(), task)
}

// GetTasks retrieves all tasks from the collection
func GetTasks() ([]Task, error) {
	var tasks []Task
	cursor, err := taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Task
		if err = cursor.Decode(&task); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTask retrieves a single task by ID
func GetTask(id string) (Task, error) {
	var task Task
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	return task, err
}

// UpdateTask updates an existing task by ID
func UpdateTask(id string, updatedTask Task) (*mongo.UpdateResult, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedTask.Title},
			{Key: "detail", Value: updatedTask.Detail},
		}},
	}
	return taskCollection.UpdateOne(context.TODO(), filter, update)
}

// DeleteTask deletes a task by ID
func DeleteTask(id string) (*mongo.DeleteResult, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	return taskCollection.DeleteOne(context.TODO(), filter)
}
