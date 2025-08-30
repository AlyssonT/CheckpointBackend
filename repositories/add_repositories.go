package repositories

import (
	"gorm.io/gorm"
)

type Respositories struct {
	UserRepository   *UserRepository
	LoginRepository  *LoginRepository
	GameRepository   *GameRepository
	ReviewRepository *ReviewRepository
	FileRepository   *FileRepository
}

func NewRepositories(dbConnection *gorm.DB, avatarPath string) *Respositories {
	return &Respositories{
		UserRepository:   NewUserRepository(dbConnection),
		LoginRepository:  NewLoginRepository(dbConnection),
		GameRepository:   NewGameRepository(dbConnection),
		ReviewRepository: NewReviewRepository(dbConnection),
		FileRepository:   NewFileRepository(avatarPath),
	}
}
