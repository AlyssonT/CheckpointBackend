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

func (gr *GameRepository) GetGames(req *communication.GetGamesRequest) (*[]models.Game, int64, error) {
	pagination := communication.PaginationRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	var games []models.Game
	var totalItems int64

	scope := gr.dbConnection.Model(&models.Game{})
	if req.Query != "" {
		scope = scope.Where("name LIKE ?", "%"+req.Query+"%")
	}
	scope.Count(&totalItems)

	scope = scope.Order("metacritic DESC").Scopes(db.Paginate(&pagination))
	result := scope.Find(&games)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	return &games, totalItems, nil
}

func (gr *GameRepository) GetGameById(gameId int) (*models.Game, error) {
	var game models.Game

	result := gr.dbConnection.Preload("Genres").Where("game_id = ?", gameId).First(&game)

	if result.Error != nil {
		return nil, result.Error
	}

	return &game, nil
}
