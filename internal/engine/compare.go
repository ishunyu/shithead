package engine

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
