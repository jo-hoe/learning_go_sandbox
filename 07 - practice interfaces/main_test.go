package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSquareArea(t *testing.T) {
	s := square{sideLength: 8}
	expectedArea := 64.

	actualArea := s.getArea()

	if actualArea != expectedArea {
		t.Errorf("Expected area of %v, but received %v", expectedArea, actualArea)
	}
}

func TestTriangleArea(t *testing.T) {
	tri := triangle{base: 12, hight: 12}
	expectedArea := 72.

	actualArea := tri.getArea()

	if actualArea != expectedArea {
		t.Errorf("Expected area of %v, but received %v", expectedArea, actualArea)
	}
}

type demo struct{}

func (d demo) getArea() float64 {
	return 1.
}

func TestPrintArea(t *testing.T) {
	var d demo

	expected := fmt.Sprintf("The area is: %f", d.getArea())
	buf := new(bytes.Buffer)

	printArea(d, buf)

	actual := buf.String()

	if expected != actual {
		t.Errorf("Expected [%v], but received [%v]", expected, actual)
	}
}
