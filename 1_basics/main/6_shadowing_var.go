package main

import "fmt"

func main() {
	x := 10 // Outer variable
	{
		x := 5         // Shadows outer x inside this block
		fmt.Println(x) // Output: 5
	}
	fmt.Println(x) // Output: 10
}
