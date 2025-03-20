package handlers

import (
	"testing"

	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	"github.com/AlyssonT/CheckpointBackend/services"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
)

func setupHandlerForTest() *UserHandlers {
	cryptography := services.NewCryptography(services.DefaultCost)

	return NewUserHandlers(repositories.NewRepositories(db.SetupTestDb(&models.User{})), &cryptography)
}

func TestRegisterUser_Success(t *testing.T) {
	handler := setupHandlerForTest()

	user := testutilities.BuildFakeUser()

	name, err := handler.RegisterUser(&user)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, name)
}

func TestRegisterUser_EmailAlreadyExists(t *testing.T) {
	handler := setupHandlerForTest()

	user := testutilities.BuildFakeUser()
	_, err := handler.RegisterUser(&user)

	assert.Nil(t, err)

	_, err = handler.RegisterUser(&user)

	assert.NotNil(t, err)
}
