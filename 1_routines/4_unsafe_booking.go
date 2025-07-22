package main

import (
	"fmt"
	"sync"
)

type Ticket struct {
	tickets int
}

func (t *Ticket) BookTicket(name string) {
	fmt.Println("==============================")
	fmt.Println("Booking a ticket for: ", name)
	if t.tickets > 0 {
		t.tickets--
		fmt.Println("Ticket is booked for: ", name)
		fmt.Println("Tickets Available= ", t.tickets)
	} else {
		fmt.Println("Tickets Sold Out.")
	}
	fmt.Println("==============================")
}

func main() {
	ticket := &Ticket{tickets: 10}
	var wg sync.WaitGroup

	for i := 1; i <= 12; i++ {
		wg.Add(1)
		name := fmt.Sprintf("User-%02d", i)
		go func(n string) {
			defer wg.Done()
			ticket.BookTicket(n)
		}(name)
	}

	wg.Wait()
}
