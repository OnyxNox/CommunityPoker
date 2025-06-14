package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Represents a poker game instance.
type game struct {
	// Game instance identifier.
	Id uuid.UUID `json:"id"`
}

// Collection of poker game instances.
var games = []game{}

// Start Community Poker web server.
func StartServer() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.POST("/games/new", newGame)

	router.Run("localhost:8080")
}

// Responds with collection of poker game instances.
func getGames(context *gin.Context) {
	context.JSON(http.StatusOK, games)
}

// Creates new poker game instance and responds with its details.
func newGame(context *gin.Context) {
	newGame := game{
		Id: uuid.New(),
	}

	games = append(games, newGame)

	context.JSON(http.StatusCreated, newGame)
}
