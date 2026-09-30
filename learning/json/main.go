package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"in_stock"`
	Internal string  `json:"-"` // "-" means: never include this field in JSON
}

func main() {
	p := Product{ID: 1, Name: "Laptop", Price: 999.99, InStock: true, Internal: "secret"}

	// Go struct -> JSON bytes
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println(string(data))
	// {"id":1,"name":"Laptop","price":999.99,"in_stock":true}

	// JSON bytes -> Go struct
	incoming := []byte(`{"id":2,"name":"Mouse","price":15.5,"in_stock":false}`)
	var p2 Product
	if err := json.Unmarshal(incoming, &p2); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Println(p2)
	// {2 Mouse 15.5 false }
}