package communication

type IGDBGamesDto struct {
	Id            uint      `json:"id"`
	Cover         IGDBCover `json:"cover"`
	Name          string    `json:"name"`
	Release_dates []int     `json:"release_dates"`
	Slug          string    `json:"slug"`
	Summary       string    `json:"summary"`
}

type IGDBCover struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type IGDBPopularityResultDto struct {
	Id      int `json:"id"`
	Game_id int `json:"game_id"`
}
