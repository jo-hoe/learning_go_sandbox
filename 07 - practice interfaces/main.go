package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	hight float64
	base  float64
}

func main() {
	s := square{sideLength: 2}
	printArea(s)

	t := triangle{hight: 2, base: 3}
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return (t.hight * t.base) / 2
}

func printArea(s shape) {
	fmt.Println("The area is:", s.getArea())
}
