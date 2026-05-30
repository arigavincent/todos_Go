package database

import (
	"fmt"
	"todos/config"
	"todos/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	fmt.Println("DB Connected")
	if err := db.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		return err
	}
	fmt.Println("Migration succeeded Successfully")
	DB = db
	return nil
}
