package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/AlyssonT/CheckpointBackend/repositories"
)

type UserHandlers struct {
	repository    *repositories.UserRepository
	cryptographer interfaces.Cryptographer
}

func NewUserHandlers(repos *repositories.Respositories, cryptographer interfaces.Cryptographer) *UserHandlers {
	return &UserHandlers{
		repository:    repos.UserRepository,
		cryptographer: cryptographer,
	}
}

func (uh *UserHandlers) RegisterUser(user *communication.RegisterUserRequest) (string, error) {
	alreadyExists, err := uh.repository.VerifyEmailAlreadyExists(user)

	if err != nil {
		return "", err
	}

	if alreadyExists {
		return "", exceptions.ErrorEmailAlreadyExists
	}

	hashedPassword, err := uh.cryptographer.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	err = uh.repository.RegisterUser(user)

	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func (uh *UserHandlers) UpdateUserProfileDetails(user *communication.UserProfileDetails) error {
	userProfileDetails, err := uh.repository.GetUserProfileDetails(user.UserID)

	if err != nil {
		return err
	}

	//TODO CALL CDN TO CREATE URL

	userProfileDetails.AvatarURL = "new url"
	userProfileDetails.Bio = user.Bio
	err = uh.repository.UpdateUserProfileDetails(userProfileDetails)

	if err != nil {
		return err
	}

	return nil
}

func (uh *UserHandlers) GetUserProfile(parsedID uint) (*communication.UserProfileResponse, error) {
	userProfile, err := uh.repository.GetUserProfileDetails(parsedID)

	if err != nil {
		return nil, err
	}

	return &communication.UserProfileResponse{
		Bio:       userProfile.Bio,
		AvatarURL: userProfile.Bio,
		UserID:    parsedID,
	}, nil
}

func (uh *UserHandlers) AddGameToUser(userID uint, request *communication.AddGameToUserRequest) error {
	return uh.repository.AddGameToUser(userID, request)
}

func (uh *UserHandlers) UpdateGameToUser(userID uint, game_id uint, request *communication.UpdateGameToUserRequest) error {
	return uh.repository.UpdateUserGame(userID, game_id, request)
}

func (uh *UserHandlers) DeleteGameToUser(userID uint, request *communication.DeleteGameToUserRequest) error {
	return uh.repository.DeleteUserGame(userID, request.Game_id)
}

func (uh *UserHandlers) GetUserGames(userID uint) (*[]communication.UserGamesResponse, error) {
	userGames, err := uh.repository.GetUserGames(userID)

	if err != nil {
		return nil, err
	}

	games := make([]communication.UserGamesResponse, len(*userGames))
	for i, game := range *userGames {
		games[i] = communication.UserGamesResponse{
			Game_id:        game.GameID,
			Game_name:      game.Game.Name,
			Game_image_url: game.Game.Imagem,
			Status:         game.Status,
			Score:          game.Score,
			Review:         game.UserReview,
		}
	}

	return &games, nil
}
