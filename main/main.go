package main

import (
	"log"
	"net/http"
	"product-management-system/api"
	"product-management-system/configs"
)

func main() {
	// Load environment variables
	configs.LoadConfig()

	// Initialize routes
	router := api.SetupRoutes()

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
