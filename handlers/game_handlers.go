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
		genresResponse := make([]communication.GenreResponseData, len(game.Genres))
		for i, genre := range game.Genres {
			genresResponse[i] = communication.GenreResponseData{
				Id:          genre.ID,
				Description: genre.Name,
			}
		}
		responseGames[i] = communication.Game{
			Game_id:     game.Game_id,
			Metacritic:  game.Metacritic,
			Slug:        game.Slug,
			Name:        game.Name,
			Description: game.Description,
			Imagem:      game.Imagem,
			Genres:      genresResponse,
		}
	}

	return responseGames, totalItems, nil
}

func (gh *GameHandlers) GetGameById(gameId int) (*communication.GameWithGenres, error) {
	game, err := gh.repository.GetGameById(gameId)
	if err != nil {
		return nil, err
	}

	genresResponse := make([]communication.GenreResponseData, len(game.Genres))
	for i, genre := range game.Genres {
		genresResponse[i] = communication.GenreResponseData{
			Id:          genre.ID,
			Description: genre.Name,
		}
	}

	gameResponse := communication.GameWithGenres{
		Game: communication.Game{
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

func (gh *GameHandlers) GetGameReviews(gameId int, request *communication.GameReviewsRequest) (*communication.GameReviewsResponse, error) {
	gameReviews, reviewsAdditionalData, totalItems, err := gh.repository.GetGameReviewsData(gameId, request)
	if err != nil {
		return nil, err
	}

	reviewsMapped := make([]communication.UserReview, len(gameReviews))
	for i, review := range gameReviews {
		reviewsMapped[i] = communication.UserReview{
			UserId:   review.UserID,
			GameId:   review.GameID,
			Username: review.User.Name,
			Review:   review.UserReview,
			Status:   review.Status,
			Score:    review.Score,
		}
	}
	gameReviewsResponse := communication.GameReviewsResponse{
		ReviewsAdditionalData: *reviewsAdditionalData,
		Reviews:               reviewsMapped,
		TotalItems:            totalItems,
	}

	return &gameReviewsResponse, nil
}
