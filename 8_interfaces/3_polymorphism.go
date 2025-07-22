package main

import "fmt"

// Shape interface
type Shape interface {
	area()
}

// Circle struct
type Circle struct {
	radius float64
}

// Square struct
type Square struct {
	side float64
}

// Circle implements Shape
func (c Circle) area() {
	fmt.Printf("Area of Circle: %.2f\n", 3.14*c.radius*c.radius)
}

// Square implements Shape
func (s Square) area() {
	fmt.Printf("Area of Square: %.2f\n", s.side*s.side)
}

// Print area and type of the struct
func printArea(shape Shape) {
	fmt.Printf("Actual type: %T\n", shape)
	shape.area()
}

func main() {
	var shape Shape

	shape = Circle{radius: 5.0}
	printArea(shape)

	shape = Square{side: 4.0}
	printArea(shape)
}
