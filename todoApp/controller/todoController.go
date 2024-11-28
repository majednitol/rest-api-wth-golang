package controller

import (
	"encoding/json"
	"net/http"
	"todoApp/model"

	"github.com/gorilla/mux"
)

// GetTasks returns all tasks
func GetTasks(w http.ResponseWriter, req *http.Request) {
	tasks, err := model.GetTasks()
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if tasks == nil {
		SendErrorResponse(w, http.StatusNotFound, "No tasks found")
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// GetTask returns a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := model.GetTask(params["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	json.NewEncoder(w).Encode(task)
}

// CreateTask creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	result, err := model.CreateTask(task)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(result)
}

// UpdateTask updates an existing task by ID
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	result, err := model.UpdateTask(params["id"], task)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(result)
}

// DeleteTask deletes a task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := model.DeleteTask(params["id"])
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(result)
}
