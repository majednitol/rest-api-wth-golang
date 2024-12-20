package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// AppError is a custom error struct for structured error messages
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface for AppError
func (e *AppError) Error() string {
	return e.Message
}

// GlobalErrorHandler is the middleware for centralized error handling
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors to handle
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if appErr, ok := err.(*AppError); ok {
				// Handle AppError with structured response
				c.JSON(appErr.Code, gin.H{"error": appErr.Message})
			} else {
				// Default to internal server error for other cases
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}
	}
}
