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

	collectedDeck := make([]Card, 0, 54)
	for _, hand := range game.Hands {
		collectedDeck = append(collectedDeck, hand.InHand...)
		collectedDeck = append(collectedDeck, hand.FaceUp...)
		collectedDeck = append(collectedDeck, hand.FaceDown...)
	}
	collectedDeck = append(collectedDeck, game.Deck.Cards...)

	testAgainstStandardDeck(t, &Deck{collectedDeck})

	game.init()
	t.Log(game)

	collectedInHand := make([]Card, 0, numOfPlayers*3)
	for _, hand := range game.Hands {
		collectedInHand = append(collectedInHand, hand.InHand...)
	}
	slices.SortFunc(collectedInHand, BasicCompare)
	t.Logf("All in hand cards: %s", collectedInHand)

	lowestCard := collectedInHand[0]
	t.Logf("Lowest card: %s", collectedInHand[0])

	startingHand := game.Hands[game.currentPlayerId]
	if !slices.Contains(startingHand.InHand, lowestCard) {
		t.Fatalf("Player does not have the lowest card. lowestCard: %s, startingHand: %v", lowestCard, startingHand)
	}
}

//TODO: Add tests for comparison code
