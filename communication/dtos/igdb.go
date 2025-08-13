package communication

type IGDBGamesDto struct {
	Id            uint           `json:"id"`
	Websites      []IGDBWebsites `json:"websites"`
	Cover         IGDBCover      `json:"cover"`
	Genres        []IGDBGenre    `json:"genres"`
	Name          string         `json:"name"`
	Release_dates []int          `json:"release_dates"`
	Slug          string         `json:"slug"`
	Summary       string         `json:"summary"`
	Total_rating  float64        `json:"total_rating"`
}

type IGDBWebsites struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	Category int    `json:"category"`
}

type IGDBCover struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type IGDBGenre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IGDBPopularityResultDto struct {
	Id      int `json:"id"`
	Game_id int `json:"game_id"`
}
