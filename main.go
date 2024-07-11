package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Task struct represents a task in the TODO list
type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

var tasks []Task

func main() {
	router := mux.NewRouter()

	// Initialize sample data
	tasks = append(tasks, Task{ID: "1", Title: "Task 1", Detail: "Details of Task 1"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Detail: "Details of Task 2"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Detail: "Details of Task 2"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Detail: "Details of Task 2"})

	// Route handles & endpoints
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
fmt.Println("server running on 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// GetTasks returns all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// GetTask returns a single task by ID
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

// CreateTask creates a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = uuid.New().String() // Generate UUID for new task
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask updates an existing task by ID
func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			json.NewEncoder(w).Encode(tasks)
			return
		}
	}
}

// DeleteTask deletes a task by ID
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}
