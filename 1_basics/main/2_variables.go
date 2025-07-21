package main

import "fmt"

func main() {
	// int x = 10
	// var x: number = 10
	var x int = 10
	var y, z int = 20, 30
	var lang string = "Go"

	// Observe the abbreviation
	msg := "Hey There !"

	// Observe the multi var declaration
	s1, s2 := "s1...", "s2..."

	fmt.Println(x, y)
	fmt.Printf("z = %d", z)
	fmt.Println()
	fmt.Println(lang)
	fmt.Println(msg)
	fmt.Println(s1, s2)
}
