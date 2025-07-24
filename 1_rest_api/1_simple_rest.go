package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	// Mock: create an example slice of products
	products := []Product{
		{ID: 1, Name: "Sample1", Price: 10.0, Category: "Category1"},
		{ID: 2, Name: "Sample2", Price: 20.0, Category: "Category2"},
	}
	fmt.Println("Retrieve All:", products)
	resp, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Create:", p)
	resp, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)
	p := Product{
		ID:       id,
		Name:     "Sample",
		Price:    123.45,
		Category: "Mock",
	}
	fmt.Println("Retrieve:", p)
	resp, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	// Temp struct to decode incoming JSON (all fields optional if JSON omits them)
	var input struct {
		Name     *string  `json:"name"`
		Price    *float64 `json:"price"`
		Category *string  `json:"category"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mock existing product (in real cases you would fetch it from db)
	p := Product{
		ID:       id,
		Name:     "Old Name",
		Price:    100.0,
		Category: "Old Category",
	}

	// Set each field if new data is provided in input
	if input.Name != nil {
		p.Name = *input.Name
	}
	if input.Price != nil {
		p.Price = *input.Price
	}
	if input.Category != nil {
		p.Category = *input.Category
	}

	fmt.Println("Update:", p)
	resp, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)
	result := map[string]interface{}{
		"deleted": true,
		"id":      id,
	}
	fmt.Println("Delete:", result)
	resp, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", getAllProducts).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Println("Mock Product API running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
