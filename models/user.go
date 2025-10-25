package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID       uint   `gorm:"primaryKey" json:"id"`
	UserName     string `gorm:"not null" json:"name"`
	UserEmail    string `gorm:"unique;not null" json:"email"`
	UserPassword string `gorm:"not null" json:"-"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.UserPassword) < 8 {
		return errors.New("password too short (min 8 characters)")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.UserPassword = string(hashed)
	return nil
}
