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

	standardDeck := standardDeck()
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

func standardDeck() []Card {
	suits := []Suit{Club, Diamond, Heart, Spade}
	ranks := []Rank{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	cards := make([]Card, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{suit, rank})
		}
	}

	cards = append(cards, Card{JokerSmall, Joker})
	cards = append(cards, Card{JokerLarge, Joker})

	return cards
}

func TestGame(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	if game == nil {
		t.Fatal("game is nil")
	}

	numOfPlayersActual := len(game.Hands)
	if numOfPlayersActual != numOfPlayers {
		t.Fatalf("Number of players mismatch. Expected: %d, actual: %d.", numOfPlayers, numOfPlayersActual)
	}

	ids := make([]uint8, 0, numOfPlayers)
	for _, hand := range game.Hands {
		id := hand.Id
		if slices.Contains(ids, id) {
			t.Fatalf("Id %d already exists.", id)
		}
		ids = append(ids, id)

	}

}
