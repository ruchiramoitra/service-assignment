package v1

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	serviceHandler := &ServiceHandler{DB: db}

	router.HandleFunc("/v1/services", serviceHandler.GetServices).Methods("GET")
}
