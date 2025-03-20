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
