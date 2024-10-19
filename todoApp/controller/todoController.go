package controller

import (
	"encoding/json"
	"net/http"
	"todoApp/model"

	"github.com/gorilla/mux"
)

// GetTasks returns all tasks
func GetTasks(res http.ResponseWriter, req *http.Request) {
	tasks, err := model.GetTasks()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(tasks)
}

// GetTask returns a single task by ID
func GetTask(res http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := model.GetTask(params["id"])
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(task)
}

// CreateTask creates a new task
func CreateTask(res http.ResponseWriter, r *http.Request) {
	var task model.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	result, err := model.CreateTask(task)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(result)
}

// UpdateTask updates an existing task by ID
func UpdateTask(res http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task model.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	result, err := model.UpdateTask(params["id"], task)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(result)
}

// DeleteTask deletes a task by ID
func DeleteTask(res http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := model.DeleteTask(params["id"])
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(res).Encode(result)
}
