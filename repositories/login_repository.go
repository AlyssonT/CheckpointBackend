package repositories

import (
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

func (lr *LoginRepository) GetHashedPassword(email string) (string, error) {
	var user models.User
	err := lr.dbConnection.Where(&models.User{Email: email}).Select("password").First(&user).Error

	if err != nil {
		return "", err
	}

	return user.Password, nil
}
