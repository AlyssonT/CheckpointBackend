package testutilities

import (
	"time"

	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/jaswdr/faker/v2"
)

var f = faker.New()

func BuildFakeGame() models.Game {
	return models.Game{
		Game_id:     f.Int(),
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
