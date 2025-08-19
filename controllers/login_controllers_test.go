package controllers

import (
	"net/http"
	"testing"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/services"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Success(t *testing.T) {
	server, _ := setupApiForTest()

	user := testutilities.BuildFakeUser()
	password := user.Password

	testutilities.MakeRequest(server, "POST", "/user", user, nil)

	w := testutilities.MakeRequest(server, "POST", "/login", communication.LoginRequest{Email: user.Email, Password: password}, nil)

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

	user := testutilities.BuildFakeUser()
	password := user.Password

	testutilities.MakeRequest(server, "POST", "/user", user, nil)

	loginRequestList := []communication.LoginRequest{
		{Email: "another@email.com", Password: password},
		{Email: "invalid email", Password: password},
		{Email: user.Email, Password: "wrong password"},
	}

	for _, request := range loginRequestList {
		w := testutilities.MakeRequest(server, "POST", "/login", request, nil)

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
