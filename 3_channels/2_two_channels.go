package main

import (
	"fmt"
	"sync"
)

func sender(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sender sending a msg:")
	ch <- "Hello from sender goroutine"
}

func receiver(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Println("Receiver received:", msg)
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2) // We have two goroutines to wait for

	go sender(ch, &wg)
	go receiver(ch, &wg)

	wg.Wait() // Wait for both sender and receiver to finish
}
