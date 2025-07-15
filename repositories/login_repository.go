package repositories

import (
	"errors"

	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/gorm"
)

type LoginRepository struct {
	dbConnection *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{
		dbConnection: db,
	}
}

func (lr *LoginRepository) GetHashedPassword(email string) (*models.User, error) {
	var user models.User
	result := lr.dbConnection.Where(&models.User{Email: email}).Select("name", "id", "password").First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &models.User{}, exceptions.ErrorInvalidCredentials
	} else if result.Error != nil {
		return &models.User{}, result.Error
	}

	return &user, nil
}
