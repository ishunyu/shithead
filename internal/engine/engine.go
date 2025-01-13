package engine

import (
	"fmt"
	"math"
	"math/rand/v2"
)

type Suit uint8

const (
	Club       Suit = 0
	Diamond    Suit = 1
	Heart      Suit = 2
	Spade      Suit = 3
	JokerSmall Suit = math.MaxUint8 - 1
	JokerLarge Suit = math.MaxUint8
)

type Rank uint8

const (
	Ace   Rank = 1
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Joker Rank = math.MaxUint8
)

type Card struct {
	Suit Suit
	Rank Rank
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
	numOfCards := 54
	suits := []Suit{0, 1, 2, 3}
	ranks := []Rank{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	standardDeck := make([]Card, 0, numOfCards)

	for _, suit := range suits {
		for _, rank := range ranks {
			standardDeck = append(standardDeck, Card{suit, rank})
		}
	}

	standardDeck = append(standardDeck, Card{JokerSmall, Joker})
	standardDeck = append(standardDeck, Card{JokerLarge, Joker})

	deck := make([]Card, 0, numOfCards)
	shuffle := rand.Perm(numOfCards)

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
	switch suit {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"
	case JokerSmall:
		return "JokerSmall"
	case JokerLarge:
		return "JokerLarge"
	default:
		return "Unknown"
	}
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
