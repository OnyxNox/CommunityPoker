package main

import (
	"flag"
	"fmt"
	"net/http"

	"example.com/community_poker/pkg/poker"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Poker games collection.
var games = []poker.Game{}

// Server application entry point.
func main() {
	port := flag.Uint("port", 8080, "Port to listen on.")
	flag.Parse()

	router := gin.New()

	api := router.Group("/api")
	{
		games := api.Group("/games")
		{
			games.GET("/:gameId", getGame)
			games.GET("/", getGames)
			games.POST("/:gameId/join", joinGame)
			games.POST("/new", newGame)
		}
	}

	router.Run(fmt.Sprintf(":%d", *port))
}

// Returns poker game if found, otherwise nil.
func findGame(gameId string) *poker.Game {
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

// Responds with poker games collection.
func getGames(context *gin.Context) {
	context.JSON(http.StatusOK, games)
}

// Joins player to a poker game.
func joinGame(context *gin.Context) {
	gameId := context.Param("gameId")

	game := findGame(gameId)
	if game == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Failed to find game %s", gameId)})

		return
	}

	player := poker.Player{
		Id: uuid.New(),
	}

	if game.TryAddPlayer(player) {
		game.TryStart()
	}

	context.Status(http.StatusOK)
}

// Creates new poker game and responds with its details.
func newGame(context *gin.Context) {
	game := poker.NewGame(2, 8)

	games = append(games, *game)

	context.JSON(http.StatusCreated, game)
}
