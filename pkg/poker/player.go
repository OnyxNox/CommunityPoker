package poker

import "github.com/google/uuid"

// Represents a poker player.
type Player struct {
	// Total amount of chips player has.
	Bank uint `json:"bank"`

	// Player identifier.
	PlayerId uuid.UUID `json:"id"`

	// Cards in player's hand.
	hand []Card
}

// Initializes a new player.
func NewPlayer() Player {
	return Player{
		Bank:     500,
		PlayerId: uuid.New(),
		hand:     []Card{},
	}
}
