package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	handlers := handlers.NewHandlers(repositories.NewRepositories(dbtest))
	userControllers := NewUserControllers(handlers)
	loginControllers := NewLoginControllers(handlers)

	testServer.POST("/user", userControllers.RegisterUser)
	testServer.POST("/login", loginControllers.Login)
	authorized := testServer.Group("/")
	authorized.Use(middlewares.Authenticate(handlers.LoginHandlers.JwtService))
	{
		authorized.POST("/user/games", userControllers.AddGameToUser)
	}

	return testServer, dbtest
}

func TestRegisterUser_Success(t *testing.T) {
	server, _ := setupApiForTest()
	w := httptest.NewRecorder()

	user := testutilities.BuildFakeUser()
	jsonRequest, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	var responseJSON map[string]string
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusCreated, w.Code)

	w = httptest.NewRecorder()
	login := communication.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	jsonRequest, _ = json.Marshal(login)
	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestValidateUser(t *testing.T) {
	server, _ := setupApiForTest()
	w := httptest.NewRecorder()

	user := communication.RegisterUserRequest{
		Name:     "",
		Password: "123",
		Email:    "invalid email",
	}

	jsonRequest, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

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
	w := httptest.NewRecorder()

	cookies := testutilities.RegisterFakeUser(server, w)
	game_id := testutilities.RegisterFakeGame(db)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: game_id,
		Status:  0,
		Score:   90,
		Review:  "Great game!",
	}

	jsonRequest, _ := json.Marshal(add_game_request)
	req, _ := http.NewRequest("POST", "/user/games", bytes.NewReader(jsonRequest))
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddGameToUser_FailValidation(t *testing.T) {
	server, _ := setupApiForTest()
	w := httptest.NewRecorder()

	cookies := testutilities.RegisterFakeUser(server, w)

	add_game_request := communication.AddGameToUserRequest{
		Status: 5,
		Score:  101,
	}

	jsonRequest, _ := json.Marshal(add_game_request)
	req, _ := http.NewRequest("POST", "/user/games", bytes.NewReader(jsonRequest))
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)

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
	w := httptest.NewRecorder()

	cookies := testutilities.RegisterFakeUser(server, w)

	add_game_request := communication.AddGameToUserRequest{
		Game_id: 1,
		Status:  1,
		Score:   90,
		Review:  "Great game!",
	}

	jsonRequest, _ := json.Marshal(add_game_request)
	req, _ := http.NewRequest("POST", "/user/games", bytes.NewReader(jsonRequest))
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)

	var responseJSON communication.ResponseDTO
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
