package engine

import "testing"

var (
	clubs = []Card{
		{Suit: Club, Rank: Ace},
		{Suit: Club, Rank: Two},
		{Suit: Club, Rank: Three},
		{Suit: Club, Rank: Four},
		{Suit: Club, Rank: Five},
		{Suit: Club, Rank: Six},
		{Suit: Club, Rank: Seven},
		{Suit: Club, Rank: Eight},
		{Suit: Club, Rank: Nine},
		{Suit: Club, Rank: Ten},
		{Suit: Club, Rank: Jack},
		{Suit: Club, Rank: Queen},
		{Suit: Club, Rank: King},
	}

	diamonds = []Card{
		{Suit: Diamond, Rank: Ace},
		{Suit: Diamond, Rank: Two},
		{Suit: Diamond, Rank: Three},
		{Suit: Diamond, Rank: Four},
		{Suit: Diamond, Rank: Five},
		{Suit: Diamond, Rank: Six},
		{Suit: Diamond, Rank: Seven},
		{Suit: Diamond, Rank: Eight},
		{Suit: Diamond, Rank: Nine},
		{Suit: Diamond, Rank: Ten},
		{Suit: Diamond, Rank: Jack},
		{Suit: Diamond, Rank: Queen},
		{Suit: Diamond, Rank: King},
	}

	hearts = []Card{
		{Suit: Heart, Rank: Ace},
		{Suit: Heart, Rank: Two},
		{Suit: Heart, Rank: Three},
		{Suit: Heart, Rank: Four},
		{Suit: Heart, Rank: Five},
		{Suit: Heart, Rank: Six},
		{Suit: Heart, Rank: Seven},
		{Suit: Heart, Rank: Eight},
		{Suit: Heart, Rank: Nine},
		{Suit: Heart, Rank: Ten},
		{Suit: Heart, Rank: Jack},
		{Suit: Heart, Rank: Queen},
		{Suit: Heart, Rank: King},
	}

	spades = []Card{
		{Suit: Spade, Rank: Ace},
		{Suit: Spade, Rank: Two},
		{Suit: Spade, Rank: Three},
		{Suit: Spade, Rank: Four},
		{Suit: Spade, Rank: Five},
		{Suit: Spade, Rank: Six},
		{Suit: Spade, Rank: Seven},
		{Suit: Spade, Rank: Eight},
		{Suit: Spade, Rank: Nine},
		{Suit: Spade, Rank: Ten},
		{Suit: Spade, Rank: Jack},
		{Suit: Spade, Rank: Queen},
		{Suit: Spade, Rank: King},
	}

	jokers = []Card{
		{Suit: JokerSmall, Rank: Joker},
		{Suit: JokerLarge, Rank: Joker},
	}
)

func TestRankCompare(t *testing.T) {
	// Function to test comparison of cards within a suit
	testCompare := func(cards []Card) {
		for i := 0; i < len(cards)-1; i++ {
			if BasicComparator.Compare(cards[i], cards[i+1]) >= 0 {
				t.Errorf("Expected %v to be less than %v", cards[i], cards[i+1])
			}
		}
	}

	// Run the compare test for each suit
	testCompare(clubs)
	testCompare(diamonds)
	testCompare(hearts)
	testCompare(spades)
	testCompare(jokers)
}

func TestSuitCompare(t *testing.T) {
	// Function to test comparison of cards across suits
	testCompareAcrossSuits := func(cards1, cards2 []Card) {
		for i := 0; i < len(cards1); i++ {
			if BasicComparator.Compare(cards1[i], cards2[i]) >= 0 {
				t.Errorf("Expected %v to be less than %v", cards1[i], cards2[i])
			}
		}
	}

	// Run the compare test across suits
	testCompareAcrossSuits(clubs, diamonds)
	testCompareAcrossSuits(diamonds, hearts)
	testCompareAcrossSuits(hearts, spades)
}

func TestJoker(t *testing.T) {
	// Function to test comparison of jokers with the largest card in each suit
	testCompareWithJoker := func(cards []Card, joker Card) {
		largestCard := cards[len(cards)-1]
		if BasicComparator.Compare(largestCard, joker) >= 0 {
			t.Errorf("Expected %v to be less than %v", largestCard, joker)
		}
	}

	// Run the compare test for each joker with the largest card in each suit
	for _, joker := range jokers {
		testCompareWithJoker(clubs, joker)
		testCompareWithJoker(diamonds, joker)
		testCompareWithJoker(hearts, joker)
		testCompareWithJoker(spades, joker)
	}
}

func TestMin(t *testing.T) {
	// Function to test the min function
	testMin := func(cards []Card, expected Card) {
		result := minSlice(cards, NumericCompare)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}

	// Test for each suit
	testMin(clubs, clubs[0])
	testMin(diamonds, diamonds[0])
	testMin(hearts, hearts[0])
	testMin(spades, spades[0])
	testMin(jokers, jokers[0])
}
