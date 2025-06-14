package web

import "github.com/google/uuid"

// Represents a poker game.
type game struct {
	// Poker game identifier.
	Id uuid.UUID `json:"id"`

	// Collection of poker game players.
	Players []player `json:"players"`
}
