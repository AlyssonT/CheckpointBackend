package models

import "time"

type Game struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Game_id     int    `gorm:"unique;not null"`
	Slug        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Imagem      string
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Users       []User    `gorm:"many2many:user_games;"`
}
