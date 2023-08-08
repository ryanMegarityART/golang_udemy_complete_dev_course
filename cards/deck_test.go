package main

import (
	"errors"
	"os"
	"strconv"
	"testing"
)

func validateStandardDeck(d deck) error {

	if (len(d)) != 12 {
		return errors.New("Expected deck length of 12 but got: " + strconv.Itoa(len(d)))
	}

	if d[0] != "Ace of Spades" {
		return errors.New("Expected first card to be 'Ace of Spades' but got: " + d[0])
	}

	return nil
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	err := validateStandardDeck(d)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	err := deck.saveToFile("_deckTesting")
	if err != nil {
		t.Errorf(err.Error())
	}

	loadedDeck := newDeckFromFile("_deckTesting")
	err = validateStandardDeck(loadedDeck)
	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove("_deckTesting")
}
