package main

import (
	"fmt"
)

func main() {
	var num int = 0
	fmt.Println("Enter number:")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d is even", num)
	} else {
		fmt.Printf("%d is odd", num)
	}

	// Be careful with the intention of the else statement !

	if num > 0 {
		fmt.Println("Number is positive")
	}
	if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is 0")
	}

}
