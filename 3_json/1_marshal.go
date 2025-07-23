package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// Marshal struct to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData)) // Output: {"name":"Alice","age":30}
}
