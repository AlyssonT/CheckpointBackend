package db

import (
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("Checkpoint.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to load db")
	}

	db.AutoMigrate(&models.User{}, &models.Game{}, &models.UserGame{})

	return db
}

func GetDb() *gorm.DB {
	return db
}
