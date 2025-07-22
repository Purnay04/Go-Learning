package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker reads and prints the context value
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	username := ctx.Value("user")
	fmt.Printf("Worker %d received user from context: %v\n", id, username)
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", "alice")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}
	wg.Wait()
}
