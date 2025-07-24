package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
