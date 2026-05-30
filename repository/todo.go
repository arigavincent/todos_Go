package repository

import (
	"todos/database"
	"todos/models"

	"gorm.io/gorm"
)

func GetAllTodos() ([]models.TodoResponse, error) {
	var todos []models.TodoResponse

	result := database.DB.Preload("User").Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil

}

func GetTodoById(id int) (models.TodoResponse, error) {
	var todo models.TodoResponse
	result := database.DB.Preload("User").First(&todo, id)

	if result.Error != nil {
		return models.TodoResponse{}, result.Error
	}
	return todo, nil
}
func CreateTodo(todo *models.CreateTodoRequest) error {
	result := database.DB.Create(todo)
	return result.Error
}
func UpdateTodo(id int, todo *models.Todo) error {
	result := database.DB.
		Model(&models.Todo{}).
		Where("id = ?", id).
		Updates(todo)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func PatchTodo(id int, body map[string]any) error {
	result := database.DB.
		Model(&models.Todo{}).
		Where("id = ?", id).
		Updates(body)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func SoftDeleteTodo(id int) error {
	result := database.DB.
		Where("id = ?", id).
		Delete(&models.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func HardDeleteTodo(id int) error {
	result := database.DB.
		Unscoped().
		Where("id = ?", id).
		Delete(&models.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func RestoreTodo(id int) error {
	result := database.DB.
		Unscoped().
		Model(&models.Todo{}).
		Where("id = ?", id).
		Update("deleted_at", nil)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
