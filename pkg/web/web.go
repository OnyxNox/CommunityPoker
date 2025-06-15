package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Collection of poker games.
var games = []game{}

// Starts Community Poker web server.
func StartServer() {
	router := gin.Default()
	router.GET("/games/:gameId", getGame)
	router.GET("/games", getGames)
	router.POST("/games/:gameId/join", joinGame)
	router.POST("/games/new", newGame)

	router.Run("localhost:8080")
}

// Returns game if found, otherwise nil.
func findGame(gameId string) *game {
	for i := range games {
		if games[i].Id.String() == gameId {
			return &games[i]
		}
	}

	return nil
}

// Responds with poker game.
func getGame(context *gin.Context) {
	gameId := context.Param("gameId")

	game := findGame(gameId)
	if game == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Failed to find game %s", gameId)})

		return
	}

	context.JSON(http.StatusOK, game)
}

// Responds with collection of poker games.
func getGames(context *gin.Context) {
	context.JSON(http.StatusOK, games)
}

// Joins poker player to a poker game.
func joinGame(context *gin.Context) {
	gameId := context.Param("gameId")

	game := findGame(gameId)
	if game == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Failed to find game %s", gameId)})

		return
	}

	newPlayer := player{
		Id: uuid.New(),
	}

	if game.tryAddPlayer(newPlayer) {
		game.tryStart()
	}

	context.Status(http.StatusOK)
}

// Creates new poker game and responds with its details.
func newGame(context *gin.Context) {
	newGame := game{
		Id:             uuid.New(),
		MinPlayerCount: 2,
		MaxPlayerCount: 8,
		Players:        []player{},
	}

	games = append(games, newGame)

	context.JSON(http.StatusCreated, newGame)
}
