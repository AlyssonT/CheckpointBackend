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

	isValidCredentials, err := lc.handlers.Login(&request)
	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	if isValidCredentials {
		response := &communication.ResponseDTO{
			StatusCode: http.StatusOK,
			Message:    "logged in",
			Data:       true,
		}
		ctx.JSON(response.StatusCode, response)
	} else {
		response := &communication.ResponseDTO{
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid credentials",
			Data:       false,
		}
		ctx.JSON(response.StatusCode, response)
	}
}
