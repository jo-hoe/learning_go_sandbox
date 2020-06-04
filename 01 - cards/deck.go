package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cardSuites := []string{"♣", "♠", "♥", "♦"}
	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	cards := deck{}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+""+suite)
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	bytes, error := ioutil.ReadFile(filename)
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	return deck(strings.Split(string(bytes), ";"))
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// switch two items in one line
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ";")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0200)
}
