package engine

type Play struct {
	Hand *Hand
	Card Card
}

type Result struct {
	RoundNumber  int
	Success      bool
	NextPlayerId int
}
