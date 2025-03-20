package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDb(models ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil
	}

	return db
}
