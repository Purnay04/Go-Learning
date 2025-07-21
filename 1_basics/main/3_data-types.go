package main

import "fmt"

func main() {
	var i int
	var f float32
	var s string
	var b bool

	// no garbage values
	fmt.Println(i, f, s, b)
	fmt.Printf("i=%d, f=%2f", i, f)

	i = 10
	// all type conversions are explicit
	f = float32(i)
	fmt.Println(f)

	fmt.Printf("var f is of type %T", f)

}
