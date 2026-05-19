package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" binding:"required,min=5,max=10"`
	Completed bool   `json:"completed"`
}
