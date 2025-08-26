package repositories

import (
	"gorm.io/gorm"
)

type Respositories struct {
	UserRepository  *UserRepository
	LoginRepository *LoginRepository
	GameRepository  *GameRepository
	FileRepository  *FileRepository
}

func NewRepositories(dbConnection *gorm.DB) *Respositories {
	return &Respositories{
		UserRepository:  NewUserRepository(dbConnection),
		LoginRepository: NewLoginRepository(dbConnection),
		GameRepository:  NewGameRepository(dbConnection),
		FileRepository:  NewFileRepository("./static/avatars"),
	}
}
