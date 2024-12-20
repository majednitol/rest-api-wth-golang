package router

import (
	"github.com/counter/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the API
func InitializeRoutes(r *gin.Engine, userController *controllers.UserController) {
    r.POST("/signup", userController.Signup)
    r.POST("/login", userController.Login)
}

