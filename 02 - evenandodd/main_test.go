package main

import "testing"

func TestIsEven(t *testing.T) {
	if isEven(8) != true {
		t.Errorf("Number 8 was not correctly detected as even")
	}
}

func TestIsOdd(t *testing.T) {
	if isEven(9) != false {
		t.Errorf("Number 9 was not correctly detected as odd")
	}
}

func TestZeroIsEven(t *testing.T) {
	if isEven(0) != true {
		t.Errorf("Number 9 was not correctly detected as even")
	}
}
