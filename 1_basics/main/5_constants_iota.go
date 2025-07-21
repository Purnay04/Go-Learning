package main

import "fmt"

func main() {
	// Constants in Go
	const PI = 3.14
	const Greeting string = "Hello, World"

	fmt.Println("PI:", PI)
	fmt.Println("Greeting:", Greeting)

	const (
		A = 10
		B = "GoLang"
		C = true
	)

	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)

	// Typed and Inferred Constants
	const TypedInt int = 100  // typed
	const UntypedFloat = 3.14 // untyped (type inferred)
	fmt.Println("TypedInt:", TypedInt)
	fmt.Println("UntypedFloat:", UntypedFloat)

	// Iota Example
	const (
		Small  = iota // 0
		Medium        // 1
		Large         // 2
	)

	fmt.Println("Sizes:", Small, Medium, Large)

	// Iota with Custom Values and skipping values
	const (
		_      = iota      // 0 (skip)
		Bronze = 10 * iota // 10 (iota = 1)
		Silver             // 20 (iota = 2)
		Gold               // 30 (iota = 3)
	)

	fmt.Println("Medals:", Bronze, Silver, Gold)
}
