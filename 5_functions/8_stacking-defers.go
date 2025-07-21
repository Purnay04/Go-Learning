package main

import "fmt"

func main() {
	fmt.Println("Start Main...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 0, 1, 2, 3, 4 .... 9
	}

	fmt.Println("Done...")
}
