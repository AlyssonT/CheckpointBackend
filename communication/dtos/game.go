package communication

type Game struct {
	ID          uint   `json:"id"`
	Game_id     int    `json:"game_id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Imagem      string `json:"imagem"`
}

type GetGamesRequest struct {
	PaginationRequest
	Query string `json:"query"`
}
