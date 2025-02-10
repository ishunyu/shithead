package engine

import (
	"fmt"
)

type Game struct {
	DrawPile        *Deck
	DiscardPile     *Deck
	InPlayPile      *Deck
	Hands           []Hand
	currentPlayerId int
	comparator      CardComparator
}

const InitialPlayerId int = -1
const EndedPlayerId int = -2

func (game *Game) CurrentHand() Hand {
	if game.currentPlayerId == InitialPlayerId {
		game.init()
	}
	return game.Hands[game.currentPlayerId]
}

func NewGame(numOfPlayers int) *Game {
	deck := NewDeck()
	hands := make([]Hand, 0, numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		hands = append(hands, Hand{
			Id:       uint8(i),
			InHand:   make([]Card, 0, 3),
			FaceUp:   make([]Card, 0, 3),
			FaceDown: make([]Card, 0, 3),
		})
	}

	// Deal hands
	dealCard(deck, hands, 3, (*Hand).dealFaceDown)
	dealCard(deck, hands, 3, (*Hand).dealFaceUp)
	dealCard(deck, hands, 3, (*Hand).dealInHand)

	return &Game{
		DrawPile:        deck,
		Hands:           hands,
		currentPlayerId: InitialPlayerId,
		comparator:      BasicComparator,
	}
}

func (game *Game) PlayHand(hand *Hand) {

}

func (game *Game) String() string {
	s := ""
	s += "Game:\n"
	s += "Deck: " + game.DrawPile.String() + "\n"
	s += "Hands: [\n"
	for _, hand := range game.Hands {
		s += fmt.Sprintf("  %+v\n", hand)
	}
	s += "]\n"
	s += fmt.Sprintf("currentPlayerId: %v\n", game.currentPlayerId)
	s += fmt.Sprintf("comparator: %v", game.comparator)
	return s
}

func dealCard(deck *Deck, hands []Hand, numOfRounds int, acceptCard func(hand *Hand, card Card)) {
	for r := 0; r < numOfRounds; r++ {
		for i := range hands {
			hand := &hands[i]
			acceptCard(hand, deck.DrawCard())
		}
	}
}

func (game *Game) init() {
	// Find player with lowest card
	minCard := minSlice(game.Hands[0].InHand, NumericCompare)
	startingPlayerId := 0
	for i := 1; i < len(game.Hands); i++ {
		hand := game.Hands[i]
		minCardInHand := minSlice(hand.InHand, NumericCompare)
		if game.compare(minCardInHand, minCard) < 0 {
			startingPlayerId = i
			minCard = minCardInHand
		}
	}
	game.currentPlayerId = startingPlayerId
}

func (game *Game) compare(a Card, b Card) int {
	return game.comparator.Compare(a, b)
}

func newGameComparator(compareFunc cardComparatorFunc) CardComparator {
	return CardComparatorImpl{
		compareFunc: compareFunc,
		next:        &BasicComparator,
	}
}
