package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestReviews_Success(t *testing.T) {
	server, db := setupApiForTest()

	game_id := testutilities.RegisterFakeGame(db)

	for range 6 {
		cookies, _ := testutilities.RegisterFakeUser(server)
		addGameToUserRequest := communication.AddGameToUserRequest{
			Game_id: game_id,
			Status:  0,
			Score:   uint(rand.Intn(100)),
			Review:  "review",
		}

		w := testutilities.MakeRequest(server, "POST", "/user/games", addGameToUserRequest, cookies)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	w := testutilities.MakeRequest(server, "GET", "/reviews/latest", nil, nil)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)
	reviews, err := testutilities.ConvertDataFromResponse[[]communication.Review](responseJSON.Data)

	assert.Nil(t, err)
	assert.Equal(t, 5, len(reviews))
}
