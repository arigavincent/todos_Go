package models

import "gorm.io/gorm"

// User Table in DB
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Todos    []Todo `json:"todos" gorm:"foreignKey:UserID"`
}

// Request body for creating a new user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Clipped User data for listing users without todo details in todo response
type UserSummary struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Full User data for detailed responses, including todo details
type UserResponse struct {
	ID    uint          `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Todos []TodoSummary `json:"todos"`
}
