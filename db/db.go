package db

import (
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	Db, error := gorm.Open(sqlite.Open("Checkpoint.db"), &gorm.Config{})

	if error != nil {
		panic("Failed to load db")
	}

	Db.AutoMigrate(&models.User{}, &models.Game{}, &models.UserGame{})

	return Db
}
