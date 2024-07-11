package controller

import (
	"encoding/json"
	"net/http"
	"todoApp/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetTasks returns all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Tasks)
}

// GetTask returns a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range model.Tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Task{})
}

// CreateTask creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = uuid.New().String() // Generate UUID for new task
	model.Tasks = append(model.Tasks, task)
	json.NewEncoder(w).Encode(model.Tasks)
}

// UpdateTask updates an existing task by ID
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range model.Tasks {
		if item.ID == params["id"] {
			model.Tasks = append(model.Tasks[:index], model.Tasks[index+1:]...)
			var task model.Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			model.Tasks = append(model.Tasks, task)
			json.NewEncoder(w).Encode(model.Tasks)
			return
		}
	}
}

// DeleteTask deletes a task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range model.Tasks {
		if item.ID == params["id"] {
			model.Tasks = append(model.Tasks[:index], model.Tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(model.Tasks)
}
