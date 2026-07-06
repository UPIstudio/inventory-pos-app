package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luthfi/inventory-pos-app/backend/repositories"
)

func GetProductsHandler(w http.ResponseWriter, e *http.Request) {
	products, err := repositories.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
