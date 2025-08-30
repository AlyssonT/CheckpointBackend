package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/repositories"
)

type ReviewHandlers struct {
	gameRepository   *repositories.GameRepository
	reviewRepository *repositories.ReviewRepository
}

func NewReviewHandlers(repos *repositories.Respositories) *ReviewHandlers {
	return &ReviewHandlers{
		gameRepository:   repos.GameRepository,
		reviewRepository: repos.ReviewRepository,
	}
}

func (rh *ReviewHandlers) GetLatestReviews() ([]communication.Review, error) {
	reviews, err := rh.reviewRepository.GetLatestReviews()

	if err != nil {
		return nil, err
	}

	responseReviews := make([]communication.Review, len(reviews))
	for i, r := range reviews {
		responseReviews[i] = communication.Review{
			User: communication.ReviewUser{
				Name:      r.User.Name,
				AvatarURL: r.User.Profile.AvatarURL,
			},
			Game: communication.ReviewGame{
				Name:     r.Game.Name,
				ImageURL: r.Game.Imagem,
			},
			Status:     r.Status,
			Score:      r.Score,
			UserReview: r.UserReview,
		}
	}

	return responseReviews, nil
}
