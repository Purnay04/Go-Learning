package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Reacts to cancellation
			fmt.Printf("Worker %d: cancelled (reason: %v)\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Launch 2 workers that listen for cancellation
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Let the workers run for a second
	time.Sleep(1 * time.Second)

	fmt.Println("Main: sending cancel signal!")
	cancel() // Signal cancellation

	wg.Wait()
	fmt.Println("All workers finished (cancelled).")
}
