package models

type Genre struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null"`
	Games []Game `gorm:"many2many:game_genres;"`
}
