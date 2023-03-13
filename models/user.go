package models

import (
	"time"
)

// User
// This type is used to define user model
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email" validate:"required,email" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"password" validate:"required" gorm:"type:varchar(64);not null"`
	Name      string    `json:"name" validate:"required" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserLogin
// This type is used to define login data model
type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
