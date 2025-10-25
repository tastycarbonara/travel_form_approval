package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tastycarbonara/travel_form_approval/db"
	"github.com/tastycarbonara/travel_form_approval/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var receievedUser models.CreateUserRequest
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&receievedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.UserName = receievedUser.Name
	user.UserEmail = receievedUser.Email
	user.UserPassword = receievedUser.Password

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})
}
