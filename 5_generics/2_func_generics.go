package main

import "fmt"

// Contains returns true if the value exists in the slice
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func main() {
	ints := []int{1, 2, 3, 4}
	fmt.Println(Contains(ints, 3)) // true

	words := []string{"apple", "banana", "cherry"}
	fmt.Println(Contains(words, "grape")) // false

	floats := []float64{2.1, 3.3, 4.4}
	fmt.Println(Contains(floats, 4.4)) // true
}
