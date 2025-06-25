package communication

import (
	"mime/multipart"
)

type RegisterUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserProfileDetails struct {
	Bio        string
	AvatarData multipart.File
	UserID     uint
}

type UserProfileResponse struct {
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatarUrl"`
	UserID    uint   `json:"userID"`
}

type AddGameToUserRequest struct {
	Game_id uint   `json:"game_id" binding:"required"`
	Status  uint   `json:"status" binding:"oneof=0 1 2 3"`
	Score   uint   `json:"score" binding:"omitempty,min=0,max=100"`
	Review  string `json:"review" binding:"omitempty,max=500"`
}

type UserGamesResponse struct {
	Game_id        uint   `json:"game_id"`
	Game_name      string `json:"game_name"`
	Game_image_url string `json:"game_image_url"`
	Status         uint   `json:"status"`
	Score          uint   `json:"score"`
	Review         string `json:"review"`
}
