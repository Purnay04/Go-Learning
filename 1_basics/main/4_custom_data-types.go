package main

import "fmt"

func main() {
	// Custom Data Types in Go
	type Age int

	var johnAge Age = 30
	fmt.Println("John's age is:", johnAge)

	// Alias for existing type
	type id = int
	var a id = 5
	fmt.Println("Id value is:", a)
}
