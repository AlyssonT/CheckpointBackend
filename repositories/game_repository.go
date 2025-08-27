package repositories

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/gorm"
)

type GameRepository struct {
	dbConnection *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{
		dbConnection: db,
	}
}

func (gr *GameRepository) GetGames(req *communication.GetGamesRequest) ([]models.Game, int64, error) {
	pagination := communication.PaginationRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	var games []models.Game
	var totalItems int64

	scope := gr.dbConnection.Preload("Genres").Model(&models.Game{})
	if req.Query != "" {
		scope = scope.Where("LOWER(name) LIKE LOWER(?)", "%"+req.Query+"%")
	}
	scope.Count(&totalItems)

	scope = scope.Order("metacritic DESC").Scopes(db.Paginate(&pagination))
	result := scope.Find(&games)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return games, totalItems, nil
}

func (gr *GameRepository) GetGameById(gameId int) (*models.Game, error) {
	var game models.Game

	result := gr.dbConnection.Preload("Genres").Model(&models.Game{}).
		Where("game_id = ?", gameId).
		First(&game)

	if result.Error != nil {
		return nil, result.Error
	}

	return &game, nil
}

func (gr *GameRepository) GetGameReviewsData(
	gameId int,
	request *communication.GameReviewsRequest,
) ([]models.UserGame, *communication.ReviewsAdditionalData, int64, error) {
	var gameReviews []models.UserGame

	scope := gr.dbConnection.Model(&models.UserGame{}).Preload("User").Where("game_id = ?", gameId)

	var totalItems int64
	err := scope.Count(&totalItems).Error
	if err != nil {
		return nil, nil, 0, err
	}

	var reviewsAdditionalData communication.ReviewsAdditionalData
	err = gr.dbConnection.Model(&models.UserGame{}).
		Where("game_id = ?", gameId).
		Select(`
			COALESCE(ROUND(AVG(score)), 0) AS average_rating,
			SUM(CASE WHEN status = ? THEN 1 ELSE 0 END) AS playing,
			SUM(CASE WHEN status = ? THEN 1 ELSE 0 END) AS finished,
			SUM(CASE WHEN status = ? THEN 1 ELSE 0 END) AS backlog,
			SUM(CASE WHEN status = ? THEN 1 ELSE 0 END) AS dropped
		`,
			communication.StatusPlaying,
			communication.StatusFinished,
			communication.StatusBacklog,
			communication.StatusDropped,
		).
		Scan(&reviewsAdditionalData).Error
	if err != nil {
		return nil, nil, 0, err
	}

	scope = scope.Scopes(db.Paginate(&request.PaginationRequest))
	err = scope.Find(&gameReviews).Error
	if err != nil {
		return nil, nil, 0, err
	}

	return gameReviews, &reviewsAdditionalData, totalItems, nil
}
