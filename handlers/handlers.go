package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"todos/models"
	"todos/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTodos(c *gin.Context) {
	result, err := repository.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid ID, ID must be numerical",
		})
		return
	}
	result, err := repository.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.CreateTodoRequest
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify user exists before creating todo
	_, err := repository.GetUserById(int(newTodo.UserID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := repository.CreateTodo(&newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.TodoResponse{
		Title:     newTodo.Title,
		Completed: newTodo.Completed,
		UserID:    newTodo.UserID,
	})
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID format"})
		return
	}
	var updatedTodo models.Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	updatedTodo.ID = uint(id)
	errr := repository.UpdateTodo(id, &updatedTodo)
	if errr != nil {
		if errors.Is(errr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Todo Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": errr.Error(),
		})

	}
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID format"})
		return
	}
	errr := repository.SoftDeleteTodo(id)
	if errr != nil {
		if errors.Is(errr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Todo Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": errr.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with ID %d has been deleted ", id),
	})
}

func PatchTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID format"})
		return
	}

	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if val, exists := body["title"]; exists {
		title, ok := val.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Title must be a string"})
			return
		}

		if len(title) < 3 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Title Must Have Atleast 3 letters",
			})
			return
		}

		if len(title) > 100 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Title Can't be longer than 100 letters",
			})
			return
		}
	}

	if val, exists := body["completed"]; exists {
		_, ok := val.(bool)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Completed must be a boolean value (true or false)"})
			return
		}
	}
	errr := repository.PatchTodo(id, body)
	if errr != nil {
		if errors.Is(errr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Todo Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": errr.Error(),
		})

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo patched successfully",
	})
}

func HardDeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID format"})
		return
	}
	errr := repository.HardDeleteTodo(id)
	if errr != nil {
		if errors.Is(errr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Todo Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": errr.Error(),
		})

	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with ID %d has been deleted Permanently", id),
	})
}

func RestoreTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID format"})
		return
	}

	err = repository.RestoreTodo(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Todo Not Found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Todo with ID %d restored", id),
	})
}
