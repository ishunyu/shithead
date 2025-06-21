package engine

import (
	"fmt"
)

type Game struct {
	DrawPile        *Deck
	InPlayPile      *Deck
	DiscardPile     *Deck
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
		game.Init()
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
		InPlayPile:      &Deck{Cards: make([]Card, 0)},
		DiscardPile:     &Deck{Cards: make([]Card, 0)},
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
		if game.compareCards(play.Card, topCard) < 0 {
			return PlayResult{
				Round:        game.round,
				Success:      false,
				Status:       Play_CardTooLow,
				NextPlayerId: game.currentPlayerId,
			}
		}
	}

	game.InPlayPile.AddCard(play.Card)
	drawCard, error := game.DrawPile.DrawCard()
	if error == nil {
		play.Hand.InHand = append(play.Hand.InHand, drawCard)
	}

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
			card, err := deck.DrawCard()
			if err == nil {
				acceptCard(hand, card)
			}
		}
	}
}

func (game *Game) Init() {
	// Find player with lowest card
	minCard := minSlice(game.Hands[0].InHand, NumericCompare)
	startingPlayerId := 0
	for i := 1; i < len(game.Hands); i++ {
		hand := game.Hands[i]
		minCardInHand := minSlice(hand.InHand, NumericCompare)
		if game.compareCards(minCardInHand, minCard) < 0 {
			startingPlayerId = i
			minCard = minCardInHand
		}
	}
	game.currentPlayerId = startingPlayerId
}
