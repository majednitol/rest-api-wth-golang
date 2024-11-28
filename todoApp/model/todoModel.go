package model

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"-"`
}

// Task represents a task in the system.
type Task struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title  string             `bson:"title" json:"title"`
	Detail string             `bson:"detail" json:"detail"`
}

var (
	userCollection *mongo.Collection
	taskCollection *mongo.Collection
)

// Initialize sets up the MongoDB collections for users and tasks.
func Initialize(client *mongo.Client) {
	userCollection = client.Database("tododb").Collection("users")
	taskCollection = client.Database("tododb").Collection("tasks")
}

// User-related functions

// RegisterUser adds a new user with a hashed password.
func RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = userCollection.InsertOne(context.TODO(), User{
		Username: username,
		Password: string(hashedPassword),
	})
	return err
}

// AuthenticateUser verifies the user's password.
func AuthenticateUser(username, password string) (User, error) {
	var user User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, errors.New("invalid password")
	}
	return user, nil
}

// Task-related functions

// CreateTask inserts a new task into the collection.
func CreateTask(task Task) (*mongo.InsertOneResult, error) {
	if err := validateTask(task); err != nil {
		return nil, err
	}

	task.ID = primitive.NewObjectID()
	return taskCollection.InsertOne(context.TODO(), task)
}

// GetTasks retrieves all tasks from the collection.
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

// GetTask retrieves a task by its ID.
func GetTask(id string) (Task, error) {
	objID, err := validateObjectID(id)
	if err != nil {
		return Task{}, err
	}

	var task Task
	err = taskCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&task)
	return task, err
}

// UpdateTask updates an existing task by ID.
func UpdateTask(id string, updatedTask Task) (*mongo.UpdateResult, error) {
	objID, err := validateObjectID(id)
	if err != nil {
		return nil, err
	}
	if err := validateTask(updatedTask); err != nil {
		return nil, err
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: updatedTask.Title},
		{Key: "detail", Value: updatedTask.Detail},
	}}}
	return taskCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}, update)
}

// DeleteTask deletes a task by its ID.
func DeleteTask(id string) (*mongo.DeleteResult, error) {
	objID, err := validateObjectID(id)
	if err != nil {
		return nil, err
	}

	return taskCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objID}})
}

// Validation functions

// validateTask ensures that the task fields are not empty.
func validateTask(task Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}
	if task.Detail == "" {
		return errors.New("detail cannot be empty")
	}
	return nil
}

// validateObjectID verifies if the given string is a valid MongoDB ObjectID.
func validateObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, errors.New("invalid ID format")
	}
	return objID, nil
}
