package engine

type CardCompare func(a, b Card) int

// NumericCompare is a CardCompare that compares cards by their numeric value.
func NumericCompare(a, b Card) int {
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

type comparatorState int

const (
	_terminate comparatorState = 0
	_continue  comparatorState = 1
)

type cardComparatorFunc func(a, b Card) (int, comparatorState)

type CardComparatorImpl struct {
	compareFunc cardComparatorFunc
	next        *CardComparator
}

func (comparator CardComparatorImpl) Compare(a, b Card) int {
	t, c := comparator.compareFunc(a, b)
	if c == _terminate {
		return t
	}
	return (*comparator.next).Compare(a, b)
}

var BasicComparator CardComparator = CardComparatorImpl{
	compareFunc: func(a, b Card) (int, comparatorState) {
		return NumericCompare(a, b), _terminate
	},
	next: nil,
}

func min(a, b Card, comparatorFunc CardCompare) Card {
	if comparatorFunc(a, b) <= 0 {
		return a
	}
	return b
}

func minSlice(cards []Card, comparatorFunc CardCompare) Card {
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
