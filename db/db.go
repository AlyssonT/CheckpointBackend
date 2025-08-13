package db

import (
	"log"

	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb(optionalLogMode ...logger.LogLevel) *gorm.DB {
	logLevel := logger.Error
	if len(optionalLogMode) > 0 {
		logLevel = optionalLogMode[0]
	}

	db, err := gorm.Open(sqlite.Open("Checkpoint.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatal("Failed to load db")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Game{},
		&models.UserGame{},
		&models.UserProfile{},
		&models.Genre{},
	)

	return db
}

func GetDb() *gorm.DB {
	return db
}
