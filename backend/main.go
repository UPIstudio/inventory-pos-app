package main

import (
	"log"
	"net/http"

	"github.com/luthfi/inventory-pos-app/backend/config"
	"github.com/luthfi/inventory-pos-app/backend/handlers"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetProductsHandler(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateProductHandler(w, r)
		} else {
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handlers.UpdateProductHandler(w, r)
		} else if r.Method == http.MethodDelete {
			handlers.DeleteProductHandler(w, r)
		} else {
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	select {}
}
