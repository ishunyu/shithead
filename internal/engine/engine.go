package engine

import (
	"fmt"
	"math/rand/v2"
)

type Suit uint8

const (
	Club    Suit = 0
	Diamond Suit = 1
	Heart   Suit = 2
	Spade   Suit = 3
)

type Card struct {
	Suit Suit
	Rank uint8
}

type Deck struct {
	Cards []Card
}

type Hand struct {
	Id       uint8
	InHand   []Card
	FaceUp   []Card
	FaceDown []Card
}

type Game struct {
	Deck  *Deck
	Hands []Hand
}

func NewDeck() *Deck {
	suits := []Suit{Club, Diamond, Heart, Spade}
	ranks := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	standardDeck := make([]Card, 0, 52)

	for _, suit := range suits {
		for _, rank := range ranks {
			standardDeck = append(standardDeck, Card{suit, rank})
		}
	}

	deck := make([]Card, 0, 52)
	shuffle := rand.Perm(52)

	for _, i := range shuffle {
		deck = append(deck, standardDeck[i])
	}

	return &Deck{deck}
}

func (deck *Deck) String() string {
	s := ""
	for _, card := range deck.Cards {
		s = s + fmt.Sprintf("%s\n", card)
	}

	return s
}

func (suit Suit) String() string {
	if suit == Club {
		return "Club"
	}

	if suit == Diamond {
		return "Diamond"
	}

	if suit == Heart {
		return "Heart"
	}

	if suit == Spade {
		return "Spade"
	}

	return "Unknown"
}

func (card Card) String() string {
	return fmt.Sprintf("(%s, %d)", card.Suit, card.Rank)
}

func NewGame(numOfPlayers int) *Game {
	deck := NewDeck()
	hands := make([]Hand, 0, numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		hands = append(hands, Hand{
			Id:       uint8(i),
			InHand:   make([]Card, 0),
			FaceUp:   make([]Card, 0),
			FaceDown: make([]Card, 0),
		})
	}
	return &Game{
		Deck:  deck,
		Hands: hands,
	}
}
