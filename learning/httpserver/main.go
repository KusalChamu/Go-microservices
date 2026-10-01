package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// in-memory "database" for now — just a slice living in memory.
// It resets every time the server restarts. Real persistence comes in Phase 5 (PostgreSQL).
var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
}

var nextID = 2

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listProducts(w, r)
	case http.MethodPost:
		createProduct(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	// Decode the JSON request body into a Product struct.
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Basic validation.
	if newProduct.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if newProduct.Price <= 0 {
		http.Error(w, "price must be positive", http.StatusBadRequest)
		return
	}

	newProduct.ID = nextID
	nextID++
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newProduct)
}

func main() {
	http.HandleFunc("/products", productsHandler)

	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}