package testutilities

import (
	"log"
	"time"

	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
)

var f = faker.New()

func BuildFakeGame() models.Game {
	return models.Game{
		Game_id:     f.UIntBetween(1, 1000000),
		Slug:        f.RandomStringWithLength(10),
		Name:        f.App().Name(),
		Description: f.App().Name(),
		Imagem:      f.Internet().URL(),
		Metacritic:  f.UInt8Between(0, 100),
		UpdatedAt:   time.Now(),
		Users:       nil,
		Genres:      nil,
	}
}

func RegisterFakeGame(db *gorm.DB) uint {
	game := BuildFakeGame()

	if err := db.Create(&game).Error; err != nil {
		log.Fatal("failed to register fake game")
	}

	return game.Game_id
}
