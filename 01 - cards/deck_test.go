package main

import (
	"os"
	"testing"
)

var initialCardCount int = 52

func TestNewDeckCardAmount(t *testing.T) {
	expectedLength := initialCardCount

	d := newDeck()

	actualLength := len(d)
	if actualLength != expectedLength {
		t.Errorf("Expected length of %v but got %v", expectedLength, actualLength)
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	testCard(newDeck(), 0, "2♣", t)
}

func TestNewDeckLastCard(t *testing.T) {
	d := newDeck()
	testCard(d, len(d)-1, "A♦", t)
}

func testCard(d deck, i int, expectedCard string, t *testing.T) {
	actualCard := d[i]

	if actualCard != expectedCard {
		t.Errorf("Expected card %v but got %v", expectedCard, actualCard)
	}
}

func TestSaveToDeck(t *testing.T) {
	fileName := "_newdecktestingsave"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	if _, err := os.Stat(fileName); err != nil {
		t.Errorf("File was not created.")
	}
	os.Remove(fileName)
}

func TestNewDeckFromFile(t *testing.T) {
	fileName := "_newdecktestingload"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if loadedDeck == nil {
		t.Errorf("File was not loaded.")
	} else if len(loadedDeck) != initialCardCount {
		t.Errorf("File did not contain expected amount of cards.")
	}
	os.Remove(fileName)
}
