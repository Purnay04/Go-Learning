package main

import "fmt"

// func divRes(num int) (int, int)
func divRes(num int) (q, r int) {
	q = num / 2
	r = num % 2
	return
}

func main() {
	// Using named return values
	q, r := divRes(11)
	fmt.Println(q, r)
	// Using naked return
	// This will return the same values as above
	fmt.Println(divRes(11))
}
