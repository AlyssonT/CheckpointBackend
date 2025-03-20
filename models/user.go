package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	Active    bool      `gorm:"default:true"`
}
