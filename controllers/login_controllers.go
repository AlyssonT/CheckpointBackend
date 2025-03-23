package controllers

import (
	"net/http"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	handlers *handlers.LoginHandlers
}

func NewLoginControllers(handlers *handlers.Handlers) *LoginController {
	return &LoginController{
		handlers: handlers.LoginHandlers,
	}
}

// @Summary		Login
// @Description	Login user
// @ID				login
// @Produce		json
// @Param			request	body	communication.LoginRequest	true	"User credentials"
// @Router			/login [post]
// @Tags			Authentication
// @Success		200
// @Failure		401
// @Failure		500
func (lc *LoginController) Login(ctx *gin.Context) {
	var request communication.LoginRequest
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

	token, err := lc.handlers.Login(&request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := &communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "logged in",
		Data:       token,
	}
	ctx.JSON(response.StatusCode, response)
}
