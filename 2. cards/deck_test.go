package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {

	a := newDeck()
	if len(a) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(a))
	}

	if a[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", a[0])
	}

	if a[len(a)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", a[0])
	}
}
