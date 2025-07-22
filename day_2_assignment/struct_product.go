package main

import (
	"fmt"
)

type Product struct {
	id    int
	name  string
	price float64
}

func main() {
	var nProd int

	fmt.Print("Enter number of products:")
	fmt.Scan(&nProd)

	products := make([]Product, nProd)

	for i := 0; i < nProd; i += 1 {
		p := Product{}

		fmt.Print("Enter ID:")
		fmt.Scan(&p.id)

		fmt.Print("Enter Name:")
		fmt.Scan(&p.name)

		fmt.Print("Enter Price:")
		fmt.Scan(&p.price)

		products[i] = p
		fmt.Println()
	}

	fmt.Println("Product List:")

	tot := 0.00
	for _, val := range products {
		tot += val.price
		fmt.Printf("ID: %d Name: %s Price: %.2f", val.id, val.name, val.price)
		fmt.Println()
	}
	fmt.Println()
	fmt.Printf("Average Price: %.2f", tot/float64(len(products)))
}
