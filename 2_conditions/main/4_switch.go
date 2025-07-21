package main

import (
	"fmt"
)

func main() {

	var x int = 0
	var y int = 0
	var ch int = 0
	var res float32 = 0

	fmt.Println("Enter x:")
	fmt.Scanln(&x)

	fmt.Println("Enter y:")
	fmt.Scanln(&y)

	fmt.Println("1. Add")
	fmt.Println("2. Sub")
	fmt.Println("3. Mul")
	fmt.Println("4. Div")
	fmt.Println("Enter Choice")

	fmt.Scanln(&ch)

	switch ch {
	case 1:
		res = (float32)(x + y)
		fmt.Printf("%d + %d = %f", x, y, res)
		break

	case 2:
		res = (float32)(x - y)
		fmt.Printf("%d - %d = %f", x, y, res)
		break

	case 3:
		res = (float32)(x * y)
		fmt.Printf("%d * %d = %f", x, y, res)
		break

	case 4:
		res = float32(x) / float32(y)
		fmt.Printf("%d / %d = %f", x, y, res)
		break
	}

}
