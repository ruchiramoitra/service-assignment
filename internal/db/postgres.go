package db

import (
	"database/sql"
	"fmt"
	"kong-assignment/config"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

func NewPostgres(pgConfig *config.PostgresDbConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", pgConfig.Host, pgConfig.User, pgConfig.Password, pgConfig.DbName, pgConfig.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	log.Println("Connected to database")
	return db, nil
}
