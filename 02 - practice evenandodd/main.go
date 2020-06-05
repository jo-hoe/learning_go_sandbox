package main

import "fmt"

func main() {
	var numbers []int

	// generate numbers
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	// calculate result
	for _, n := range numbers {
		fmt.Printf("%v is ", n)
		if isEven(n) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}

func isEven(number int) bool {
	return number%2 == 0
}
