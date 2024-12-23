package engine

type Suit uint8

type Card struct {
	Suit   Suit
	Number uint8
}

type Deck struct {
	Cards []Card
}

type Hand struct {
	Id       uint8
	InHand   []Card
	FaceDown []Card
}

type Game struct {
	deck  *Deck
	hands []Hand
}
