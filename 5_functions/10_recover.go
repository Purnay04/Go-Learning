package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	defer handlePanic()
	panic("Panic Reason: Something went wrong!")
}
