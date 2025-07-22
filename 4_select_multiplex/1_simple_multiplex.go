package main

import (
	"fmt"
	"time"
)

func main() {
	bluePhone := make(chan string)
	redPhone := make(chan string)

	// A friend will call your blue phone after 1 second
	go func() {
		time.Sleep(11 * time.Second)
		bluePhone <- "Hello from the BLUE phone!"
	}()

	// Another friend will call your red phone after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		redPhone <- "Hello from the RED phone!"
	}()

	// You wait: whichever phone rings first, you answer!
	select {
	case msg := <-bluePhone:
		fmt.Println("Answered BLUE phone:", msg)
	case msg := <-redPhone:
		fmt.Println("Answered RED phone:", msg)
	}
}
