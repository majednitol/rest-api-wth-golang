package router

import (
	"fmt"
	"net/http"
	"todoApp/controller"
	"todoApp/middleware"

	"github.com/gorilla/mux"
)

// CORS middleware to set headers for CORS
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request for debugging purposes
		fmt.Println("CORS Middleware: Method =", r.Method)

		// Allow specific origin
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request by returning status OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Continue to the next handler for non-OPTIONS requests
		next.ServeHTTP(w, r)
	})
}

// SetupRouter initializes the router
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Apply CORS middleware to all routes
	router.Use(CORS)

	// Public routes
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/register", controller.Register).Methods("POST")

	// Manually handle OPTIONS for login and register routes
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware) // Apply authentication middleware

	// Tasks routes
	api.HandleFunc("/tasks", controller.GetTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", controller.GetTask).Methods("GET")
	api.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	api.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")

	// Manually handle OPTIONS for tasks routes
	api.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	api.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	return router
}