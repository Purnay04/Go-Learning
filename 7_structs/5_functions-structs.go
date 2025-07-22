package main

import "fmt"

type Point struct {
	x int
	y int
}

// this is a method of type Point
func (p Point) printPoint() {
	fmt.Println(p.x)
	fmt.Println(p.y)
}

func main() {
	point := Point{1, 2}
	point.printPoint()
}

/*

class Point{
    int x, y

    printPoint(){
        ....
    }
}

Point p = new Point()
p.printPoint()

*/
