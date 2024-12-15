package api

import (
	"encoding/json"
	"net/http"
	"product-management-system/services"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product services.ProductInput
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdProduct, err := services.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdProduct)
}
