package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/db"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/AlyssonT/CheckpointBackend/middlewares"
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupApiForTest() (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)

	testServer := gin.Default()
	dbtest := db.SetupTestDb(&models.User{}, &models.UserProfile{}, &models.Game{}, &models.UserGame{})
	handlers := handlers.NewHandlers(repositories.NewRepositories(dbtest, os.TempDir()+"/avatars"))
	userControllers := NewUserControllers(handlers)
	loginControllers := NewLoginControllers(handlers)

	testServer.POST("/user", userControllers.RegisterUser)
	testServer.POST("/login", loginControllers.Login)
	authorized := testServer.Group("/")
	authorized.Use(middlewares.Authenticate(handlers.LoginHandlers.JwtService))
	{
		authorized.GET("/user/:username/games", userControllers.GetUserGames)
		authorized.GET("/user/games/:gameId", userControllers.GetUserGameById)
		authorized.POST("/user/games", userControllers.AddGameToUser)
		authorized.PUT("/user/games/:gameId", userControllers.UpdateGameToUser)
		authorized.DELETE("/user/games/:gameId", userControllers.DeleteGameToUser)
	}

	return testServer, dbtest
}

func TestRegisterUser_Success(t *testing.T) {
	server, _ := setupApiForTest()
	user := testutilities.BuildFakeUser()

	w := testutilities.MakeRequest(server, "POST", "/user", user, nil)

	var responseJSON map[string]string
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusCreated, w.Code)

	login := communication.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	w = testutilities.MakeRequest(server, "POST", "/login", login, nil)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestValidateUser(t *testing.T) {
	server, _ := setupApiForTest()
	user := communication.RegisterUserRequest{
		Name:     "",
		Password: "123",
		Email:    "invalid email",
	}

	w := testutilities.MakeRequest(server, "POST", "/user", user, nil)

	messages, err := testutilities.ExtractAllMessagesFromResponse(w)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	expectedMessages := []string{
		"Field 'Name' is required.",
		"Field 'Email' should be a valid e-mail.",
		"Field 'Password' should have at least 6 characters.",
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.ElementsMatch(t, expectedMessages, messages)
}

func TestAddGameToUser_Success(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddGameToUser_FailValidation(t *testing.T) {
	server, _ := setupApiForTest()
	cookies, _ := testutilities.RegisterFakeUser(server)

	add_game_request := communication.AddGameToUserRequest{
		Status: 5,
		Score:  101,
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	messages, err := testutilities.ExtractAllMessagesFromResponse(w)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	expectedMessages := []string{
		"Field 'Game_id' is required.",
		"Field 'Status' should be one of [0 1 2 3].",
		"Field 'Score' should have 100 or less characters.",
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.ElementsMatch(t, expectedMessages, messages)
}

func TestAddGameToUser_FailGameDontExist(t *testing.T) {
	server, _ := setupApiForTest()
	cookies, _ := testutilities.RegisterFakeUser(server)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: 1,
		Status:  1,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateGameToUser_Success(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	update_game_request := communication.UpdateGameToUserRequest{
		Status: 1,
		Score:  95,
		Review: "Amazing game!",
	}

	w = testutilities.MakeRequest(server, "PUT", fmt.Sprintf("/user/games/%d", game_id), update_game_request, cookies)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateGameToUser_FailValidation(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	update_game_request := communication.UpdateGameToUserRequest{
		Status: 10,
		Score:  101,
		Review: "Amazing game!",
	}

	w = testutilities.MakeRequest(server, "PUT", fmt.Sprintf("/user/games/%d", game_id), update_game_request, cookies)

	messages, err := testutilities.ExtractAllMessagesFromResponse(w)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	expectedMessages := []string{
		"Field 'Status' should be one of [0 1 2 3].",
		"Field 'Score' should have 100 or less characters.",
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.ElementsMatch(t, expectedMessages, messages)
}

func TestUpdateGameToUser_FailGameDoesntExist(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	update_game_request := communication.UpdateGameToUserRequest{
		Status: 1,
		Score:  95,
		Review: "Amazing game!",
	}

	w = testutilities.MakeRequest(server, "PUT", fmt.Sprintf("/user/games/%d", 1000001), update_game_request, cookies)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteGameToUser_Success(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	w = testutilities.MakeRequest(server, "DELETE", fmt.Sprintf("/user/games/%d", game_id), nil, cookies)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteGameToUser_Fail(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	w = testutilities.MakeRequest(server, "DELETE", fmt.Sprintf("/user/games/%d", 1000001), nil, cookies)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetGamesUser_Success(t *testing.T) {
	server, db := setupApiForTest()

	cookies, user := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	w = testutilities.MakeRequest(server, "GET", fmt.Sprintf("/user/%s/games", user.Name), nil, cookies)
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	userGamesData, err := testutilities.ConvertDataFromResponse[communication.UserGamesResponse](responseJSON.Data)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, userGamesData.Games, 1)
}

func TestGetGameUserById_Success(t *testing.T) {
	server, db := setupApiForTest()

	cookies, _ := testutilities.RegisterFakeUser(server)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	w := testutilities.MakeRequest(server, "POST", "/user/games", add_game_request, cookies)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	w = testutilities.MakeRequest(server, "GET", fmt.Sprintf("/user/games/%d", game_id), nil, cookies)
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	userGamesData, err := testutilities.ConvertDataFromResponse[communication.UserGame](responseJSON.Data)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, add_game_request.Review, userGamesData.Review)
}
