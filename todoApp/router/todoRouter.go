package router

import (
	"todoApp/controller"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router and sets up the routes
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Define routes and associate with controller functions
	router.HandleFunc("/tasks", controller.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", controller.GetTask).Methods("GET")
	router.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")

	return router
}
