package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tastycarbonara/travel_form_approval/db"
	"github.com/tastycarbonara/travel_form_approval/models"
	"github.com/tastycarbonara/travel_form_approval/routes"
)

// type Role struct {
// 	role_id   uint   `gorm:"primaryKey"`
// 	role_name string `gorm:"unique;not null"`
// }

func main() {

	godotenv.Load()
	db.Connect()

	db.DB.AutoMigrate(&models.User{})

	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
