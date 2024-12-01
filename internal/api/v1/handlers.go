package v1

import (
	"encoding/json"
	"fmt"
	"kong-assignment/internal/models"
	"kong-assignment/internal/storage"
	"net/http"
)

type ServiceHandler struct {
	ServiceRepo storage.ServiceRepo
}

func (handler *ServiceHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	paginationToken := r.URL.Query().Get("pagination_token")

	queryParams := models.QueryParams{
		Sort:            sort,
		Limit:           limit,
		Offset:          offset,
		PaginationToken: paginationToken,
	}
	services, paginationToken, err := handler.ServiceRepo.GetServices(queryParams)
	if err != nil {
		fmt.Println("Error getting data: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode the services and pagination token to JSON and write it to the response writer

	response := struct {
		Services        []models.Service `json:"services"`
		PaginationToken string           `json:"pagination_token"`
	}{
		Services:        services,
		PaginationToken: paginationToken,
	}
	json.NewEncoder(w).Encode(response)
}

func (handler *ServiceHandler) SearchService(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")

	if name == "" && id == "" {
		http.Error(w, "name or id is required", http.StatusBadRequest)
		return
	}
	queryParams := models.QueryParams{
		Name: name,
		Id:   id,
	}
	services, err := handler.ServiceRepo.SearchService(queryParams)
	if err != nil {
		fmt.Println("Error getting data: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}
