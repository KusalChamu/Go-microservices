package main

import "fmt"

func restock(p Product, amount int){
	p.Stock = p.Stock+amount
}

func restockPtr(p *Product, amount int){
	p.Stock = p.Stock+amount
}

type Product struct {
	Name string
	Stock int
}

func main() {
	laptop := Product{Name: "Laptop", Stock: 5}

	restock(laptop, 10)
	fmt.Println(laptop.Stock) // 5 — unchanged! restock only touched a copy

	restockPtr(&laptop, 10)
	fmt.Println(laptop.Stock) // 15 — changed, because we passed the address
}