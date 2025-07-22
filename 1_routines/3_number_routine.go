package main

import (
	"fmt"
	"sync"
)

// Function that prints numbers 1-5, with the goroutine's name
func printNumbers(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
	}
}

func main() {
	var wg sync.WaitGroup

	// Launch 3 goroutines concurrently
	wg.Add(3)
	go printNumbers("Goroutine-One", &wg)
	go printNumbers("Goroutine-Two", &wg)
	go printNumbers("Goroutine-Three", &wg)

	wg.Wait()
	fmt.Println("All goroutines complete!")
}
