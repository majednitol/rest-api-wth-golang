package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"todoApp/util"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		fmt.Println("Raw Authorization Header:", authHeader)

		// Normalize the Authorization header by removing duplicate "Bearer"
		parts := strings.Fields(authHeader) // Split by spaces
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1] // Extract the token
		fmt.Println("Extracted Token:", token)

		userID, err := util.ValidateJWT(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Store userID in context
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
