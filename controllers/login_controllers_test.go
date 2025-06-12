package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/services"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	server, _ := setupApiForTest()
	w := httptest.NewRecorder()

	user := testutilities.BuildFakeUser()
	password := user.Password
	jsonRequest, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	jsonRequest, _ = json.Marshal(&communication.LoginRequest{Email: user.Email, Password: password})

	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	var responseJSON map[string]string
	json.Unmarshal(w.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, w.Code)

	jwtService := services.NewJwt()
	_, err := jwtService.VerifyToken(responseJSON["data"])

	assert.Nil(t, err)
}

func TestLogin_Fail(t *testing.T) {
	server, _ := setupApiForTest()
	w := httptest.NewRecorder()

	user := testutilities.BuildFakeUser()
	password := user.Password
	jsonRequest, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	loginRequestList := []communication.LoginRequest{
		{Email: "another@email.com", Password: password},
		{Email: "invalid email", Password: password},
		{Email: user.Email, Password: "wrong password"},
	}

	for _, request := range loginRequestList {
		w = httptest.NewRecorder()
		jsonRequest, _ = json.Marshal(&request)

		req, _ = http.NewRequest("POST", "/login", bytes.NewReader(jsonRequest))
		server.ServeHTTP(w, req)

		var responseJSON map[string]string
		json.Unmarshal(w.Body.Bytes(), &responseJSON)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		jwtService := services.NewJwt()
		_, err := jwtService.VerifyToken(responseJSON["data"])

		assert.ErrorIs(t, exceptions.ErrorInvalidCredentials, err)
	}
}
