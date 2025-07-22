package main

import (
	"fmt"
	"sync"
	"time"
)

func chat(name string, recv <-chan string, send chan<- string, messages []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range messages {
		send <- fmt.Sprintf("%s: %s", name, msg)
		reply := <-recv
		fmt.Println(reply)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	var wg sync.WaitGroup

	aliceToBob := make(chan string, 1)
	bobToAlice := make(chan string, 1)

	wg.Add(2) // We'll launch two goroutines

	go chat("Alice", bobToAlice, aliceToBob, []string{
		"Hi Bob!",
		"How are you?",
		"Nice weather today.",
		"See you!",
	}, &wg)

	go chat("Bob", aliceToBob, bobToAlice, []string{
		"Hi Alice!",
		"I'm good, thanks. You?",
		"Indeed, it's sunny.",
		"Bye!",
	}, &wg)

	wg.Wait() // Wait for both chat goroutines to finish
}
