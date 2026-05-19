package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	App_Env string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Env file not found")
	}
	return &Config{
		Port:    getEnv("PORT", "3000"),
		App_Env: getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
