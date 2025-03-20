package handlers

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/AlyssonT/CheckpointBackend/repositories"
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

func (uh *LoginHandlers) Login(credentials *communication.LoginRequest) (bool, error) {
	dbHashedPassword, err := uh.repository.GetHashedPassword(credentials.Email)
	if err != nil {
		return false, err
	}

	validPassword := uh.cryptographer.CheckPassword(dbHashedPassword, credentials.Password)

	return validPassword, nil
}
