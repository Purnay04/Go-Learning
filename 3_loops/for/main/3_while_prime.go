package main

import "fmt"

func main() {
	num := 23
	i := 2

	for i < num {
		if num%i == 0 {
			fmt.Printf("%d is not prime.", num)
			fmt.Println()
			break
		} else {
			i++
		}
	}

	if num == i {
		fmt.Printf("%d is prime.", num)
		fmt.Println()
	}

}
