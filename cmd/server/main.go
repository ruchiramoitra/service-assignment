package main

import (
	"kong-assignment/config"
	v1 "kong-assignment/internal/api/v1"
	"kong-assignment/internal/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	dbInstance, err := storage.NewPostgres(config)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	service := storage.NewPostgresStorage(dbInstance)

	r := mux.NewRouter()

	v1.RegisterRoutes(r, service)

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
