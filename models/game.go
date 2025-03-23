package models

import "time"

type Game struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Slug        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Metacritic  int32
	Release     time.Time
	Imagem      string
	Playtime    int32
	Users       []User `gorm:"many2many:user_games;"`
}
