package main

import "fmt"

func swap(x, y *string) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a, b := "Hello", "World"
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
