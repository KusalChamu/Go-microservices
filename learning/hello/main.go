package main

import "fmt"

type Product struct {
	ID int
	Name string
	Price float64
	Stock int
}

func main() {
	p1 := Product{
		ID :1,
		Name :"laptop",
		Price :999.99,
		Stock :5,
	}
	

	fmt.Println(p1)

	products := []Product{
		{ID :1,
		Name :"laptop",
		Price :999.99,
		Stock :5},
		{ID :2,
		Name :"mouse",
		Price :9.99,
		Stock :50},
	}

	for _, p := range products {
		fmt.Printf("%d: %s - $%.2f (%d in stock)\n", p.ID, p.Name, p.Price, p.Stock)
	}

}