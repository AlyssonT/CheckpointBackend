package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/configs"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/services"
	"gorm.io/gorm"
)

var shouldContinue = true
var processedIds map[string]bool

func updateGenresTableByGameGenres(genres *[]communication.GenreData, db *gorm.DB) {
	for _, genre := range *genres {
		genre_id, _ := strconv.Atoi(genre.Id)
		err := db.First(&models.Genre{ID: genre_id}).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.Genre{
				ID:   genre_id,
				Name: genre.Description,
			})
		}
	}
}

func addGenresToGame(genres *[]communication.GenreData, game *models.Game, db *gorm.DB) {
	var game_genres_models []models.Genre
	for _, genre := range *genres {
		genre_id, _ := strconv.Atoi(genre.Id)
		var genre models.Genre
		db.First(&genre, genre_id)

		game_genres_models = append(game_genres_models, genre)
	}
	if len(game_genres_models) > 0 {
		db.Model(game).Association("Genres").Append(game_genres_models)
	}
}

func saveOnDb(apps *[]communication.SteamAppData, db *gorm.DB) {
	steamStoreApiHelper := services.NewSteamStoreApiHelper()
	processedIds = ReadProcessedIds()
	fmt.Println(len(processedIds))

	fmt.Println("Starting processing. Press Ctrl+C to stop safely.")

	for _, app := range *apps {

		if !shouldContinue {
			fmt.Println("Stopping processing safely...")
			break
		}

		if processedIds[fmt.Sprint(app.Appid)] {
			fmt.Println("skipping id...", app.Appid)
			continue
		}

		var foundGame models.Game
		result := db.Where(&models.Game{Game_id: app.Appid}).Take(&foundGame)

		time.Sleep(time.Second / 8)

		response, err := steamStoreApiHelper.Route("appdetails?appids=" + fmt.Sprintf("%d", app.Appid)).Run()
		if err != nil {
			continue
		}

		var appResponse communication.SteamAppDetailResponseDto
		err = json.Unmarshal(response, &appResponse)
		if err != nil {
			fmt.Println("error on processing", app.Appid)
			continue
		}

		for _, game := range appResponse {
			fmt.Println(game.Data.Metacritic.Score, game.Data.Name)
		}

		for _, appData := range appResponse {
			app := appData.Data
			processedIds[fmt.Sprint(app.Game_id)] = true
			if app.Type != "game" {
				continue
			}

			if len(app.Genres) > 0 {
				updateGenresTableByGameGenres(&app.Genres, db)
			}

			metacriticScore := uint8(0)
			if app.Metacritic.Score != 0 {
				metacriticScore = app.Metacritic.Score
			}
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				newGame := &models.Game{
					Game_id:     app.Game_id,
					Slug:        strings.ToLower(strings.ReplaceAll(app.Name, " ", "")),
					Name:        app.Name,
					Description: app.Summary,
					Imagem:      app.ImageURL,
					Metacritic:  metacriticScore,
					UpdatedAt:   time.Now(),
				}
				db.Create(newGame)
				if len(app.Genres) > 0 {
					addGenresToGame(&app.Genres, newGame, db)
				}
			} else {
				db.Model(&foundGame).Updates(models.Game{
					Game_id:     app.Game_id,
					Slug:        strings.ToLower(strings.ReplaceAll(app.Name, " ", "")),
					Name:        app.Name,
					Description: app.Summary,
					Imagem:      app.ImageURL,
					Metacritic:  metacriticScore,
					UpdatedAt:   time.Now(),
				})
			}
		}
	}
}

func setupSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nReceived termination signal. Preparing to shut down...")
		shouldContinue = false

		fmt.Println("Saving processed IDs...")
		SaveProcessedIds(processedIds)
		fmt.Println("Processed IDs saved successfully.")
	}()
}

func main() {
	configs.BuildConfigsDbSync()
	dbConnection := db.InitDb()

	setupSignalHandler()

	steamApiHelper := services.NewSteamApiHelper()

	response, err :=
		steamApiHelper.
			Route("ISteamApps/GetAppList/v2").Run()

	if err != nil {
		log.Fatal("error on get steam api data")
	}

	var app_ids_respose communication.SteamAppListResponseDto
	err = json.Unmarshal(response, &app_ids_respose)
	if err != nil {
		log.Fatal("error on parsing response")
	}

	apps := app_ids_respose.Applist.Apps

	saveOnDb(&apps, dbConnection)

	if shouldContinue {
		fmt.Println("Processing completed normally. Saving processed IDs...")
		SaveProcessedIds(processedIds)
		fmt.Println("Processed IDs saved successfully.")
	}

	fmt.Println("Application terminated safely.")
}
