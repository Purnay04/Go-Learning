package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {

    sum, diff := 0, 0
    sum = add(42, 13)
    diff = sub(42, 13)
	fmt.Println(sum)
	fmt.Println(diff)
}