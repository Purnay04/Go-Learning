package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("Hello from a goroutine!")
	wg.Done()
}

func main() {
	fmt.Print("Starting goroutine...\n")
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
}
