package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Role struct {
// 	role_id   uint   `gorm:"primaryKey"`
// 	role_name string `gorm:"unique;not null"`
// }

// type User struct {
// 	user_id       uint   `gorm:"primaryKey"`
// 	user_name     string `gorm:"not null"`
// 	user_email    string `gorm:"unique;not null"`
// 	user_password string `gorm:"not null"`
// }

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var psqlInfo string = CreateConnString()

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	} else {
		_ = db
		fmt.Println("connected successfully!")
	}

	// err = db.AutoMigrate(&Role{}, &User{})
	// if err != nil {
	// 	log.Fatal("failed to migrate:", err)
	// }
}

func CreateConnString() string {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	return psqlInfo
}
