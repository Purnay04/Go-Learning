package main

import "fmt"

// Shape interface with an Area method
type Shape interface {
	Area() float64
}

// Circle struct implementing Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Square struct implementing Shape
type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func printShapeDetails(shape Shape) {
	// Try to assert as Circle
	if c, ok := shape.(Circle); ok {
		fmt.Println("This is a Circle with radius:", c.Radius)
	} else if s, ok := shape.(Square); ok {
		fmt.Println("This is a Square with side:", s.Side)
	} else {
		fmt.Println("Unknown shape type.")
	}
}

func main() {
	var s Shape

	s = Circle{Radius: 2.5}
	printShapeDetails(s) // Output: This is a Circle with radius: 2.5

	s = Square{Side: 4.0}
	printShapeDetails(s) // Output: This is a Square with side: 4
}
