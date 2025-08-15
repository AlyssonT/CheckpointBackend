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

func (gh *GameHandlers) GetGames(req *communication.GetGamesRequest) ([]communication.Game, int64, error) {
	games, totalItems, err := gh.repository.GetGames(req)
	if err != nil {
		return nil, 0, err
	}

	responseGames := make([]communication.Game, len(*games))
	for i, game := range *games {
		responseGames[i] = communication.Game{
			ID:          game.ID,
			Game_id:     game.Game_id,
			Metacritic:  game.Metacritic,
			Slug:        game.Slug,
			Name:        game.Name,
			Description: game.Description,
			Imagem:      game.Imagem,
		}
	}

	return responseGames, totalItems, nil
}

func (gh *GameHandlers) GetGameById(gameId int) (*communication.GameWithGenres, error) {
	game, err := gh.repository.GetGameById(gameId)
	if err != nil {
		return nil, err
	}

	var genresResponse []communication.GenreResponseData
	for _, genre := range game.Genres {
		genresResponse = append(genresResponse, communication.GenreResponseData{
			Id:          genre.ID,
			Description: genre.Name,
		})
	}

	gameResponse := communication.GameWithGenres{
		Game: communication.Game{
			ID:          game.ID,
			Game_id:     game.Game_id,
			Metacritic:  game.Metacritic,
			Slug:        game.Slug,
			Name:        game.Name,
			Description: game.Description,
			Imagem:      game.Imagem,
		},
		Genres: genresResponse,
	}

	return &gameResponse, nil
}
