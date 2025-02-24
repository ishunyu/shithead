package engine

type Play struct {
	Hand *Hand
	Card Card
}

type Status int

const (
	Error         Status = 0
	Success       Status = 1
	WrongPlayer   Status = 2
	CardNotInHand Status = 3
)

type Result struct {
	Round        int
	Success      bool
	Status       Status
	NextPlayerId int
}
