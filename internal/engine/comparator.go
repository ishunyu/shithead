package engine

//
// Compare related functions
//

func (game *Game) compareCards(a Card, b Card) int {
	return game.comparator.Compare(a, b)
}

func newGameComparator(compareFunc cardComparatorFunc) CardComparator {
	return CardComparatorImpl{
		compareFunc: compareFunc,
		next:        &BasicComparator,
	}
}
