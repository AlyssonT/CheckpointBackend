package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	"github.com/AlyssonT/CheckpointBackend/services"
)

type LoginHandlers struct {
	repository    *repositories.LoginRepository
	cryptographer interfaces.Cryptographer
}

func NewLoginHandlers(repos *repositories.Respositories, cryptographer interfaces.Cryptographer) *LoginHandlers {
	return &LoginHandlers{
		repository:    repos.LoginRepository,
		cryptographer: cryptographer,
	}
}

func (uh *LoginHandlers) Login(credentials *communication.LoginRequest) (string, error) {
	user, err := uh.repository.GetHashedPassword(credentials.Email)
	if err != nil {
		return "", err
	}

	validPassword := uh.cryptographer.CheckPassword(user.Password, credentials.Password)
	if !validPassword {
		return "", exceptions.ErrorInvalidCredentials
	}

	token, err := services.GenerateToken(credentials.Email, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
