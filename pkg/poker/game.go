package poker

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Represents a poker game's status.
type Status uint8

const (
	Created Status = iota
	Starting
	PreFlop
	Flop
	Turn
	River
	Showdown
	Ended
)

// Converts the enumerator's numeric value to its corresponding string.
func (status Status) String() string {
	return [...]string{"Created", "Starting", "PreFlop", "Flop", "Turn", "River", "Showdown", "Ended"}[status]
}

// Converts the enumerator's numeric value to its corresponding string during JSON encoding.
func (status Status) MarshalJSON() ([]byte, error) {
	return []byte(`"` + status.String() + `"`), nil
}

// Represents a poker game.
type Game struct {
	Blinds [2]uint `json:"blinds"`

	CurrentPlayerId uuid.UUID `json:"currentPlayer"`

	DealerPlayerId uuid.UUID `json:"dealerPlayerId"`

	GameId uuid.UUID `json:"gameId"`

	MaxPlayerCount int `json:"maxPlayerCount"`

	MinPlayerCount int `json:"minPlayerCount"`

	Players []Player `json:"players"`

	Pot uint `json:"pot"`

	Status Status `json:"status"`

	TurnEndTimestamp time.Time `json:"turnEndTimestamp"`

	deck []Card
}

// Initializes a new poker game.
func NewGame(minPlayerCount, maxPlayerCount int) *Game {
	deck := NewDeck()

	Shuffle(deck)

	game := Game{
		Blinds:         [2]uint{30, 90},
		GameId:         uuid.New(),
		MaxPlayerCount: maxPlayerCount,
		MinPlayerCount: minPlayerCount,
		Players:        []Player{},
		Status:         Created,
		deck:           deck,
	}

	log.Printf("Created Game %v", game.GameId)

	return &game
}

// Adds the player to the poker game if the maximum player count threshold hasn't been reached,
// otherwise noop.
func (game *Game) TryAddPlayer(player Player) bool {
	if len(game.Players) >= game.MaxPlayerCount {
		log.Printf("Failed to add Player %v to Game %v; maximum player count (%d) reached", player.PlayerId, game.GameId, game.MaxPlayerCount)

		return false
	}

	game.Players = append(game.Players, player)

	log.Printf("Added Player %v to Game %v", player.PlayerId, game.GameId)

	return true
}

// Starts the poker game if the minimum player count threshold has been reached, otherwise noop.
func (game *Game) TryStart() bool {
	playerCount := len(game.Players)

	if game.Status >= Starting {
		log.Printf("Failed to start Game %v; has already been started", game.GameId)

		return false
	} else if playerCount < game.MinPlayerCount {
		log.Printf("Failed to start Game %v; not enough Players", game.GameId)

		return false
	}

	game.TurnEndTimestamp = time.Now().Add(10 * time.Second).UTC()

	game.Status = Starting

	log.Printf("Starting Game %v", game.GameId)

	go sleepUntilAndRun(game.TurnEndTimestamp, func() {
		playerCount = len(game.Players)

		dealerIndex := rand.Intn(playerCount)

		game.CurrentPlayerId = game.Players[(dealerIndex+1)%playerCount].PlayerId
		game.DealerPlayerId = game.Players[dealerIndex%playerCount].PlayerId

		game.dealCards()

		game.Status = PreFlop

		log.Printf("Started Game %v with %d players", game.GameId, len(game.Players))
	})

	return true
}

// Deals each players private cards.
func (game *Game) dealCards() {
	const cardsPerPlayer = 2

	for range cardsPerPlayer {
		for i := range game.Players {
			if len(game.deck) == 0 {
				return
			}

			game.Players[i].hand = append(game.Players[i].hand, game.deck[0])

			game.deck = game.deck[1:]
		}
	}
}

// Sleeps goroutine until specified start run time, then runs the given function.
func sleepUntilAndRun(startRunTime time.Time, fn func()) {
	duration := time.Until(startRunTime)

	if duration > 0 {
		time.Sleep(duration)
	}

	fn()
}
