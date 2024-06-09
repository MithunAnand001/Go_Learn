package main

import (
	"learn_service/handler"
	"learn_service/repository"
	"learn_service/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    // Create a new router
    router := mux.NewRouter()
	TestingHandler := handler.TestingHandler{TestingService: service.NewTestingService(repository.NewTestingRepository())}
    // Define a route and its handler
    router.HandleFunc("/hello", TestingHandler.HealthCheck).Methods("GET")

	router.HandleFunc("/healthcheck",TestingHandler.HealthCheck).Methods("OPTIONS")

    // Start the HTTP server and specify the router as the request handler
    http.ListenAndServe(":8080", router)
}