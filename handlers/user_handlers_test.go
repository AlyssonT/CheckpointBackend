package handlers

import (
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	"github.com/AlyssonT/CheckpointBackend/services"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupUserHandlerForTest(Db *gorm.DB) *UserHandlers {
	if Db == nil {
		Db = db.SetupTestDb(&models.User{}, &models.UserProfile{})
	}
	cryptography := services.NewCryptography(services.DefaultCost)
	jwtService := services.NewJwt()

	return NewUserHandlers(repositories.NewRepositories(Db), &cryptography, &jwtService)
}

func TestRegisterUser_Success(t *testing.T) {
	handler := setupUserHandlerForTest(nil)

	user := testutilities.BuildFakeUser()

	token, err := handler.RegisterUser(&user)

	assert.Nil(t, err)

	_, err = handler.jwtService.VerifyToken(token)

	assert.Nil(t, err)
}

func TestRegisterUser_EmailAlreadyExists(t *testing.T) {
	handler := setupUserHandlerForTest(nil)

	user := testutilities.BuildFakeUser()
	_, err := handler.RegisterUser(&user)

	assert.Nil(t, err)

	_, err = handler.RegisterUser(&user)

	assert.NotNil(t, err)
}

func TestUpdateUserDetails_Success(t *testing.T) {
	handler := setupUserHandlerForTest(nil)

	user := testutilities.BuildFakeUser()
	_, err := handler.RegisterUser(&user)

	assert.Nil(t, err)

	err = handler.UpdateUserProfileDetails(&communication.UserProfileDetails{
		UserID: 0,
		Bio:    "New Bio",
	})

	assert.Nil(t, err)
}
