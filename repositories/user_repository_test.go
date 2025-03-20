package repositories

import (
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func setupRepositoryForTests() *UserRepository {
	db := db.SetupTestDb(&models.User{})

	return NewUserRepository(db)
}

func TestRegisterUser_Success(t *testing.T) {
	repo := setupRepositoryForTests()

	f := faker.New()
	email := f.Internet().Email()

	err := repo.RegisterUser(&communication.RegisterUserRequest{
		Name:     f.Person().FirstName(),
		Password: f.Internet().Password(),
		Email:    email,
	})
	assert.Nil(t, err)

	var user models.User
	err = repo.dbConnection.Where("email = ?", email).Take(&user).Error

	assert.Nil(t, err)
	assert.Equal(t, email, user.Email)
}

func TestVerifyEmailAlreadyExists(t *testing.T) {
	repo := setupRepositoryForTests()

	f := faker.New()
	email := f.Internet().Email()

	user := communication.RegisterUserRequest{
		Name:     f.Person().FirstName(),
		Password: f.Internet().Password(),
		Email:    email,
	}

	err := repo.RegisterUser(&user)

	assert.Nil(t, err)

	alreadyExists, err := repo.VerifyEmailAlreadyExists(&user)

	assert.Nil(t, err)
	assert.True(t, alreadyExists)

	otherUser := communication.RegisterUserRequest{
		Name:     "Nome",
		Password: "123456",
		Email:    "useremail@email.com",
	}

	alreadyExists, err = repo.VerifyEmailAlreadyExists(&otherUser)

	assert.Nil(t, err)
	assert.False(t, alreadyExists)
}
