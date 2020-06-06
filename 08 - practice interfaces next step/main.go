package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeFileToConsoleViaCopy()
}

func writeFileToConsoleViaCopy() {
	arg := os.Args[1]
	file, err := os.Open(arg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

// Function does not follow the exercise. However, this logic came first to my mind.
// It works, but may perform worse (time complexity wise).
// This is an assumption however, to validate this the implementation of io.Copy
// has to be checked or performance test have to be written.
func writeFileToConsoleViaBuffer() {
	arg := os.Args[1]
	bS, err := ioutil.ReadFile(arg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bS))
}
