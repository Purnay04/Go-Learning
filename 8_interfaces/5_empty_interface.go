package main

import "fmt"

var anyValue interface{}

// similar to an Object class

func main() {
	anyValue = 42
	fmt.Println("anyValue:", anyValue)

	anyValue = "hello"
	fmt.Println("anyValue:", anyValue)

	anyValue = []int{1, 2, 3}
	fmt.Println("anyValue:", anyValue)

}
