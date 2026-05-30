package routes

import (
	"todos/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/todos", handlers.GetAllTodos)
		api.GET("/todos/:id", handlers.GetTodoById)
		api.POST("/todos", handlers.CreateTodo)
		api.PUT("/todos/:id", handlers.UpdateTodo)
		api.PATCH("/todos/:id", handlers.PatchTodo)
		api.DELETE("/todos/:id", handlers.DeleteTodo)
		api.DELETE("/todos_perm/:id", handlers.HardDeleteTodo)
		api.PATCH("/todos_restore/:id", handlers.RestoreTodo)
	}

	user := api.Group("/users")
	{
		user.GET("/", handlers.GetAllUsers)
		user.GET("/:id", handlers.GetUserById)
		user.POST("", handlers.CreateUser)
		user.PUT("/:id", handlers.UpdateUser)
		user.PATCH("/:id", handlers.PatchUser)
		user.DELETE("/:id", handlers.DeleteUser)
		user.DELETE("/perm/:id", handlers.HardDeleteUser)
		user.PATCH("/restore/:id", handlers.RestoreUser)
	}
}
