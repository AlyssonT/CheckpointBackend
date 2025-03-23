package controllers

import (
	"net/http"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/handlers"
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

	createdName, err := uc.handlers.RegisterUser(&request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusCreated,
		Message:    "User created succesfully.",
		Data:       createdName,
	}
	ctx.JSON(response.StatusCode, response)
}
