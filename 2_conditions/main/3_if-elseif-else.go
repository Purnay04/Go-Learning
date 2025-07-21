package main

import (
	"fmt"
)

func main() {
	var num int = 0
	fmt.Println("Enter number:")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Printf("%d is +ve", num)
	} else if num < 0 {
		fmt.Printf("%d is -ve", num)
	} else {
		fmt.Printf("Num is 0")
	}

}
