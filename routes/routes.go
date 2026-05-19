package routes

import (
	"todos2/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/todos", handlers.Get_Todos)
	router.GET("/todos/:id", handlers.Get_Todo_By_ID)
	router.POST("/todos", handlers.Create_Todo)
	router.PUT("/todos/:id", handlers.Update_Todo)
	router.DELETE("/todos/:id", handlers.Delete_Todo)

}
