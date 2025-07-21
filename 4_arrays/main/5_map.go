package main

import "fmt"

func main() {

	map1 := map[int]string{
		11: "Dog",
		2:  "Cat",
		31: "Cow",
		4:  "Bird",
		33: "Rabbit",
	}

	fmt.Println("Map1: ", map1)

	// print a particular key
	fmt.Print("Key-11, value=", map1[11])
	fmt.Println()

	names := make(map[int]string)
	names[10] = "John D."
	names[4] = "Derek K."
	fmt.Println("Names=", names)

	fmt.Println("Printing map using range for")
	for i, pet := range map1 {
		fmt.Println(i, pet)
	}
}
