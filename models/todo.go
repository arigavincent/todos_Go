package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo Table in DB
type Todo struct {
	gorm.Model
	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"default:false"`
	UserID    uint   `json:"user_id" gorm:"not null"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
}

// Request body for creating a new todo
type CreateTodoRequest struct {
	Title     string `json:"title" binding:"required,min=3,max=100"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id" binding:"required"`
}

// Clipped Todo data for listing todos without user details in user response
type TodoSummary struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// full Todo data for detailed responses, including user details
type TodoResponse struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	Completed bool        `json:"completed"`
	UserID    uint        `json:"user_id"`
	User      UserSummary `json:"user"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
