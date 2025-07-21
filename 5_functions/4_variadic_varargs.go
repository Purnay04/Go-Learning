package main

import "fmt"

// func sum(nums []int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total += n
// 	}
// 	return total
// }

// sum is a variadic function that adds any number of integers

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {

	// Using the fixed-size array version
	// fmt.Println(sum([]int{1, 2, 3})) // 6
	// fmt.Println(sum([]int{10, 20}))  // 30
	// fmt.Println(sum([]int{}))        // 0

	// Using the variadic function
	fmt.Println(sum(1, 2, 3)) // 6
	fmt.Println(sum(10, 20))  // 30
	fmt.Println(sum())        // 0
}
