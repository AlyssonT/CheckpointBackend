package communication

type ReviewUser struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarURL"`
}

type ReviewGame struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageURL"`
}

type Review struct {
	User       ReviewUser `json:"user"`
	Game       ReviewGame `json:"game"`
	Status     uint       `json:"status"`
	Score      uint       `json:"score"`
	UserReview string     `json:"userReview"`
}
