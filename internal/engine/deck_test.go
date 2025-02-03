package engine

import (
	"slices"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := NewDeck()
	if deck == nil {
		t.Fatal("deck is nil")
	}

	t.Logf("\nDeck: \n%s", deck)

	testAgainstStandardDeck(t, deck)
}

func testAgainstStandardDeck(t *testing.T, deck *Deck) {
	standardDeck := newStandardDeck()
	numOfCardsInStandardDeck := len(standardDeck)
	numOfCards := len(deck.Cards)
	if numOfCards != numOfCardsInStandardDeck {
		t.Fatalf("A deck should have %d cards, but this one has %d cards", numOfCardsInStandardDeck, numOfCards)
	}

	for _, card := range deck.Cards {
		i := slices.Index(standardDeck, card)
		if i == -1 {
			t.Fatalf("Deck has unknown card: %s", card)
		} else {
			standardDeck = slices.Delete(standardDeck, i, i+1)
		}
	}

	if len(standardDeck) != 0 {
		t.Fatalf("Deck does not conform to a standard deck.")
	}
}
