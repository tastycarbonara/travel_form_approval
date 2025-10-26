package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID       uint      `gorm:"primaryKey" json:"id"`
	UserName     string    `gorm:"not null" json:"name"`
	UserEmail    string    `gorm:"unique;not null" json:"email"`
	UserPassword string    `gorm:"not null" json:"-"`
	IsActive     bool      `gorm:"not null" json:"is_active"`
	IsDeleted    bool      `gorm:"not null" json:"is_deleted"`
	CreatedDate  time.Time `gorm:"not null" json:"created_date"`
	ModifiedDate time.Time `gorm:"not null" json:"modified_date"`
	CreatedBy    string    `gorm:"not null" json:"created_by"`
	ModifiedBy   string    `gorm:"not null" json:"modified_by"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
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
