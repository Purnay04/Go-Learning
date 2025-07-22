package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from a goroutine!")
}

func main() {
	go sayHello() // Launches sayHello() in a new goroutine
	fmt.Println("Hello from main!")

	// main thread may exit before the goroutine runs
	// To prevent this, we can add a sleep or use a WaitGroup
	time.Sleep(1 * time.Second)
}
