package communication

type SteamAppListResponseDto struct {
	Applist steamApps `json:"applist"`
}

type steamApps struct {
	Apps []SteamAppData `json:"apps"`
}

type SteamAppData struct {
	Name  string `json:"name"`
	Appid uint   `json:"appid"`
}

type MetacriticData struct {
	Score uint8  `json:"score"`
	Url   string `json:"url"`
}

type GenreData struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type GameData struct {
	Type       string         `json:"type"`
	Name       string         `json:"name"`
	Game_id    uint           `json:"steam_appid"`
	Summary    string         `json:"short_description"`
	ImageURL   string         `json:"header_image"`
	Metacritic MetacriticData `json:"metacritic"`
	Genres     []GenreData    `json:"genres"`
}

type SteamAppDetailResponseDto map[string]struct {
	Success bool     `json:"success"`
	Data    GameData `json:"data"`
}
