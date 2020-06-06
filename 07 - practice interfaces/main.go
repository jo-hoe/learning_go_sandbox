package main

import (
	"fmt"
	"io"
)

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
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return (t.hight * t.base) / 2
}

func printArea(s shape, wr io.Writer) {
	message := fmt.Sprintf("The area is: %f", s.getArea())
	wr.Write([]byte(message))
}
