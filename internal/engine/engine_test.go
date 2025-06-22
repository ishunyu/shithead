package engine

import (
	"slices"
	"testing"
)

func TestLeftOf(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	game.Init()

	for i := 0; i < numOfPlayers; i++ {
		leftPlayerId := game.leftOf(i)
		expectedLeftPlayerId := (i - 1 + numOfPlayers) % numOfPlayers
		if leftPlayerId != expectedLeftPlayerId {
			t.Errorf("Expected left of player %d to be %d, got %d", i, expectedLeftPlayerId, leftPlayerId)
		}
	}
}

func TestRightOf(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	game.Init()

	for i := 0; i < numOfPlayers; i++ {
		rightPlayerId := game.rightOf(i)
		expectedRightPlayerId := (i + 1) % numOfPlayers
		if rightPlayerId != expectedRightPlayerId {
			t.Errorf("Expected right of player %d to be %d, got %d", i, expectedRightPlayerId, rightPlayerId)
		}
	}
}

func TestNextPlayerId(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	game.Init()

	for i := 0; i < numOfPlayers; i++ {
		nextPlayerId := game.nextPlayerId()
		expectedNextPlayerId := (game.currentPlayerId + game.direction + numOfPlayers) % numOfPlayers
		if nextPlayerId != expectedNextPlayerId {
			t.Errorf("Expected next player after %d to be %d, got %d", i, expectedNextPlayerId, nextPlayerId)
		}
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

	play := newPlay(startingHand, minSlice(startingHand.InHand, NumericCompare))

	result := game.PlayHand(play)
	if !result.Success {
		t.Fatal("Expected play to succeed, but it failed.")
	}

	if result.Round != 1 {
		t.Fatalf("Round number mismatch. Expected: 1, actual: %d.", result.Round)
	}

	if !game.isNextTo(result.NextPlayerId, startingHand.Id) {
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
	game.Init()

	// Attempt to play a card from the left player's hand, which should fail
	startingHand := game.CurrentHand()
	handToTheLeft := game.leftOf(startingHand.Id)
	play := newPlay(&game.Hands[handToTheLeft], minSlice(startingHand.InHand, NumericCompare))

	result := game.PlayHand(play)
	if result.Success {
		t.Fatal("Expected play to fail, but it succeeded.")
	}
	if startingHand.Id != game.CurrentHand().Id {
		t.Fatalf("Current hand should not have changed. Expected: %d, actual: %d.", startingHand.Id, game.CurrentHand().Id)
	}

	// Attempt to play a card not from the current hand, which should also fail
	play = newPlay(&startingHand, minSlice(game.Hands[handToTheLeft].InHand, NumericCompare))
	result = game.PlayHand(play)
	if result.Success {
		t.Fatal("Expected play to fail, but it succeeded.")
	}
	if startingHand.Id != game.CurrentHand().Id {
		t.Fatalf("Current hand should not have changed. Expected: %d, actual: %d.", startingHand.Id, game.CurrentHand().Id)
	}
}

// Test playing a Ten card should empty the in-play pile.
func TestPlayHand10(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	game.Init()

	startingHand := &game.Hands[game.currentPlayerId]
	tenCard := Card{Rank: Ten, Suit: Spade} // Assuming Ten is a valid rank and Spades is a valid suit

	// Add a Ten card to the starting hand for testing
	startingHand.InHand = append(startingHand.InHand, tenCard)

	// Simulate some cards in play
	game.InPlayPile.Cards = game.DrawPile.Cards[:5]
	game.DrawPile.Cards = game.DrawPile.Cards[5:]

	play := newPlay(startingHand, tenCard)

	result := game.PlayHand(play)
	if !result.Success {
		t.Fatal("Expected play with Ten to succeed, but it failed.")
	}

	if len(game.InPlayPile.Cards) != 0 {
		t.Fatal("In-play pile should be empty after playing a Ten.")
	}
}

// Test playing an eight card should appear transparent to the next player.
func TestPlayHand8(t *testing.T) {
	numOfPlayers := 4
	game := NewGame(numOfPlayers)
	game.Init()

	startingHand := &game.Hands[game.currentPlayerId]
	eightCard := Card{Rank: Eight, Suit: Spade} // Assuming Eight is a valid rank and Spades is a valid suit

	// Add an Eight card to the starting hand for testing
	startingHand.InHand = append(startingHand.InHand, eightCard)

	// Add a card to the in-play pile to simulate a game state
	game.InPlayPile.Cards = append(game.InPlayPile.Cards, Card{Rank: Two, Suit: Heart})

	// Ensure the next player has a higher card to play
	expectedNextPlayerId := game.nextPlayerId()
	nextPlayerHandCard := Card{Rank: Three, Suit: Diamond}
	game.Hands[expectedNextPlayerId].InHand = append(game.Hands[expectedNextPlayerId].InHand, nextPlayerHandCard)

	play := newPlay(startingHand, eightCard)

	result := game.PlayHand(play)
	if !result.Success {
		t.Fatal("Expected play with Eight to succeed, but it failed.")
	}

	nextPlay := newPlay(&game.Hands[expectedNextPlayerId], nextPlayerHandCard)
	nextResult := game.PlayHand(nextPlay)
	if !nextResult.Success {
		t.Fatal("Expected next player to play successfully after Eight, but it failed.")
	}
}
