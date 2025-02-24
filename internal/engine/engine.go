package engine

import (
	"fmt"
	"slices"
)

type Game struct {
	DrawPile        *Deck
	DiscardPile     *Deck
	InPlayPile      *Deck
	Hands           []Hand
	comparator      CardComparator
	round           int
	currentPlayerId int
	direction       int
}

const NotStartedPlayerId int = -1
const EndedPlayerId int = -2
const ErrorPlayerId int = -3

func (game *Game) CurrentHand() Hand {
	if game.currentPlayerId == NotStartedPlayerId {
		game.init()
	}
	return game.Hands[game.currentPlayerId]
}

func NewGame(numOfPlayers int) *Game {
	deck := NewDeck()
	hands := make([]Hand, 0, numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		hands = append(hands, Hand{
			Id:       i,
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
		round:           0,
		currentPlayerId: NotStartedPlayerId,
		direction:       1,
		comparator:      BasicComparator,
	}
}

func (game *Game) PlayHand(play Play) Result {
	// Check the correct player played the turn
	if play.Hand.Id != game.currentPlayerId {
		return Result{
			Round:        game.round,
			Success:      false,
			Status:       WrongPlayer,
			NextPlayerId: game.currentPlayerId,
		}
	}

	// Check if the card is in the player's hand
	if len(play.Hand.InHand) != 0 && !slices.Contains(play.Hand.InHand, play.Card) {
		return Result{
			Round:        game.round,
			Success:      false,
			Status:       CardNotInHand,
			NextPlayerId: game.currentPlayerId,
		}
	}

	if len(play.Hand.FaceUp) != 0 && !slices.Contains(play.Hand.FaceUp, play.Card) {
		return Result{
			Round:        game.round,
			Success:      false,
			Status:       CardNotFaceUp,
			NextPlayerId: game.currentPlayerId,
		}
	}

	if len(play.Hand.FaceDown) != 0 && !slices.Contains(play.Hand.FaceDown, play.Card) {
		return Result{
			Round:        game.round,
			Success:      false,
			Status:       CardNotFaceDown,
			NextPlayerId: game.currentPlayerId,
		}
	}

	play.Hand.RemoveCard(play.Card)
	game.InPlayPile.AddCard(play.Card)
	game.round++

	return Result{
		Round:        1,
		Success:      true,
		NextPlayerId: game.leftOf(game.currentPlayerId),
	}
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

func (game *Game) leftOf(playerId int) int {
	return (playerId - 1) % len(game.Hands)
}

func (game *Game) rightOf(playerId int) int {
	return (playerId + 1) % len(game.Hands)
}

func (game *Game) nextTo(playerAId int, playerBId int) bool {
	return game.leftOf(playerAId) == playerBId || game.rightOf(playerAId) == playerBId
}
