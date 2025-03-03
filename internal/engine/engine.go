package engine

import (
	"fmt"
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

func (game *Game) PlayHand(play Play) PlayResult {
	// Check the correct player played the turn
	if play.Hand.Id != game.currentPlayerId {
		return PlayResult{
			Round:        game.round,
			Success:      false,
			Status:       Play_WrongPlayer,
			NextPlayerId: game.currentPlayerId,
		}
	}

	// Check if the card is in the player's hand
	status := play.Hand.removeCard(play.Card)
	if status != Success {
		return PlayResult{
			Round:        game.round,
			Success:      false,
			Status:       status,
			NextPlayerId: game.currentPlayerId,
		}
	}

	// Check if the card is higher than the top of the in play pile
	if len(game.InPlayPile.Cards) > 0 {
		topCard := game.InPlayPile.Cards[len(game.InPlayPile.Cards)-1]
		if game.compare(play.Card, topCard) < 0 {
			return PlayResult{
				Round:        game.round,
				Success:      false,
				Status:       Play_CardTooLow,
				NextPlayerId: game.currentPlayerId,
			}
		}
	}

	game.InPlayPile.AddCard(play.Card)

	return game.concludePlay(play)
}

func (game *Game) concludePlay(play Play) PlayResult {
	game.round++
	game.currentPlayerId = game.nextPlayerId()

	return PlayResult{
		Round:        game.round,
		Success:      true,
		Status:       Success,
		NextPlayerId: game.currentPlayerId,
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

func (game *Game) nextPlayerId() int {
	return (game.currentPlayerId + game.direction) % len(game.Hands)
}

func (game *Game) nextTo(playerAId int, playerBId int) bool {
	return game.leftOf(playerAId) == playerBId || game.rightOf(playerAId) == playerBId
}
