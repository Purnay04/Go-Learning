package main

import (
	"fmt"
	"sync"
)

func chat(name string, send chan<- string, recv <-chan string, messages []string, startFirst bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range messages {
		if startFirst {
			send <- fmt.Sprintf("%s: %s", name, msg)
			fmt.Println(<-recv)
		} else {
			fmt.Println(<-recv)
			send <- fmt.Sprintf("%s: %s", name, msg)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	aliceToBob := make(chan string)
	bobToAlice := make(chan string)

	aliceMsgs := []string{"Hi Bob!", "How are you?", "Nice weather today.", "See you!"}
	bobMsgs := []string{"Hi Alice!", "I'm fine, thanks!", "Yes, it's great.", "Bye!"}

	wg.Add(2)

	go chat("Alice", aliceToBob, bobToAlice, aliceMsgs, true, &wg) // Alice sends first
	go chat("Bob", bobToAlice, aliceToBob, bobMsgs, false, &wg)    // Bob receives first

	wg.Wait()
}
