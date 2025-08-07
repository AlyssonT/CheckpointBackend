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
)

func saveOnDb(games *[]communication.IGDBGamesDto, db *gorm.DB) {
	skipWords := []string{
		"sexy", "sex", "nude", "nudity",
		"erotic", "pornographic", "adult",
		"sexual", "sensual", "xxx", "fuck",
		"suggestive", "intimate", "risqué",
		"torture", "femboy", "girl", "cum",
		"cumming", "dude",
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

		var foundGame models.Game
		result := db.Where(&models.Game{Game_id: game.Id}).Take(&foundGame)

		newImage := strings.Replace(game.Cover.Url, "t_thumb", "t_720p", 1)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Create(&models.Game{
				Game_id:     game.Id,
				Slug:        game.Slug,
				Name:        game.Name,
				Description: game.Summary,
				Imagem:      newImage,
				UpdatedAt:   time.Now(),
			})
		} else {
			db.Model(&foundGame).Updates(models.Game{
				Game_id:     game.Id,
				Slug:        game.Slug,
				Name:        game.Name,
				Description: game.Summary,
				Imagem:      newImage,
				UpdatedAt:   time.Now(),
			})
		}
	}
}

func main() {
	configs.BuildConfigsDbSync()
	dbConnection := db.InitDb()

	igdbHelper := services.NewIGDBApiHelper()

	for i := 0; i < 3; i += 1 {
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
				Req("fields cover.url,name,summary,release_dates,slug;where id = (" + strings.Join(games_id, ",") + ");limit 500;").
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
