package repositories

import (
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	dbConnection *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		dbConnection: db,
	}
}

func (rr *ReviewRepository) GetLatestReviews() ([]models.UserGame, error) {
	var reviews []models.UserGame
	err := rr.dbConnection.Model(&models.UserGame{}).
		Preload("Game").
		Preload("User.Profile").
		Limit(5).
		Order("updated_at desc").
		Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}
