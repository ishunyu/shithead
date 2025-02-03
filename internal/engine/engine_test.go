package engine

import (
	"slices"
	"testing"
)

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
