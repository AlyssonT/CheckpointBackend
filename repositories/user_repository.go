package repositories

import (
	"errors"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbConnection *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		dbConnection: db,
	}
}

func (ur *UserRepository) RegisterUser(user *communication.RegisterUserRequest) error {
	result := ur.dbConnection.Create(&models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) VerifyEmailAlreadyExists(user *communication.RegisterUserRequest) (bool, error) {
	var foundUser models.User
	result := ur.dbConnection.Where(&models.User{Email: user.Email}).Take(&foundUser)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
