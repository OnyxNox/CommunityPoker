package web

import (
	"log"

	"github.com/google/uuid"
)

// Represents a poker game.
type game struct {
	// Poker game identifier.
	Id uuid.UUID `json:"id"`

	// Maximum number of players allowed in the poker game.
	MaxPlayerCount int `json:"maxPlayerCount"`

	// Minimum number of players in order for the game to start.
	MinPlayerCount int `json:"minPlayerCount"`

	// Collection of poker game players.
	Players []player `json:"players"`

	// Flag identifying the poker game has been started.
	isStarted bool
}

// Adds the player to the poker game if the maximum player count threshold hasn't been reached,
// otherwise noop.
func (game *game) tryAddPlayer(player player) bool {
	if len(game.Players) < game.MaxPlayerCount {
		game.Players = append(game.Players, player)

		return true
	}

	return false
}

// Starts the poker game if the minimum player count threshold has been reached, otherwise noop.
func (game *game) tryStart() bool {
	if !game.isStarted && len(game.Players) >= game.MinPlayerCount {
		log.Println("Starting new game server")

		game.isStarted = true

		return true
	}

	return false
}
