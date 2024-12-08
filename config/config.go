package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database Database
	Server   Server
}

type Database struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

type Server struct {
	APP_PORT string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	config := &Config{
		Server: Server{
			APP_PORT: os.Getenv("APP_PORT"),
		},
		Database: Database{
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_PORT:     os.Getenv("DB_PORT"),
			DB_USERNAME: os.Getenv("DB_USERNAME"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_DATABASE: os.Getenv("DB_DATABASE"),
		},
	}

	return config, nil
}
