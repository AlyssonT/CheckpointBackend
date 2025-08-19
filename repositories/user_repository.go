package repositories

import (
	"errors"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	exceptions "github.com/AlyssonT/CheckpointBackend/communication/exceptions"
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

func (ur *UserRepository) RegisterUser(user *communication.RegisterUserRequest) (uint, error) {
	tx := ur.dbConnection.Begin()

	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := tx.Create(newUser).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	profile := &models.UserProfile{
		UserID:    newUser.ID,
		Bio:       "",
		AvatarURL: "",
	}

	if err := tx.Create(profile).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	result := tx.Commit()
	if result.Error != nil {
		return 0, result.Error
	}

	return newUser.ID, nil
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

func (ur *UserRepository) GetUser(userID uint) (*models.User, error) {
	var user models.User
	result := ur.dbConnection.Where("ID = ?", userID).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
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

func (ur *UserRepository) AddGameToUser(userID uint, game_data *communication.AddGameToUserRequest) error {
	var user models.User
	if err := ur.dbConnection.First(&user, userID).Error; err != nil {
		return err
	}

	var game models.Game
	if err := ur.dbConnection.First(&game, game_data.Game_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.ErrorGameNotFound
		}
		return err
	}

	if err := ur.dbConnection.First(&models.UserGame{}, "user_id = ? AND game_id = ?", userID, game_data.Game_id).Error; err == nil {
		return exceptions.ErrorGameAlreadyAddedUser
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	userGame := models.UserGame{
		UserID:     user.ID,
		GameID:     game_data.Game_id,
		Status:     game_data.Status,
		Score:      game_data.Score,
		UserReview: game_data.Review,
	}

	if err := ur.dbConnection.Create(&userGame).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserGames(userID uint) ([]models.UserGame, error) {
	var user_games []models.UserGame
	result := ur.dbConnection.
		Preload("Game").
		Where("user_id = ?", userID).
		Find(&user_games)

	if result.Error != nil {
		return nil, result.Error
	}

	return user_games, nil
}

func (ur *UserRepository) UpdateUserGame(userID uint, game_id uint, game_data *communication.UpdateGameToUserRequest) error {
	var userGame models.UserGame
	result := ur.dbConnection.Where("user_id = ? AND game_id = ?", userID, game_id).First(&userGame)

	if result.Error != nil {
		return result.Error
	}

	userGame.Status = game_data.Status
	userGame.Score = game_data.Score
	userGame.UserReview = game_data.Review

	if err := ur.dbConnection.Save(&userGame).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUserGame(userID uint, game_id uint) error {
	var userGame models.UserGame
	result := ur.dbConnection.Where("user_id = ? AND game_id = ?", userID, game_id).First(&userGame)

	if result.Error != nil {
		return result.Error
	}

	if err := ur.dbConnection.Delete(&userGame).Error; err != nil {
		return err
	}

	return nil
}
