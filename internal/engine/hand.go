package engine

import "slices"

type Hand struct {
	Id       int
	InHand   []Card
	FaceUp   []Card
	FaceDown []Card
}

type handResult int

func (hand *Hand) dealFaceDown(card Card) {
	hand.FaceDown = append(hand.FaceDown, card)
}

func (hand *Hand) dealFaceUp(card Card) {
	hand.FaceUp = append(hand.FaceUp, card)
}

func (hand *Hand) dealInHand(card Card) {
	hand.InHand = append(hand.InHand, card)
}

func (hand *Hand) removeCard(card Card) Status {
	if len(hand.InHand) != 0 {
		if !slices.Contains(hand.InHand, card) {
			return Hand_NotInHand
		}
		hand.InHand = slices.DeleteFunc(hand.InHand, func(c Card) bool {
			return c == card
		})
		return Success
	}

	if len(hand.FaceUp) != 0 {
		if !slices.Contains(hand.FaceUp, card) {
			return Hand_NotFaceUp
		}
		hand.FaceUp = slices.DeleteFunc(hand.FaceUp, func(c Card) bool {
			return c == card
		})
		return Success
	}

	if len(hand.FaceDown) != 0 {
		if !slices.Contains(hand.FaceDown, card) {
			return Hand_NotFaceDown
		}
		hand.FaceDown = slices.DeleteFunc(hand.FaceDown, func(c Card) bool {
			return c == card
		})
		return Success
	}

	return Hand_NotFound
}
