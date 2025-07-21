package main

import "fmt"

func main() {
	fmt.Println("Before panic")
	panic("Something went wrong!")

	// Not executed
	fmt.Println("After panic")
}
