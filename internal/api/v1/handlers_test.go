package v1_test

import (
	v1 "kong-assignment/internal/api/v1"
	"kong-assignment/internal/mocks"
	"kong-assignment/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for ServiceRepo
	mockServiceRepo := mocks.NewMockServiceRepo(ctrl)

	// Define the query parameters
	queryParams := models.QueryParams{
		Sort:   "asc",
		Limit:  "10",
		Offset: "0",
	}

	// Define the expected services to return from the mock
	expectedServices := []models.Service{
		{
			Id:            "1",
			Name:          "test-service-1",
			Description:   "Test service 1",
			Versions:      []string{"1.0", "1.1"},
			TotalVersions: 2,
		},
		{
			Id:            "2",
			Name:          "test-service-2",
			Description:   "Test service 2",
			Versions:      []string{"1.1"},
			TotalVersions: 1,
		},
	}

	// Set up the mock expectation for GetServices
	mockServiceRepo.EXPECT().
		GetServices(queryParams).
		Return(expectedServices, nil).
		Times(1)

	// Create the handler
	handler := &v1.ServiceHandler{
		ServiceRepo: mockServiceRepo,
	}

	req, err := http.NewRequest("GET", "/services?name=test-service&sort=asc&limit=10&offset=0", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.GetServices(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `[{"id":"1","name":"test-service-1","description":"Test service 1","versions":["1.0","1.1"],"total_versions":2},{"id":"2","name":"test-service-2","description":"Test service 2","versions":["1.1"],"total_versions":1}]`
	assert.JSONEq(t, expectedBody, rr.Body.String())

}
