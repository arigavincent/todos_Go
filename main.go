package main

import (
	"log"
	"net/http"

	"todos2/config"
	"todos2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello My Name Is Mr.Ariga A Passionate Go Developer",
			"status":  "success",
		})
	})

	routes.RegisterRoutes(router)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
