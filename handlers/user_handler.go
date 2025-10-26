package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/tastycarbonara/travel_form_approval/db"
	"github.com/tastycarbonara/travel_form_approval/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	isTokenValid, _, err := checkToken(r)
	if !isTokenValid && err != nil {
		http.Error(w, "Token is invalid", http.StatusUnauthorized)
		return
	}

	var users []models.User
	db.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	isTokenValid, loggedUser, err := checkToken(r)
	if !isTokenValid && err != nil {
		http.Error(w, "Token is invalid", http.StatusUnauthorized)
		return
	}

	var receievedUser models.CreateUserRequest
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&receievedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existing_user = db.DB.Where("user_email = ?", receievedUser.Email).First(&user)
	if existing_user.Error == nil {
		http.Error(w, "Email already exist", http.StatusBadRequest)
		return
	}

	user.UserName = receievedUser.Name
	user.UserEmail = receievedUser.Email
	user.UserPassword = receievedUser.Password
	user.IsActive = true
	user.IsDeleted = false
	user.CreatedDate = time.Now()
	user.ModifiedDate = time.Now()
	user.CreatedBy = loggedUser
	user.ModifiedBy = loggedUser

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

func Login(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userResult := db.DB.Where("user_email = ?", request.Email).
		Where("is_active = ?", true).Where("is_deleted = ? ", false).First(&user)
	if userResult.Error != nil {
		if errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusUnauthorized)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(request.Password))
	if err != nil {
		http.Error(w, "Wrong username or password", http.StatusBadRequest)
		return
	}

	token, err := createToken(user.UserEmail)
	if err != nil {
		http.Error(w, "Can't generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login success",
		"user":    user,
		"token":   token,
	})
}
