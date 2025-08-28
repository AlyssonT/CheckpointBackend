package db

import (
	"fmt"
	"log"
	"time"

	"github.com/AlyssonT/CheckpointBackend/configs"
	"github.com/AlyssonT/CheckpointBackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb(optionalLogMode ...logger.LogLevel) *gorm.DB {
	logLevel := logger.Error
	cfg := configs.GetConfigs()
	sslMode := "require"
	if cfg.Environment == "develop" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Sao_Paulo",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, sslMode,
	)
	dialector := postgres.Open(dsn)

	if len(optionalLogMode) > 0 {
		logLevel = optionalLogMode[0]
	}

	if configs.GetConfigs().Environment == "prod" {
		logLevel = logger.Silent
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatal("Failed to load db")
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Game{},
		&models.UserGame{},
		&models.UserProfile{},
		&models.Genre{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func GetDb() *gorm.DB {
	return db
}
