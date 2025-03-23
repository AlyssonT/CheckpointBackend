package models

import "time"

type UserGame struct {
	UserID    uint   `gorm:"primaryKey"`
	GameID    uint   `gorm:"primaryKey"`
	User      User   `gorm:"foreignKey:UserID"`
	Game      Game   `gorm:"foreignKey:GameID"`
	Status    string `gorm:"not null"`
	Score     float32
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
