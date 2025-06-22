package engine

import "testing"

func TestCardComparator(t *testing.T) {
	// Create a custom comparator that compares cards by rank only
	customComparator := newGameComparator(func(a, b Card) (int, comparatorState) {
		if a.Rank < b.Rank {
			return -1, _terminate
		} else if a.Rank > b.Rank {
			return 1, _terminate
		}
		return 0, _continue
	})

	// Test the custom comparator with cards of the same suit
	for i := 0; i < len(clubs)-1; i++ {
		if customComparator.Compare(clubs[i], clubs[i+1]) >= 0 {
			t.Errorf("Expected %v to be less than %v", clubs[i], clubs[i+1])
		}
	}
}

func TestGameCardComparator(t *testing.T) {
}