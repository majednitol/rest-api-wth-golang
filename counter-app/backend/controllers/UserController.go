package controllers

import (
	"net/http"
	"time"

	"github.com/counter/middleware"
	models "github.com/counter/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	Model *models.UserModel
}

// JWTSecret is the secret key for signing JWT tokens
var JWTSecret = []byte("your_secret_key")

// Signup creates a new user
func (uc *UserController) Signup(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	_, err := uc.Model.CreateUser(body.Username, body.Password)
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}

// Login authenticates a user and returns a JWT token
func (uc *UserController) Login(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
		})
		return
	}

	user, err := uc.Model.FindUserByUsername(body.Username)
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
		})
		return
	}

	if err := uc.Model.ValidatePassword(user.Password, body.Password); err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID.Hex(),
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		c.Error(&middleware.AppError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
