package communication

type Game struct {
	ID          uint                `json:"id"`
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

const (
	StatusPlaying = iota
	StatusCompleted
	StatusWanted
	StatusDropped
)
