package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

// GenerateJWT generates a new JWT token for a user
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

// ValidateJWT validates a JWT token and identifies errors
func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return "", errors.New("token expired")
			case ve.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				return "", errors.New("invalid token signature")
			}
		}
		return "", fmt.Errorf("token parse error: %v", err)
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	// Validate exp claim
	exp, ok := claims["exp"].(float64)
	if !ok || int64(exp) < time.Now().Unix() {
		return "", errors.New("token has expired or is invalid")
	}

	// Extract and validate userID
	userID, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return userID, nil
}
