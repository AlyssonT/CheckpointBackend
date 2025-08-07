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

	assert.Equal(t, http.StatusOK, w.Code)

	var authCookie *http.Cookie
	cookies := w.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			authCookie = cookie
			break
		}
	}

	assert.NotNil(t, authCookie)
	assert.True(t, authCookie.HttpOnly)
	assert.True(t, authCookie.Secure)
	assert.Equal(t, http.SameSiteStrictMode, authCookie.SameSite)

	jwtService := services.NewJwt()
	_, err := jwtService.VerifyToken(authCookie.Value)

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

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		cookies := w.Result().Cookies()
		var token string
		for _, cookie := range cookies {
			if cookie.Name == "auth_token" {
				token = cookie.Value
				break
			}
		}

		jwtService := services.NewJwt()
		_, err := jwtService.VerifyToken(token)

		assert.ErrorIs(t, exceptions.ErrorInvalidCredentials, err)
	}
}
