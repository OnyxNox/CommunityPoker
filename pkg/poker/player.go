package poker

import "github.com/google/uuid"

// Represents a poker player.
type Player struct {
	// Poker player identifier.
	Id uuid.UUID `json:"id"`
}
