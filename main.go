package main

import (
	"learn_service/automigrate"
	"learn_service/handler"
	"learn_service/repository"
	"learn_service/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	automigrate.MigrateTables()
    // Create a new router
    router := mux.NewRouter()
	TestingHandler := handler.TestingHandler{TestingService: service.NewTestingService(repository.NewTestingRepository())}
    // Define a route and its handler
    router.HandleFunc("/hello", TestingHandler.HealthCheck).Methods("GET")

	router.HandleFunc("/healthcheck",TestingHandler.HealthCheck).Methods("OPTIONS")

    // Start the HTTP server and specify the router as the request handler
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }}