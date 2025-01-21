package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 16{
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as the first card, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs as the last card, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()

	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 160{
		t.Errorf("Expected %d cards in deck, but got %d", len(deck), len(loadedDeck))
	}

	os.Remove(filename)
}