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

// query to get all services
const BASE_QUERY = "SELECT s.id AS service_id, s.name AS service_name,  s.description AS service_description, COALESCE(COUNT(v.id), 0) AS total_versions,  COALESCE(json_agg(v.name) FILTER (WHERE v.id IS NOT NULL), '[]') AS versions FROM services s LEFT JOIN versions v ON s.id = v.service_id"

func (ps *PostgresStorage) GetServices(queryParams models.QueryParams) ([]models.Service, error) {
	var conditions []string

	var query = BASE_QUERY

	if queryParams.Sort != "" {
		conditions = append(conditions, "ORDER BY "+queryParams.Sort)
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
	query += " GROUP BY s.id"
	rows, err := ps.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service_id, service_name, service_description string
		var total_versions int
		var versions []uint8
		err := rows.Scan(&service_id, &service_name, &service_description, &total_versions, &versions)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		service := models.Service{
			Id:            service_id,
			Name:          service_name,
			Description:   service_description,
			TotalVersions: total_versions,
			Versions:      []string{string(versions)},
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return services, nil
}

func (ps *PostgresStorage) SearchService(queryParams models.QueryParams) ([]models.Service, error) {
	var query = BASE_QUERY + " WHERE "

	// can be queried by name or id
	if queryParams.Name != "" {
		query += fmt.Sprintf("s.name = '%s'", strings.Trim(queryParams.Name, `"'`))
	} else {
		query += fmt.Sprintf("s.id = %s", strings.Trim(queryParams.Id, `"'`))
	}

	// Add conditions to the query
	query += " GROUP BY s.id"
	rows, err := ps.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service_id, service_name, service_description string
		var total_versions int
		var versions []uint8
		err := rows.Scan(&service_id, &service_name, &service_description, &total_versions, &versions)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		service := models.Service{
			Id:            service_id,
			Name:          service_name,
			Description:   service_description,
			TotalVersions: total_versions,
			Versions:      []string{string(versions)},
		}

		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return services, nil
}
