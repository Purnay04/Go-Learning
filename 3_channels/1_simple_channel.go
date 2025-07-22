package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) // create a channel of type string

	// Launch a goroutine that sends a message on the channel
	go func() {
		ch <- "Hello from goroutine"
	}()

	// Receive the message from the channel in main goroutine
	msg := <-ch
	fmt.Println(msg)
}
