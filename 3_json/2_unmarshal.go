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
	jsonStr := `{"name":"Bob","age":25}`

	var p Person
	// Unmarshal JSON into struct
	if err := json.Unmarshal([]byte(jsonStr), &p); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p) // Output: {Name:Bob Age:25}
}
