package v1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"kong-assignment/internal/models"
	"net/http"
)

type ServiceHandler struct {
	DB *sql.DB
}

func (handler *ServiceHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("name")
	sort := r.URL.Query().Get("sort")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	query := "SELECT * FROM services"
	if nameFilter != "" {
		query += " WHERE name = $1"
	}
	if sort != "" {
		query += " ORDER BY " + sort
	}
	if limit != "" {
		query += " LIMIT " + limit
	}
	if offset != "" {
		query += " OFFSET " + offset
	}

	fmt.Println(query)

	rows, err := handler.DB.Query(query)
	if err != nil {
		fmt.Println("Error querying database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var services []models.Service

	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.Description, &service.Name, &service.Version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)

}
