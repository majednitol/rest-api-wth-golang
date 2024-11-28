package router

import (
	"todoApp/controller"
	"todoApp/middleware"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/register", controller.Register).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/tasks", controller.GetTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", controller.GetTask).Methods("GET")
	api.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")

	return router
}
