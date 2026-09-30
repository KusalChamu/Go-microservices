package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

// Method with a POINTER receiver — can modify the struct
func (p *Product) Restock(amount int) {
	p.Stock += amount
}

// Method with a VALUE receiver — only reads, cannot modify
func (p Product) Describe() string {
	return fmt.Sprintf("%s - $%.2f (%d in stock)", p.Name, p.Price, p.Stock)
}

func main() {
	laptop := Product{Name: "Laptop", Price: 999.99, Stock: 5}

	laptop.Restock(10)
	fmt.Println(laptop.Describe())
	// Laptop - $999.99 (15 in stock)
}