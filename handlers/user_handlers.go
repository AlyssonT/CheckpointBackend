package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/AlyssonT/CheckpointBackend/repositories"
)

type UserHandlers struct {
	repository      *repositories.UserRepository
	fileRespository *repositories.FileRepository
	cryptographer   interfaces.Cryptographer
	jwtService      interfaces.JwtService
}

func NewUserHandlers(repos *repositories.Respositories, cryptographer interfaces.Cryptographer, jwtService interfaces.JwtService) *UserHandlers {
	return &UserHandlers{
		repository:      repos.UserRepository,
		fileRespository: repos.FileRepository,
		cryptographer:   cryptographer,
		jwtService:      jwtService,
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
	userID, err := uh.repository.RegisterUser(user)

	if err != nil {
		return "", err
	}

	token, err := uh.jwtService.GenerateToken(user.Name, user.Email, userID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (uh *UserHandlers) UpdateUserProfileDetails(user *communication.UserProfileDetails) error {
	userProfileDetails, err := uh.repository.GetUserProfileDetails(user.UserID)

	if err != nil {
		return err
	}

	if user.AvatarData != nil {
		avatarURL, err := uh.fileRespository.SaveAvatar(user.AvatarData, user.UserID)
		if err != nil {
			return err
		}
		userProfileDetails.AvatarURL = avatarURL
	}

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
		AvatarURL: userProfile.AvatarURL,
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

func (uh *UserHandlers) GetUserGames(userID uint) ([]communication.UserGame, int64, error) {
	userGames, totalItems, err := uh.repository.GetUserGames(userID)

	if err != nil {
		return nil, 0, err
	}

	games := make([]communication.UserGame, len(userGames))
	for i, game := range userGames {
		games[i] = communication.UserGame{
			Game: communication.Game{
				Game_id:     game.Game.Game_id,
				Metacritic:  game.Game.Metacritic,
				Slug:        game.Game.Slug,
				Name:        game.Game.Name,
				Description: game.Game.Description,
				Imagem:      game.Game.Imagem,
			},
			Status: game.Status,
			Score:  game.Score,
			Review: game.UserReview,
		}
	}

	return games, totalItems, nil
}

func (uh *UserHandlers) GetUserGameById(userID uint, gameId uint) (*communication.UserGame, error) {
	userGame, err := uh.repository.GetUserGameById(userID, gameId)

	if err != nil {
		return nil, err
	}

	game := communication.UserGame{
		Game: communication.Game{
			Game_id:     userGame.UserID,
			Metacritic:  userGame.Game.Metacritic,
			Slug:        userGame.Game.Slug,
			Name:        userGame.Game.Name,
			Description: userGame.Game.Description,
			Imagem:      userGame.Game.Imagem,
		},
		Status: userGame.Status,
		Score:  userGame.Score,
		Review: userGame.UserReview,
	}

	return &game, nil
}
