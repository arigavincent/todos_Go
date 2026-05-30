package repository

import (
	"todos/database"
	"todos/models"

	"gorm.io/gorm"
)

func GetAllUsers() ([]models.UserResponse, error) {
	var users []models.User

	result := database.DB.Preload("Todos").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return toUserResponses(users), nil

}

func GetUserById(id int) (models.UserResponse, error) {
	var user models.User
	result := database.DB.Preload("Todos").First(&user, id)

	if result.Error != nil {
		return models.UserResponse{}, result.Error
	}
	return toUserResponse(user), nil
}
func CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}
func UpdateUser(id int, user *models.User) error {
	result := database.DB.
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func PatchUser(id int, body map[string]any) error {
	result := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(body)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteUser(id int) error {
	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func SoftDeleteUser(id int) error {
	return DeleteUser(id)
}

func HardDeleteUser(id int) error {
	result := database.DB.Unscoped().Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func toUserResponses(users []models.User) []models.UserResponse {
	responses := make([]models.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, toUserResponse(user))
	}
	return responses
}

func toUserResponse(user models.User) models.UserResponse {
	todos := make([]models.TodoSummary, 0, len(user.Todos))
	for _, todo := range user.Todos {
		todos = append(todos, models.TodoSummary{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	return models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Todos: todos,
	}
}

func RestoreUser(id int) error {
	result := database.DB.
		Unscoped().
		Model(&models.User{}).
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
