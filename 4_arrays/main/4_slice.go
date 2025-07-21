package main

import "fmt"

func main() {
	// primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice range includes the first element, but excludes the last one.
	// var s []int = primes[1:4]
	// fmt.Println("Printing Slice")
	// fmt.Println(s)

	// changing slice also changes the original array
	// s[0] = 100
	// fmt.Println("Printing Array after modification")
	// fmt.Println(primes)
	// fmt.Println("Printing Slice after modification")
	// fmt.Println(s)

	// array literal
	// names := [3]string{"John", "Arya", "Mary"}
	// fmt.Println(names)

	// slice literals
	// namesSlice := []string{"John", "Arya", "Mary"}
	// fmt.Println(namesSlice)

	// slice default values
	// a := names[0:3]
	// a := names[:3]
	// a := names[0:]
	// a := names[:]
	// fmt.Println(a)

	// slice first always creates an underlying array if used as a literal !
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(s)
	// s = s[:5]
	// fmt.Println(s)
	// fmt.Printf("len=%d, cap=%d", len(s), cap(s))

	// Dynamically creating zeroes arrays

	// s := make([]int, 5)
	// fmt.Printf("Len=%d, Cap=%d\n", len(s), cap(s))

	// r := make([]int, 1, 5)
	// fmt.Printf("Len=%d, Cap=%d\n", len(r), cap(r))

	// append data to a slice

	// s is a nil slice
	// var s []int
	// fmt.Println(s)

	// s = append(s, 0)
	// fmt.Println(s)

	// s = append(s, 1)
	// fmt.Println(s)

	// s = append(s, 2, 3, 4, 5, 6)
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)
	// fmt.Printf("Len=%d, Cap=%d\n", len(s), cap(s))

	// s = s[:]
	// fmt.Println(s)

	// r := make([][]int, 2, 10)
	// fmt.Println(r)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[:1]
	fmt.Printf("Len=%d, Cap=%d\n", len(s), cap(s))

}
