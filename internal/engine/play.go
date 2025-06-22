package engine

import "fmt"

type Play struct {
	Hand       *Hand
	Card       Card // Card that the play represents
	ActualCard Card // Actual card that was played
}

type PlayResult struct {
	Round        int
	Success      bool
	Status       Status
	NextPlayerId int
}

func newPlay(hand *Hand, card Card) Play {
	return Play{
		Hand:       hand,
		Card:       card,
		ActualCard: card, // In this context, ActualCard is the same as Card
	}
}

func newPlayWithJoker(hand *Hand, card Card, joker Card) (Play, error) {
	if joker.Rank != Joker {
		return Play{}, fmt.Errorf("joker card must have a rank of Joker: %s", joker)
	}

	if !(joker.Suit == JokerSmall || joker.Suit == JokerLarge) {
		return Play{}, fmt.Errorf("joker card has the wrong suit: %s", joker)
	}

	return Play{
		Hand:       hand,
		Card:       card,
		ActualCard: joker,
	}, nil
}