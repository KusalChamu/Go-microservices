package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello from the server!")
}

func main(){
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("Server starting on http://localhost:8080")

	err :=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
