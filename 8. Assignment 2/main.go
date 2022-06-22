package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 2, height: 2}
	s := square{sideLength: 2}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	areaOfTriangle := 0.5 * t.height * t.base
	return areaOfTriangle
}

func (s square) getArea() float64 {
	areaOfSquare := s.sideLength * s.sideLength
	return areaOfSquare
}
