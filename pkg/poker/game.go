package poker

import (
	"log"

	"github.com/google/uuid"
)

// Represents a poker game's status.
type Status uint8

const (
	Created Status = iota
	Starting
	Started
	Flopping
	Flopped
	Turning
	Turned
	Flipping
	Flipped
	Ending
	Ended
)

type Game struct {
	Id uuid.UUID `json:"id"`

	deck []Card

	maxPlayerCount int

	minPlayerCount int

	players []Player

	status Status
}

func NewGame(minPlayerCount, maxPlayerCount int) *Game {
	deck := NewDeck()

	Shuffle(deck)

	return &Game{
		Id:             uuid.New(),
		deck:           deck,
		maxPlayerCount: maxPlayerCount,
		minPlayerCount: minPlayerCount,
		players:        []Player{},
		status:         Created,
	}
}

// Adds the player to the poker game if the maximum player count threshold hasn't been reached,
// otherwise noop.
func (game *Game) TryAddPlayer(player Player) bool {
	if len(game.players) >= game.maxPlayerCount {
		log.Printf("Failed to add Player %v to Game %v; maximum player count (%d) reached", player.Id, game.Id, game.maxPlayerCount)

		return false
	}

	game.players = append(game.players, player)

	log.Printf("Added Player %v to Game %v", player.Id, game.Id)

	return true
}

// Starts the poker game if the minimum player count threshold has been reached, otherwise noop.
func (game *Game) TryStart() bool {
	if game.status >= Started {
		log.Printf("Failed to start Game %v; has already been started", game.Id)

		return false
	} else if len(game.players) < game.minPlayerCount {
		log.Printf("Failed to start Game %v; not enough players", game.Id)

		return false
	}

	game.status = Started

	log.Printf("Started Game %v", game.Id)

	return true
}
