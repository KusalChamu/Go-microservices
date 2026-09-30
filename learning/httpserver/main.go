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

func productsHandler(w http.ResponseWriter, r *http.Request){
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 999.99},
		{ID: 2, Name: "Mouse", Price: 15.50},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}

}

func main(){
	http.HandleFunc("/products",productsHandler)
	fmt.Println("Server starting on http://localhost:8080")

	err :=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
