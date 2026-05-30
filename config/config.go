package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	AppEnv     string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		Port:       getEnv("PORT", "3000"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "ariga"),
		DBPassword: getEnv("DB_PASSWORD", "1308Ariga"),
		DBName:     getEnv("DB_NAME", "gorm"),
		DBPort:     getEnv("DB_PORT", "5432"),
		AppEnv:     getEnv("APP_ENV", "development"),
	}, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
