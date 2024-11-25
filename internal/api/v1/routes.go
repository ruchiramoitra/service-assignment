package v1

import (
	"kong-assignment/internal/storage"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, service storage.ServiceRepo) {
	serviceHandler := &ServiceHandler{serviceRepo: service}

	router.HandleFunc("/v1/services", serviceHandler.GetServices).Methods("GET")
}
