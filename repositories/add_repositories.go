package repositories

import (
	"gorm.io/gorm"
)

type Respositories struct {
	UserRepository  *UserRepository
	LoginRepository *LoginRepository
}

func NewRepositories(dbConnection *gorm.DB) *Respositories {
	return &Respositories{
		UserRepository:  NewUserRepository(dbConnection),
		LoginRepository: NewLoginRepository(dbConnection),
	}
}
