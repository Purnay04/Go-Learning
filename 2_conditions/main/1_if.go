package main

import (
	"fmt"
)

func main() {
	var num int = 0
	fmt.Println("Enter number:")
	fmt.Scanln(&num)

	if num >= 18 {
		fmt.Printf("%d is allowed to vote", num)
	}

}
