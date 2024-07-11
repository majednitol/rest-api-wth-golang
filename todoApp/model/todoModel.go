package model

// Task struct represents a task in the TODO list
type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// Define the tasks slice
var Tasks []Task
