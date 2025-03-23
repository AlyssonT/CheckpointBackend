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
	"github.com/AlyssonT/CheckpointBackend/models"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupApiForTest() *gin.Engine {
	gin.SetMode(gin.TestMode)

	testServer := gin.Default()
	dbtest := db.SetupTestDb(&models.User{})
	userControllers := NewUserControllers(handlers.NewHandlers(repositories.NewRepositories(dbtest)))
	loginControllers := NewLoginControllers(handlers.NewHandlers(repositories.NewRepositories(dbtest)))

	testServer.POST("/user", userControllers.RegisterUser)
	testServer.POST("/login", loginControllers.Login)

	return testServer
}

func TestRegisterUser_Success(t *testing.T) {
	server := setupApiForTest()
	w := httptest.NewRecorder()

	user := testutilities.BuildFakeUser()
	jsonRequest, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	var responseJSON map[string]string
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, responseJSON["data"], user.Name)
}

func TestValidateUser(t *testing.T) {
	server := setupApiForTest()
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
