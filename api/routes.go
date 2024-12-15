package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Product routes
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", GetProductByID).Methods("GET")
	router.HandleFunc("/products", GetProducts).Methods("GET")

	return router
}

// GetProductByID retrieves a product by its ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	// Mock response for now
	product := map[string]interface{}{
		"id":          productID,
		"name":        "Sample Product",
		"description": "This is a sample product.",
		"price":       100.0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProducts retrieves all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Mock response for now
	products := []map[string]interface{}{
		{"id": 1, "name": "Product A", "price": 50.0},
		{"id": 2, "name": "Product B", "price": 75.0},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
