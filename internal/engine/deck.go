package engine

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

type Suit uint8

const (
	Club       Suit = 0
	Diamond    Suit = 1
	Heart      Suit = 2
	Spade      Suit = 3
	JokerSmall Suit = 4
	JokerLarge Suit = 5
)

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
	Joker Rank = 14
)

type Card struct {
	Suit Suit
	Rank Rank
}

var StandardDeck []Card = newStandardDeck()

func newStandardDeck() []Card {
	suits := []Suit{Club, Diamond, Heart, Spade}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	cards := make([]Card, 0, 54)
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{suit, rank})
		}
	}

	cards = append(cards, Card{JokerSmall, Joker})
	cards = append(cards, Card{JokerLarge, Joker})

	return cards
}

func validate(card Card) bool {
	return slices.Contains[[]Card](StandardDeck, card)
}

func (card Card) String() string {
	return fmt.Sprintf("(%s, %d)", card.Suit, card.Rank)
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	numOfCards := 54
	suits := []Suit{Club, Diamond, Heart, Spade}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
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

func (deck *Deck) DrawCard() Card {
	card, remaining := deck.Cards[0], deck.Cards[1:]
	deck.Cards = remaining
	return card
}

func (deck *Deck) String() string {
	s := ""
	for i, card := range deck.Cards {
		if i == len(deck.Cards)-1 {
			s += card.String()
			break
		}
		s += card.String() + ", "
	}
	return s
}
