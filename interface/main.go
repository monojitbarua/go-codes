package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (s circle) area() float64 {
	return 3.14 * s.radius * s.radius
}

func shapeInfo(s shape) {
	fmt.Printf("Shape Type: %T\n", s)
	fmt.Println("Area of shape: ", s.area())
}

func main() {

	s := square{side: 5.0}
	c := circle{radius: 5.0}

	shapeInfo(s)
	shapeInfo(c)
}
