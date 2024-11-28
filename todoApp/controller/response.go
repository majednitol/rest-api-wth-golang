package controller

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error message returned to the client
type ErrorResponse struct {
	Message string `json:"message"`
}

// SendErrorResponse sends a JSON-formatted error response
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
