package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/configs"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/services"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func updateGenresTableByGameGenres(genres []communication.IGDBGenre, db *gorm.DB) {
	for _, genre := range genres {
		err := db.First(&models.Genre{ID: genre.Id}).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.Genre{
				ID:   genre.Id,
				Name: genre.Name,
			})
		}
	}
}

func addGenresToGame(genres []communication.IGDBGenre, game *models.Game, db *gorm.DB) {
	if len(genres) == 0 {
		return
	}
	var game_genres_models []models.Genre
	for _, game_genre := range genres {
		var genre models.Genre
		db.First(&genre, game_genre.Id)

		game_genres_models = append(game_genres_models, genre)
	}
	if len(game_genres_models) > 0 {
		db.Model(game).Association("Genres").Append(game_genres_models)
	}
}

func saveOnDb(games *[]communication.IGDBGamesDto, db *gorm.DB) {
	skipWords := []string{
		"sexy", "sex", "nude", "nudity",
		"erotic", "pornographic", "adult",
		"sexual", "sensual", "xxx", "fuck",
		"suggestive", "intimate", "risqué",
		"torture", "femboy", "girl", "cum",
		"cumming", "dude", "mommy", "hentai", "harem",
	}

	for _, game := range *games {
		gameName := strings.ToLower(game.Name)

		shouldSkip := false
		for _, word := range skipWords {
			if strings.Contains(gameName, word) {
				shouldSkip = true
				break
			}
		}

		if shouldSkip {
			continue
		}

		newImage := strings.Replace(game.Cover.Url, "t_thumb", "t_cover_big_2x", 1)

		if len(game.Genres) > 0 {
			updateGenresTableByGameGenres(game.Genres, db)
		}

		var foundGame models.Game
		result := db.Where(&models.Game{Game_id: game.Id}).Take(&foundGame)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newGame := &models.Game{
				Game_id:     game.Id,
				Slug:        game.Slug,
				Name:        game.Name,
				Description: game.Summary,
				Imagem:      newImage,
				Metacritic:  uint8(game.Total_rating),
				UpdatedAt:   time.Now(),
			}
			db.Create(newGame)
			addGenresToGame(game.Genres, newGame, db)
		} else {
			db.Model(&foundGame).Updates(models.Game{
				Game_id:     game.Id,
				Slug:        game.Slug,
				Name:        game.Name,
				Description: game.Summary,
				Imagem:      newImage,
				Metacritic:  uint8(game.Total_rating),
				UpdatedAt:   time.Now(),
			})
		}
	}
}

func main() {
	configs.BuildConfigsDbSync()
	configs.BuildConfigs()
	dbConnection := db.InitDb(logger.Silent)

	igdbHelper := services.NewIGDBApiHelper()

	for i := 0; i < 1000; i += 1 {
		fmt.Println(i)
		var games_id []string
		response, err :=
			igdbHelper.
				Route("popularity_primitives").
				Req(fmt.Sprintf("fields game_id; sort value desc; limit 500; offset %d; where popularity_type = 1;", 500*i)).
				Run()
		if err != nil {
			log.Fatal("error on get igdb api data")
		}

		var games_id_respose []communication.IGDBPopularityResultDto
		err = json.Unmarshal(response, &games_id_respose)
		if err != nil {
			log.Fatal("error on parsing response")
		}
		for _, game := range games_id_respose {
			games_id = append(games_id, fmt.Sprintf("%d", game.Game_id))
		}

		response, err =
			igdbHelper.
				Route("games").
				Req("fields genres.name,total_rating,cover.url,name,summary,release_dates,slug,websites.*;where id = (" + strings.Join(games_id, ",") + ");limit 500;").
				Run()
		if err != nil {
			log.Fatal("error on get igdb api data")
		}

		var games []communication.IGDBGamesDto
		err = json.Unmarshal(response, &games)
		if err != nil {
			log.Fatal("error on parsing games data")
		}

		saveOnDb(&games, dbConnection)
	}
}
