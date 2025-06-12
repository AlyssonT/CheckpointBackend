package models

import "time"

type UserGame struct {
	UserID     uint `gorm:"primaryKey"`
	GameID     uint `gorm:"primaryKey"`
	User       User `gorm:"foreignKey:UserID"`
	Game       Game `gorm:"foreignKey:GameID"`
	Status     uint `gorm:"not null"`
	Score      uint `gorm:"not null"`
	UserReview string
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
