package main

import "fmt"

type Address struct {
	city string
	line string
	pin  string
}

type Employee struct {
	id      string
	name    string
	salary  float32
	address Address
}

func main() {

	a := Address{"Pune", "Some address line", "123456"}
	emp := Employee{"QWA-2435", "John Doe", 12345.6, a}

	// emp := Employee{"QWA-2435", "John Doe", 12345.6, Address{"Pune", "Some address line", "123456"}}
	fmt.Println(emp)
}
