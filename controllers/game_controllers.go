package controllers

import (
	"net/http"
	"strconv"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	handlers *handlers.GameHandlers
}

func NewGameControllers(handlers *handlers.Handlers) *GameController {
	return &GameController{
		handlers: handlers.GameHandlers,
	}
}

// @Summary		Get games
// @Description	Get a list of games
// @ID				get-games
// @Produce		json
// @Param			page		query	int		false	"Page index"
// @Param			pageSize	query	int		false	"Page size"
// @Param			query		query	string	false	"Query for search"
// @Router			/games [get]
// @Tags			Games
// @Security		cookieAuth
// @Success		200
// @Failure		401
// @Failure		500
func (gc *GameController) GetGames(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	request := communication.GetGamesRequest{
		PaginationRequest: communication.PaginationRequest{
			Page:     page,
			PageSize: pageSize,
		},
		Query: ctx.Query("query"),
	}

	games, totalItems, err := gc.handlers.GetGames(&request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := &communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "",
		Data: gin.H{
			"games":      games,
			"totalItems": totalItems,
		},
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Get game by id
// @Description	Get game by id
// @ID				get-game-by-id
// @Produce		json
// @Router			/games/{gameId} [get]
// @Param			gameId	path	string	true	"Game id"
// @Tags			Games
// @Security		cookieAuth
// @Success		200
// @Failure		401
// @Failure		500
func (gc *GameController) GetGameById(ctx *gin.Context) {
	pathId := ctx.Param("gameId")
	gameId, err := strconv.Atoi(pathId)

	if err != nil {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidGameId)
		ctx.JSON(response.StatusCode, response)
		return
	}

	game, err := gc.handlers.GetGameById(gameId)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := &communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       game,
	}
	ctx.JSON(response.StatusCode, response)
}
