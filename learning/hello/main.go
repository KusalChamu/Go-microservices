package main

import "fmt"



func main() {
	
	prices:=map[string]float64{
		"laptop":999.39,
		"mouse":56.88,
	}
	

	for name,price := range prices{
		fmt.Println(name,price)
	}

}