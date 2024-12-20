package router

import (
	"github.com/counter/controllers"
	"github.com/counter/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the API
func SetupRoutes(r *gin.Engine, counterController *controllers.CounterController) {
	// Public routes (e.g., for login, signup)
 // You need to implement the Signup handler

	// Secure routes
	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware()) // Apply JWT middleware to secure these routes
	{
		api.GET("/counters", counterController.GetCounters)
		api.POST("/counter", counterController.CreateCounter)
		api.PUT("/counter/:id", counterController.UpdateCounter)
		api.DELETE("/counter/:id", counterController.DeleteCounter)
	}
}
