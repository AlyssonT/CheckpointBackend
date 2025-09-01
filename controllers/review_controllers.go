package controllers

import (
	"net/http"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	gameHandlers   *handlers.GameHandlers
	reviewHandlers *handlers.ReviewHandlers
}

func NewReviewControllers(handlers *handlers.Handlers) *ReviewController {
	return &ReviewController{
		gameHandlers:   handlers.GameHandlers,
		reviewHandlers: handlers.ReviewHandlers,
	}
}

// @Summary		Latest Reviews
// @Description	Get Latest Reviews
// @ID				latestReviews
// @Produce		json
// @Router			/reviews/latest [get]
// @Tags			Reviews
// @Success		200
// @Failure		500
func (rc *ReviewController) GetLatestReviews(ctx *gin.Context) {
	reviews, err := rc.reviewHandlers.GetLatestReviews()

	if err != nil {
		response := exceptions.ErrorHandler(err)
		ctx.JSON(response.StatusCode, response)
		return
	}

	response := communication.ResponseDTO{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       reviews,
	}
	ctx.JSON(response.StatusCode, response)
}
