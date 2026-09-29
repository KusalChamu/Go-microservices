package main

import "fmt"



func main() {
	
	var products []string
	products = append(products, "Laptop","mouse","keyboard")


	prices := []float64{19.99,5.99}
	fmt.Println(prices)

    fmt.Println(products[0:2])

}