package handlers

import (
	"github.com/AlyssonT/CheckpointBackend/repositories"
	"github.com/AlyssonT/CheckpointBackend/services"
)

type Handlers struct {
	UserHandlers  *UserHandlers
	LoginHandlers *LoginHandlers
}

func NewHandlers(repositories *repositories.Respositories) *Handlers {
	cryptography := services.NewCryptography(services.DefaultCost)
	jwtService := services.NewJwt()

	return &Handlers{
		UserHandlers:  NewUserHandlers(repositories, &cryptography),
		LoginHandlers: NewLoginHandlers(repositories, &cryptography, &jwtService),
	}
}
