package storage_test

import (
	"kong-assignment/internal/models"
	"kong-assignment/internal/storage"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetServices(t *testing.T) {
	// Create a mock database and service storage instance
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	ps := storage.NewPostgresStorage(db)

	tests := []struct {
		name        string
		queryParams models.QueryParams
		mockQuery   string
		mockRows    *sqlmock.Rows
		wantErr     bool
		expected    []models.Service
	}{
		{
			name: "valid query params",
			queryParams: models.QueryParams{
				Name:   "test-service",
				Sort:   "name",
				Limit:  "10",
				Offset: "0",
			},
			// Adjust query to match how the actual code will build it
			mockQuery: `SELECT \* FROM services WHERE name = test-service ORDER BY name LIMIT 10 OFFSET 0`,
			mockRows: sqlmock.NewRows([]string{"id", "name", "description", "version"}).
				AddRow(1, "test-service-1", "Description of test-service-1", "1.0").
				AddRow(2, "test-service-2", "Description of test-service-2", "1.1"),
			wantErr: false,
			expected: []models.Service{
				{Id: "1", Name: "test-service-1", Description: "Description of test-service-1", Version: "1.0"},
				{Id: "2", Name: "test-service-2", Description: "Description of test-service-2", Version: "1.1"},
			},
		},
		{
			name: "query error",
			queryParams: models.QueryParams{
				Name: "test-service",
			},
			mockQuery: "SELECT * FROM services WHERE name = test-service",
			mockRows:  nil, // No rows will be returned
			wantErr:   true,
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the query
			if tt.mockRows != nil {
				mock.ExpectQuery(tt.mockQuery).WillReturnRows(tt.mockRows)
			}
			// Call GetServices method
			services, err := ps.GetServices(tt.queryParams)

			// Check for errors
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, services)
			} else {

				require.NoError(t, err)
				assert.Equal(t, tt.expected, services)
			}

			// Ensure all expectations were met
			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}
