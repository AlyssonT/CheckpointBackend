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
