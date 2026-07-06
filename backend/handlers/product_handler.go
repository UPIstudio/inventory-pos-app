package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luthfi/inventory-pos-app/backend/models"
	"github.com/luthfi/inventory-pos-app/backend/repositories"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method tidak diterima", http.StatusMethodNotAllowed)
		return
	}

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.SaveProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
