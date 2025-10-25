package routes

import (
	"github.com/gorilla/mux"
	"github.com/tastycarbonara/travel_form_approval/handlers"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
}
