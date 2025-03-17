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

	ids := make([]int, 0, numOfPlayers)
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
	collectedDeck = append(collectedDeck, game.DrawPile.Cards...)

	testAgainstStandardDeck(t, &Deck{collectedDeck})
}

func TestInitGame(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)

	game.Init()
	t.Log(game)

	collectedInHand := make([]Card, 0, numOfPlayers*3)
	for _, hand := range game.Hands {
		collectedInHand = append(collectedInHand, hand.InHand...)
	}
	slices.SortFunc(collectedInHand, NumericCompare)
	t.Logf("All in hand cards: %s", collectedInHand)

	lowestCard := collectedInHand[0]
	t.Logf("Lowest card: %s", collectedInHand[0])

	startingHand := game.Hands[game.currentPlayerId]
	if !slices.Contains(startingHand.InHand, lowestCard) {
		t.Fatalf("Player does not have the lowest card. lowestCard: %s, startingHand: %v", lowestCard, startingHand)
	}
}

func TestPlayHandSuccess(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	game.Init()
	startingHand := &game.Hands[game.currentPlayerId]

	play := Play{
		Hand: startingHand,
		Card: minSlice(startingHand.InHand, NumericCompare),
	}

	result := game.PlayHand(play)
	if !result.Success {
		t.Fatal("Expected play to succeed, but it failed.")
	}

	if result.Round != 1 {
		t.Fatalf("Round number mismatch. Expected: 1, actual: %d.", result.Round)
	}

	if !game.nextTo(result.NextPlayerId, startingHand.Id) {
		t.Fatalf("Next player mismatch. startingPlayerId: %d, nextPlayerId: %d.", startingHand.Id, result.NextPlayerId)
	}

	if slices.Contains(startingHand.InHand, play.Card) {
		t.Fatalf("Card should not be found in starting hand. Card: %s, startingHand: %v", play.Card, startingHand.InHand)
	}

	if len(startingHand.InHand) != 3 {
		t.Fatalf("Starting hand should still have 3 cards. Actual: %d.", startingHand.InHand)
	}
}

func TestPlayHandFail(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	startingHand := &game.Hands[game.currentPlayerId]

	play := Play{
		Hand: &game.Hands[game.leftOf(startingHand.Id)],
		Card: minSlice(startingHand.InHand, NumericCompare),
	}
	result := game.PlayHand(play)
	if result.Success {
		t.Fatal("Expected play to fail, but it succeeded.")
	}
}
