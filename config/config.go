package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var GodotenvLoad = godotenv.Load

type PostgresDbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func LoadConfig() (*PostgresDbConfig, error) {
	if err := GodotenvLoad(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" || user == "" || password == "" || dbName == "" || port == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid port number")
	}

	return &PostgresDbConfig{
		Host:     host,
		Port:     portInt,
		User:     user,
		Password: password,
		DbName:   dbName,
	}, nil
}
