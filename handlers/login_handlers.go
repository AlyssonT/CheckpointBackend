package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/AlyssonT/CheckpointBackend/repositories"
)

type LoginHandlers struct {
	repository    *repositories.LoginRepository
	cryptographer interfaces.Cryptographer
	jwtService    interfaces.JwtService
}

func NewLoginHandlers(repos *repositories.Respositories, cryptographer interfaces.Cryptographer, jwtService interfaces.JwtService) *LoginHandlers {
	return &LoginHandlers{
		repository:    repos.LoginRepository,
		cryptographer: cryptographer,
		jwtService:    jwtService,
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

	token, err := uh.jwtService.GenerateToken(credentials.Email, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
