package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret is the secret key for signing JWT tokens
var JWTSecret = []byte("your_secret_key")

// JWTMiddleware validates the JWT token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(&AppError{
				Code:    http.StatusUnauthorized,
				Message: "Authorization header is missing",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is as expected
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return JWTSecret, nil
		})

		if err != nil {
			var errorMessage string
			switch err {
			case jwt.ErrSignatureInvalid:
				errorMessage = "Invalid token signature"
			case jwt.ErrTokenMalformed:
				errorMessage = "Malformed token"
			case jwt.ErrTokenExpired:
				errorMessage = "Token has expired"
			case jwt.ErrTokenNotValidYet:
				errorMessage = "Token is not yet valid"
			default:
				errorMessage = "Failed to parse token"
			}
			c.Error(&AppError{
				Code:    http.StatusUnauthorized,
				Message: errorMessage,
			})
			return
		}

		if !token.Valid {
			c.Error(&AppError{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["userID"] == nil {
			c.Error(&AppError{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token claims",
			})
			return
		}
		c.Set("userID", claims["userID"])
		c.Next()
	}
}
