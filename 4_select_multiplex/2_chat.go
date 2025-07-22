package main

import (
	"fmt"
	"sync"
	"time"
)

func chatter(
	name string,
	recv <-chan string,
	send chan<- string,
	greetings []string,
	done chan struct{}, // Now a non-read-only channel; any participant can close it
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for _, msg := range greetings {
		select {
		case send <- fmt.Sprintf("%s: %s", name, msg):
			// Sent
		case <-done:
			fmt.Printf("%s: Shutdown requested. Exiting.\n", name)
			return
		}

		select {
		case reply := <-recv:
			fmt.Printf("[%s]: got: %s\n", name, reply)
		case <-time.After(2 * time.Second):
			fmt.Printf("[%s]: timed out waiting for reply. Triggering shutdown.\n", name)
			close(done) // Initiate shutdown when timeout occurs
			return
		case <-done:
			fmt.Printf("%s: Shutdown requested while waiting. Exiting.\n", name)
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	aliceToBob := make(chan string, 1)
	bobToAlice := make(chan string, 1)
	done := make(chan struct{})

	aliceLines := []string{"Hi Bob!", "How are you?", "Nice weather today.", "See you!"}
	bobLines := []string{"Hi Alice!", "I'm good, thanks. You?", "Indeed, it's sunny."}

	wg.Add(2)
	go chatter("Alice", bobToAlice, aliceToBob, aliceLines, done, &wg)
	go chatter("Bob", aliceToBob, bobToAlice, bobLines, done, &wg)

	// Let them chat for more than 2 seconds, then trigger shutdown
	time.Sleep(3 * time.Second)
	close(done) // Shutdown happens after the set duration

	wg.Wait()
	fmt.Println("Chat ended gracefully.")
}
