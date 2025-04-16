package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/repositories"
)

type GameHandlers struct {
	repository *repositories.GameRepository
}

func NewGameHandlers(repos *repositories.Respositories) *GameHandlers {
	return &GameHandlers{
		repository: repos.GameRepository,
	}
}

func (gh *GameHandlers) GetGames(req *communication.GetGamesRequest) (*[]communication.Game, int64, error) {
	games, totalItems, err := gh.repository.GetGames(req)
	if err != nil {
		return nil, 0, err
	}

	responseGames := make([]communication.Game, len(*games))
	for i, game := range *games {
		responseGames[i] = communication.Game{
			ID:          game.ID,
			Game_id:     game.Game_id,
			Slug:        game.Slug,
			Name:        game.Name,
			Description: game.Description,
			Imagem:      game.Imagem,
		}
	}

	return &responseGames, totalItems, nil
}
