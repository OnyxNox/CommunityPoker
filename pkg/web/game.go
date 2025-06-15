package web

import (
	"log"

	"example.com/community_poker/pkg/poker"

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

	deck []poker.Card

	// Flag identifying the poker game has been started.
	isStarted bool
}

// Adds the player to the poker game if the maximum player count threshold hasn't been reached,
// otherwise noop.
func (game *game) tryAddPlayer(player player) bool {
	if len(game.Players) >= game.MaxPlayerCount {
		log.Printf("Failed to add Player %d to Game %d; maximum player count reached", player.Id, game.Id)

		return false
	}

	game.Players = append(game.Players, player)

	log.Printf("Added Player %d to Game %d", player.Id, game.Id)

	return true
}

// Starts the poker game if the minimum player count threshold has been reached, otherwise noop.
func (game *game) tryStart() bool {
	if game.isStarted {
		log.Printf("Failed to start Game %d; has already been started", game.Id)

		return false
	} else if len(game.Players) < game.MinPlayerCount {
		log.Printf("Failed to start Game %d; not have enough players", game.Id)

		return false
	}

	game.deck = poker.NewDeck()

	game.isStarted = true

	log.Printf("Started Game %d", game.Id)

	return true
}
