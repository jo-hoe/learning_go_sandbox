package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type persianBot struct{}
type invalidBot struct{}

func main() {
	eb := englishBot{}
	pb := persianBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (persianBot) getGreeting() string {
	return "Salom!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
