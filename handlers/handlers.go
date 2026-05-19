package handlers

import (
	"net/http"
	"strconv"
	"sync"
	"todos2/models"

	"github.com/gin-gonic/gin"
)

var mu sync.RWMutex
var idNext = 5

var todos = []models.Todo{
	{ID: 1, Title: "Learn Go By myself", Completed: false},
	{ID: 2, Title: "Go to church", Completed: false},
	{ID: 3, Title: "Eat By myself", Completed: false},
	{ID: 4, Title: "Play Soccer By myself", Completed: false},
}

func Get_Todo_By_ID(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An Error Occured",
			"Error":   err.Error(),
		})
		return
	}
	mu.RLock()
	defer mu.RUnlock()
	for _, todo := range todos {
		if todo.ID == idInt {
			c.JSON(http.StatusOK, todo)
			return

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}

func Get_Todos(c *gin.Context) {
	mu.RLock()
	defer mu.RUnlock()
	c.JSON(http.StatusOK, todos)
}

func Create_Todo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	newTodo.ID = idNext
	todos = append(todos, newTodo)
	idNext++
	c.JSON(http.StatusCreated, newTodo)

}

func Update_Todo(c *gin.Context) {
	var updatedTodo models.Todo
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An Error Occured",
			"Error":   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	mu.Lock()
	defer mu.Unlock()
	for i, todo := range todos {
		if todo.ID == idInt {
			updatedTodo.ID = idInt
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
func Delete_Todo(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An Error Occured",
			"Error":   err.Error(),
		})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, todo := range todos {
		if todo.ID == idInt {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func Patch_Todo(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An Error Occured",
			"Error":   err.Error(),
		})
		return
	}
	var body map[string]any
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, todo := range todos {
		if todo.ID == idInt {
			if val, exists := body["title"]; exists {
				if title, ok := val.(string); ok {
					if len(title) < 5 {
						c.JSON(http.StatusBadRequest, gin.H{"error": "title must be at least 5 characters long"})
						return
					}
					if len(title) > 100 {
						c.JSON(http.StatusBadRequest, gin.H{"error": "title must be at most 100 characters long"})
						return
					}
					todos[i].Title = title
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "title must be a string"})
					return
				}
			}
			if val, exists := body["completed"]; exists {
				if completed, ok := val.(bool); ok {
					todos[i].Completed = completed
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "completed must be a boolean"})
					return
				}
			}
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}
