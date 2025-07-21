package main

import "fmt"

// They can be used as a callback or passed as arguments to other functions.
func compute(operation func(int, int) int) int {
	return operation(10, 2)
}

func main() {
	// Anonymous functions are functions that are defined without a name.
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 5)) // Output: 8

	// Anonymous functions can be assigned to variables.
	result := func(x int, y int) int {
		return x * y
	}(2, 4)
	fmt.Println(result) // Output: 8

	// Anonymous functions can be passed as arguments to other functions.
	// Similar to lambda functions or arrow functions in other languages.
	// Also known as Higher-order functions.
	result = compute(func(a, b int) int {
		return a - b
	})
	fmt.Println(result) // Output: 8

}
