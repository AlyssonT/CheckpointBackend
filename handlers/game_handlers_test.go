package handlers

import (
	"os"
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupGameHandlersForTest(Db *gorm.DB) (*GameHandlers, *gorm.DB) {
	if Db == nil {
		Db = db.SetupTestDb(&models.Game{}, &models.Genre{})
	}
	return NewGameHandlers(repositories.NewRepositories(Db, os.TempDir()+"/avatars")), Db
}

func TestGetGamesHandler_Success(t *testing.T) {
	gameHandlers, db := setupGameHandlersForTest(nil)
	gameListSize := 11

	var games []models.Game
	for range gameListSize {
		games = append(games, testutilities.BuildFakeGame())
	}
	games[len(games)-1].Name = "test123 name"
	db.Create(&games)

	req := communication.GetGamesRequest{
		PaginationRequest: communication.PaginationRequest{
			Page:     1,
			PageSize: gameListSize,
		},
		Query: "",
	}

	resGames, totalItems, err := gameHandlers.GetGames(&req)
	assert.Nil(t, err)

	assert.Equal(t, int64(gameListSize), totalItems)

	expectedNames := make([]string, len(games))
	for i, g := range games {
		expectedNames[i] = g.Name
	}

	resultNames := make([]string, len(resGames))
	for i, g := range resGames {
		resultNames[i] = g.Name
	}

	assert.ElementsMatch(t, expectedNames, resultNames)

	req.Query = "test123"
	resGames, totalItems, err = gameHandlers.GetGames(&req)
	assert.Nil(t, err)

	assert.Equal(t, int64(1), totalItems)
	assert.Equal(t, "test123 name", (resGames)[0].Name)
}

func TestGetGamesHandler_NoQueries(t *testing.T) {
	gameHandlers, db := setupGameHandlersForTest(nil)
	gameListSize := 10

	var games []models.Game
	for range gameListSize {
		games = append(games, testutilities.BuildFakeGame())
	}
	db.Create(&games)

	req := communication.GetGamesRequest{}

	resGames, totalItems, err := gameHandlers.GetGames(&req)
	assert.Nil(t, err)

	assert.Equal(t, int64(gameListSize), totalItems)

	expectedNames := make([]string, len(games))
	for i, g := range games {
		expectedNames[i] = g.Name
	}

	resultNames := make([]string, len(resGames))
	for i, g := range resGames {
		resultNames[i] = g.Name
	}

	assert.ElementsMatch(t, expectedNames, resultNames)
}
