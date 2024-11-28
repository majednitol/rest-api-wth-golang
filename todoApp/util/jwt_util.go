package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define custom error for expired tokens
var ErrExpiredToken = errors.New("token has expired")

var jwtKey = []byte("your_secret_key")

// GenerateJWT generates a new JWT token for a user
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

// ValidateJWT validates a JWT token
func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		// Check if the error is because of token expiration
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", ErrExpiredToken // Return custom error for expired token
			}
		}
		return "", errors.New("invalid token") // General invalid token error
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["userID"] == nil {
		return "", errors.New("invalid token claims")
	}

	return claims["userID"].(string), nil
}
