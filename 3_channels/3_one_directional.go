package main

import (
	"fmt"
)

// Sender function: only sends data
func sender(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send values to channel
		fmt.Println("Sent:", i)
	}
	close(ch) // Close the channel when done
}

// Receiver function: only receives data
func receiver(ch <-chan int) {
	for val := range ch { // Receive values from channel
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int) // Create a bidirectional channel

	go sender(ch) // Pass as send-only to sender
	receiver(ch)  // Pass as receive-only to receiver
}
