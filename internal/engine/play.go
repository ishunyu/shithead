package engine

type Play struct {
	Hand *Hand
	Card Card
}

type Status int

const (
	Success          Status = 0
	Error            Status = 1
	Play_WrongPlayer Status = 101
	Play_CardTooLow  Status = 102
	Hand_NotFound    Status = 201
	Hand_NotInHand   Status = 202
	Hand_NotFaceUp   Status = 203
	Hand_NotFaceDown Status = 204
)

type PlayResult struct {
	Round        int
	Success      bool
	Status       Status
	NextPlayerId int
}
