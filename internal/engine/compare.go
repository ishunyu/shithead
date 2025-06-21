package engine

// Compare related functions

func (game *Game) compareCards(a Card, b Card) int {
	return game.comparator.Compare(a, b)
}

func newGameComparator(compareFunc cardComparatorFunc) CardComparator {
	return CardComparatorImpl{
		compareFunc: compareFunc,
		next:        &BasicComparator,
	}
}

func (game *Game) leftOf(playerId int) int {
	return (playerId - 1 + len(game.Hands)) % len(game.Hands)
}

func (game *Game) rightOf(playerId int) int {
	return (playerId + 1) % len(game.Hands)
}

func (game *Game) nextPlayerId() int {
	return (game.currentPlayerId + game.direction + len(game.Hands)) % len(game.Hands)
}

func (game *Game) isNextTo(playerAId int, playerBId int) bool {
	return game.leftOf(playerAId) == playerBId || game.rightOf(playerAId) == playerBId
}