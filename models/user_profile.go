package models

type UserProfile struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"uniqueIndex"`
	Bio       string `gorm:"type:text"`
	AvatarURL string
}
