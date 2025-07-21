package main

import "fmt"

func main() {

	res := 0
	for i := 1; i <= 10; i++ {
		res = i * 2
		fmt.Printf("2 * %d = %d", i, res)
		fmt.Println()
	}
}
