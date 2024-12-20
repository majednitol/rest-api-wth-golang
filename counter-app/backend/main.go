package main

import (
	"github.com/counter/controllers"
	"github.com/counter/middleware"
	models "github.com/counter/model"
	"github.com/counter/router"
	routes "github.com/counter/router"
	"github.com/counter/services"
	utils "github.com/counter/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	client := utils.ConnectDatabase("mongodb://localhost:27017")
	defer client.Disconnect(nil)

	counterModel := &models.CounterModel{Collection: client.Database("counterdb").Collection("counters")}
	counterController := &controllers.CounterController{Model: counterModel}

	r := gin.Default()
	userCollection := client.Database("counterdb").Collection("users")
	userModel := &models.UserModel{Collection: userCollection}
	userController := &controllers.UserController{Model: userModel}

    // Setup Routes
    router.InitializeRoutes(r, userController)

	// CORS configuration for React app
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow your React app's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Global error handler middleware
	r.Use(middleware.GlobalErrorHandler())

	// Setup routes
	routes.SetupRoutes(r, counterController)

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		services.HandleWebSocket(c.Writer, c.Request)
	})

	r.Run(":8080")
}
