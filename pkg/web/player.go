package web

import "github.com/google/uuid"

// Represents a poker player.
type player struct {
	// Poker player identifier.
	Id uuid.UUID `json:"id"`
}
