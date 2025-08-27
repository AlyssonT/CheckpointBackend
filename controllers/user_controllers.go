package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	handlers *handlers.UserHandlers
}

func NewUserControllers(handlers *handlers.Handlers) *UserController {
	return &UserController{
		handlers: handlers.UserHandlers,
	}
}

// @Summary		User data
// @Description	Get user data from cookie token
// @ID				cookie-token-user-data
// @Produce		json
// @Router			/me [get]
// @Tags			User
// @Security		cookieAuth
// @Success		200
// @Failure		401
// @Failure		500
func (uc *UserController) Me(ctx *gin.Context) {
	userData, exists := ctx.Get("userData")
	if !exists {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	parsedUserData, ok := userData.(interfaces.UserClaims)
	if !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       parsedUserData,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Register user
// @Description	Register user in the database
// @ID				register-user
// @Produce		json
// @Param			request	body	communication.RegisterUserRequest	true	"User data"
// @Router			/users [post]
// @Tags			Authentication
// @Success		201
// @Failure		400
// @Failure		409
// @Failure		500
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var request communication.RegisterUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		messages := exceptions.CreateValidationErrorMessages(err)
		response := communication.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation error",
			Data:       messages,
		}
		ctx.JSON(response.StatusCode, response)
		return
	}

	token, err := uc.handlers.RegisterUser(&request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	ctx.SetCookie("auth_token", token, int(time.Hour.Seconds())*24, "/", "", true, true)

	response := communication.ResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "User created succesfully.",
		Data:       nil,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Update user details
// @Description	Update user profile details like bio, avatar etc.
// @Tags			User
// @Accept			multipart/form-data
// @Produce		json
// @Security		cookieAuth
// @Router			/user/profile [put]
// @Param			bio		formData	string	false	"User Bio"
// @Param			avatar	formData	file	false	"User Avatar"
// @Success		200
// @Failure		400
// @Failure		500
func (uc *UserController) UpdateUserProfileDetails(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	bio := ctx.PostForm("bio")
	file, _, err := ctx.Request.FormFile("avatar")

	if err != nil {
		if !strings.Contains(err.Error(), "http: no such file") {
			response := exceptions.ErrorHandler(exceptions.ErrorInvalidAvatarData)
			ctx.JSON(response.StatusCode, response)
			return
		}
	}

	userProfileDetails := communication.UserProfileDetails{
		UserID:     parsedID,
		Bio:        bio,
		AvatarData: file,
	}

	uc.handlers.UpdateUserProfileDetails(&userProfileDetails)

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "Profile details updated succesfully.",
		Data:       true,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Get user profile
// @Description	Get user profile
// @Tags			User
// @Produce		json
// @Security		cookieAuth
// @Router			/user/profile [get]
// @Success		200
// @Failure		401
// @Failure		500
func (uc *UserController) GetUserProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	userProfile, err := uc.handlers.GetUserProfile(parsedID)

	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       userProfile,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Add game to user
// @Description	Add a game to the user's collection
// @Tags			User
// @Produce		json
// @Security		cookieAuth
// @Router			/user/games [post]
// @Param			request	body	communication.AddGameToUserRequest	true	"Game data"
// @Success		200
// @Failure		401
// @Failure		500
func (uc *UserController) AddGameToUser(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	var request communication.AddGameToUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		messages := exceptions.CreateValidationErrorMessages(err)
		response := communication.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation error",
			Data:       messages,
		}
		ctx.JSON(response.StatusCode, response)
		return
	}

	err = uc.handlers.AddGameToUser(parsedID, &request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "Game added to user successfully.",
		Data:       nil,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Update user game
// @Description	Update user game
// @Tags			User
// @Produce		json
// @Security		cookieAuth
// @Router			/user/games/{gameId} [put]
// @Param			gameId	path	string									true	"Game ID to update"
// @Param			request	body	communication.UpdateGameToUserRequest	true	"Game data"
// @Success		200
// @Failure		401
// @Failure		500
func (uc *UserController) UpdateGameToUser(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	game_id := ctx.Param("gameId")
	game_idUint, err := strconv.ParseUint(game_id, 10, 32)

	if err != nil {
		response := communication.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid game ID",
			Data:       nil,
		}
		ctx.JSON(response.StatusCode, response)
		return
	}

	var request communication.UpdateGameToUserRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		messages := exceptions.CreateValidationErrorMessages(err)
		response := communication.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "Validation error",
			Data:       messages,
		}
		ctx.JSON(response.StatusCode, response)
		return
	}

	err = uc.handlers.UpdateGameToUser(parsedID, uint(game_idUint), &request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "Game updated successfully.",
		Data:       nil,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Delete user game
// @Description	Delete user game
// @Tags			User
// @Produce		json
// @Security		cookieAuth
// @Router			/user/games/{gameId} [delete]
// @Param			gameId	path	string	true	"Game ID to delete"
// @Success		200
// @Failure		401
// @Failure		500
func (uc *UserController) DeleteGameToUser(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	game_id := ctx.Param("gameId")
	game_idUint, err := strconv.ParseUint(game_id, 10, 32)

	if err != nil {
		response := communication.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid game ID",
			Data:       nil,
		}
		ctx.JSON(response.StatusCode, response)
		return
	}

	request := communication.DeleteGameToUserRequest{
		Game_id: uint(game_idUint),
	}

	err = uc.handlers.DeleteGameToUser(parsedID, &request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "Game deleted successfully.",
		Data:       nil,
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Get user games
// @Description	Get user games
// @Tags			User
// @Produce		json
// @Security		cookieAuth
// @Router			/user/games [get]
// @Success		200
// @Failure		401
// @Failure		500
func (uc *UserController) GetUserGames(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	games, totalItems, err := uc.handlers.GetUserGames(parsedID)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := &communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       communication.UserGamesResponse{Games: games, TotalItems: totalItems},
	}
	ctx.JSON(response.StatusCode, response)
}

// @Summary		Get user game
// @Description	Get user game by id
// @Tags			User
// @Produce		json
// @Security		cookieAuth
// @Router			/user/games/{gameId} [get]
// @Param			gameId	path	string	true	"Game ID"
// @Success		200
// @Failure		401
// @Failure		404
// @Failure		500
func (uc *UserController) GetUserGameById(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	parsedID, ok := userID.(uint)

	if !exists || !ok {
		response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
		ctx.JSON(response.StatusCode, response)
		return
	}

	game_id := ctx.Param("gameId")
	parsedGameId, err := strconv.ParseUint(game_id, 10, 32)

	if err != nil {
		response := communication.ResponseDTO{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid game ID",
			Data:       nil,
		}
		ctx.JSON(response.StatusCode, response)
		return
	}

	game, err := uc.handlers.GetUserGameById(parsedID, uint(parsedGameId))
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
