package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var GodotenvLoad = viper.AutomaticEnv

type PostgresDbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func LoadConfig() (*PostgresDbConfig, error) {
	if _, err := os.Stat(".env"); err == nil {
		// If the file exists, load the .env file
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading .env file: %w", err)
		}
	}
	viper.AutomaticEnv()
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

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
