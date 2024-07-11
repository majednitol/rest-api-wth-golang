package main

import (
	"fmt"
	"log"
	"net/http"
	"todoApp/model"
	"todoApp/router"
)

func main() {
	r := router.SetupRouter()

	// Initialize sample data
	model.Tasks = append(model.Tasks, model.Task{ID: "1", Title: "Task 1", Detail: "Details of Task 1"})
	model.Tasks = append(model.Tasks, model.Task{ID: "2", Title: "Task 2", Detail: "Details of Task 2"})
	model.Tasks = append(model.Tasks, model.Task{ID: "3", Title: "Task 3", Detail: "Details of Task 3"})

	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
