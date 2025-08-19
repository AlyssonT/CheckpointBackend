package models

import "time"

type Game struct {
	Game_id     uint   `gorm:"primaryKey"`
	Slug        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Imagem      string
	Metacritic  uint8
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Users       []User    `gorm:"many2many:user_games;"`
	Genres      []Genre   `gorm:"many2many:game_genres;"`
}
