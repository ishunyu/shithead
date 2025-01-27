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

type Hand struct {
	Id       uint8
	InHand   []Card
	FaceUp   []Card
	FaceDown []Card
}

func (hand *Hand) dealFaceDown(card Card) {
	hand.FaceDown = append(hand.FaceDown, card)
}

func (hand *Hand) dealFaceUp(card Card) {
	hand.FaceUp = append(hand.FaceUp, card)
}

func (hand *Hand) dealInHand(card Card) {
	hand.InHand = append(hand.InHand, card)
}

type CardCompareFunc func(a, b Card) int

func BasicCompare(a, b Card) int {
	if a == b {
		return 0
	}

	rankDiff := int(a.Rank) - int(b.Rank)
	if rankDiff != 0 {
		return rankDiff
	}

	return int(a.Suit) - int(b.Suit)
}

type CardComparator interface {
	Compare(a, b Card) int
}

type CardComparatorImplState int

const (
	_terminate = 0
	_continue  = 1
)

type CardComparatorImpl struct {
	f     func(a, b Card) (int, CardComparatorImplState)
	inner *CardComparator
}

func (comparator CardComparatorImpl) Compare(a, b Card) int {
	t, c := comparator.f(a, b)
	if c == _terminate {
		return t
	}
	return (*comparator.inner).Compare(a, b)
}

var BasicComparator CardComparator = CardComparatorImpl{
	f: func(a, b Card) (int, CardComparatorImplState) {
		return BasicCompare(a, b), _terminate
	},
	inner: nil,
}

type Game struct {
	Deck            *Deck
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
		Deck:            deck,
		Hands:           hands,
		currentPlayerId: InitialPlayerId,
		comparator:      BasicComparator,
	}
}

func (game *Game) String() string {
	s := ""
	s += "Game:\n"
	s += "Deck: " + game.Deck.String() + "\n"
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
	minCard := minSlice(game.Hands[0].InHand, BasicCompare)
	startingPlayerId := 0
	for i := 1; i < len(game.Hands); i++ {
		hand := game.Hands[i]
		minCardInHand := minSlice(hand.InHand, BasicCompare)
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

func minSlice(cards []Card, comparatorFunc CardCompareFunc) Card {
	numOfCards := len(cards)
	if numOfCards == 0 {
		panic("cards is empty")
	}

	minCard := cards[0]
	for i := 1; i < numOfCards; i++ {
		minCard = min(minCard, cards[i], comparatorFunc)
	}
	return minCard
}

func min(a, b Card, comparatorFunc CardCompareFunc) Card {
	if comparatorFunc(a, b) <= 0 {
		return a
	}
	return b
}
