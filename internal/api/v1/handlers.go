package v1

import (
	"encoding/json"
	"fmt"
	"kong-assignment/internal/models"
	"kong-assignment/internal/storage"
	"net/http"
)

type ServiceHandler struct {
	serviceRepo storage.ServiceRepo
}

func (handler *ServiceHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	sort := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	queryParams := models.QueryParams{
		Name:   name,
		Sort:   sort,
		Limit:  limit,
		Offset: offset,
	}
	services, err := handler.serviceRepo.GetServices(queryParams)
	if err != nil {
		fmt.Println("Error getting data: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)

}
