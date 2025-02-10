package engine

type Hand struct {
	Id       int
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
