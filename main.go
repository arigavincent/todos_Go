package main

import (
	"fmt"
	"todos/config"
	"todos/database"
	"todos/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	if err := database.ConnectDB(cfg); err != nil {
		panic(err)
	}
	routes.RegisterRoutes(router)

	if err := router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		panic(err)
	}
}
