package handlers

import (
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	"github.com/AlyssonT/CheckpointBackend/services"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupLoginHandlerForTest(Db *gorm.DB) (*LoginHandlers, *services.Jwt) {
	if Db == nil {
		Db = db.SetupTestDb(&models.User{})
	}
	cryptography := services.NewCryptography(services.DefaultCost)
	jwt := services.NewJwt()
	return NewLoginHandlers(repositories.NewRepositories(Db), &cryptography, &jwt), &jwt
}

func TestLoginHandler_Success(t *testing.T) {
	db := db.SetupTestDb(&models.User{})
	userHandler := setupUserHandlerForTest(db)

	user := testutilities.BuildFakeUser()
	password := user.Password

	_, err := userHandler.RegisterUser(&user)
	assert.Nil(t, err)

	handler, jwtService := setupLoginHandlerForTest(db)
	token, err := handler.Login(&communication.LoginRequest{Email: user.Email, Password: password})

	assert.Nil(t, err)

	err = jwtService.VerifyToken(token)

	assert.Nil(t, err)
}

func TestLoginHandler_Fail(t *testing.T) {
	db := db.SetupTestDb(&models.User{})
	userHandler := setupUserHandlerForTest(db)

	user := testutilities.BuildFakeUser()
	correctPassword := user.Password
	wrongPassword := "wrong password"

	_, err := userHandler.RegisterUser(&user)
	assert.Nil(t, err)

	handler, _ := setupLoginHandlerForTest(db)
	_, err = handler.Login(&communication.LoginRequest{Email: user.Email, Password: wrongPassword})

	assert.Equal(t, exceptions.ErrorInvalidCredentials, err)

	_, err = handler.Login(&communication.LoginRequest{Email: "wrong@email.com", Password: correctPassword})

	assert.Equal(t, exceptions.ErrorInvalidCredentials, err)
}
