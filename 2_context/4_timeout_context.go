package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	fmt.Println("Worker: started task...")
	select {
	case <-time.After(3 * time.Second): // Simulate a long-running task
		fmt.Println("Worker: finished task successfully!")
	case <-ctx.Done(): // Triggered if timeout/cancel occurs
		fmt.Println("Worker:", ctx.Err()) // context deadline exceeded
	}
}

func main() {
	// Create a context with 2-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx)

	// Wait enough time to see the timeout (main must stay alive)
	time.Sleep(7 * time.Second)
	fmt.Println("Main: exiting")
}
