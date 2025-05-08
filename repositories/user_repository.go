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
	tx := ur.dbConnection.Begin()

	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := tx.Create(newUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	profile := &models.UserProfile{
		UserID:    newUser.ID,
		Bio:       "",
		AvatarURL: "",
	}

	if err := tx.Create(profile).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

func (ur *UserRepository) GetUserProfileDetails(userID uint) (*models.UserProfile, error) {
	var userPrfileDetails models.UserProfile
	result := ur.dbConnection.Where(&models.UserProfile{UserID: userID}).First(&userPrfileDetails)

	if result.Error != nil {
		return nil, result.Error
	}

	return &userPrfileDetails, nil
}

func (ur *UserRepository) UpdateUserProfileDetails(userProfileDetails *models.UserProfile) error {
	result := ur.dbConnection.Model(userProfileDetails).
		Updates(models.UserProfile{
			AvatarURL: userProfileDetails.AvatarURL,
			Bio:       userProfileDetails.Bio,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
