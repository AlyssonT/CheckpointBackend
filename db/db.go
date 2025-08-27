package db

import (
	"log"

	"github.com/AlyssonT/CheckpointBackend/configs"
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb(optionalLogMode ...logger.LogLevel) *gorm.DB {
	logLevel := logger.Error
	dialector := sqlite.Open("Checkpoint.db")

	if len(optionalLogMode) > 0 {
		logLevel = optionalLogMode[0]
	}

	if configs.GetConfigs().Environment == "prod" {
		logLevel = logger.Silent
		dsn := "host=" + configs.GetConfigs().DBHost + " user=" + configs.GetConfigs().DBUser + " password=" + configs.GetConfigs().DBPassword + " dbname=" + configs.GetConfigs().DBName + " port=" + configs.GetConfigs().DBPort + " sslmode=require TimeZone=America/Sao_Paulo"
		dialector = postgres.Open(dsn)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
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
