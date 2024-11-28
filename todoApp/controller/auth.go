package controller

import (
	"encoding/json"
	"net/http"
	"todoApp/model"
	"todoApp/util"
)

// Login handles user authentication
func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := model.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		SendErrorResponse(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	token, err := util.GenerateJWT(user.ID.Hex())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Register handles new user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := model.RegisterUser(creds.Username, creds.Password)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
