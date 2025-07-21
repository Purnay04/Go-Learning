package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i
	var q *int = &j

	fmt.Println(i)
	fmt.Println(*p)

	fmt.Println(j)
	fmt.Println(*q)

	*p = 100
	fmt.Println(*p)
}
