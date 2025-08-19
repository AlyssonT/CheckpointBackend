package communication

type Game struct {
	Game_id     uint                `json:"game_id"`
	Metacritic  uint8               `json:"metacritic"`
	Slug        string              `json:"slug"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Imagem      string              `json:"imagem"`
	Genres      []GenreResponseData `json:"genres"`
}

type GenreResponseData struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type GameWithGenres struct {
	Game
	Genres []GenreResponseData `json:"genres"`
}

type GetGamesRequest struct {
	PaginationRequest
	Query string `json:"query"`
}

type UserReview struct {
	UserId   uint   `json:"userId"`
	GameId   uint   `json:"gameId"`
	Username string `json:"username"`
	Review   string `json:"review"`
	Status   uint   `json:"status"`
	Score    uint   `json:"score"`
}

type ReviewsAdditionalData struct {
	AverageRating uint `json:"averageRating"`
	Playing       uint `json:"playing"`
	Finished      uint `json:"finished"`
	Backlog       uint `json:"backlog"`
	Dropped       uint `json:"dropped"`
}

type GameReviewsResponse struct {
	ReviewsAdditionalData
	Reviews    []UserReview `json:"reviews"`
	TotalItems int64        `json:"totalItems"`
}

type GameReviewsRequest struct {
	PaginationRequest
}

const (
	StatusPlaying = iota
	StatusFinished
	StatusBacklog
	StatusDropped
)
