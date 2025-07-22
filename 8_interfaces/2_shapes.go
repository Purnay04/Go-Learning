package main

import "fmt"

type Shape interface {
	area()
}

type Circle struct {
	radius float32
}

func (circle Circle) area() {
	fmt.Printf("Area of circle=%f\n", 3.14*circle.radius*circle.radius)
}

type Square struct {
	side int
}

func (square Square) area() {
	fmt.Printf("Area of square=%d\n", square.side*square.side)
}

// i can pass circle or a square or any struct which implements Shape
func findArea(shape Shape) {
	shape.area()
}

func main() {

	c := Circle{5.66}
	s := Square{6}

	// c.area()
	// s.area()

	findArea(c)
	findArea(s)

}

// Shape -- Circle
// Shape -- Square
// Shape s = new Circle()
// Shape s = new Square()
