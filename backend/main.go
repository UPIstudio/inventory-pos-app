package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo dari Backend Golang yang hidup!")
	})

	fmt.Println("Backend berjalan di port 8080...")
	http.ListenAndServe(":8080", nil)
}
