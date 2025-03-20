package testutilities

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/jaswdr/faker/v2"
)

func BuildFakeUser() communication.RegisterUserRequest {
	f := faker.New()
	return communication.RegisterUserRequest{
		Name:     f.Person().FirstName(),
		Email:    f.Internet().Email(),
		Password: f.Internet().Password(),
	}
}
