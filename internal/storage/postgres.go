package storage

import (
	"database/sql"
	"fmt"
	"kong-assignment/config"
	"kong-assignment/internal/models"
	"log"
	"strings"

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

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{DB: db}
}

func (ps *PostgresStorage) GetServices(queryParams models.QueryParams) ([]models.Service, error) {
	query := "SELECT * FROM services"
	var conditions []string

	if queryParams.Name != "" {
		conditions = append(conditions, "name = "+queryParams.Name)
	}
	if queryParams.Sort != "" {
		conditions = append(conditions, " ORDER BY "+queryParams.Sort)
	}
	if queryParams.Limit != "" {
		conditions = append(conditions, "LIMIT "+queryParams.Limit)
	}
	if queryParams.Offset != "" {
		conditions = append(conditions, "OFFSET "+queryParams.Offset)
	}

	// Add conditions to the query if any
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions[:1], " ")
		query += strings.Join(conditions[1:], " ")
	}
	fmt.Println("actual query: ", query)
	rows, err := ps.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.Id, &service.Name, &service.Description, &service.Version)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return services, nil
}
