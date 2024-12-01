package storage

import "kong-assignment/internal/models"

type ServiceRepo interface {
	GetServices(queryParams models.QueryParams) ([]models.Service, string, error)
	SearchService(queryParams models.QueryParams) ([]models.Service, error)
}
