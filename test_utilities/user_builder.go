package testutilities

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker/v2"
)

func BuildFakeUser() communication.RegisterUserRequest {
	f := faker.New()
	return communication.RegisterUserRequest{
		Name:     f.Person().FirstName(),
		Email:    f.Internet().Email(),
		Password: f.Internet().Password(),
	}
}

func RegisterFakeUser(server *gin.Engine, w *httptest.ResponseRecorder) []*http.Cookie {
	user := BuildFakeUser()
	jsonRequest, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(jsonRequest))
	server.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		log.Fatal("failed to register fake user")
	}

	login_request := communication.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	jsonRequest, _ = json.Marshal(login_request)
	req, _ = http.NewRequest("POST", "/login", bytes.NewReader(jsonRequest))

	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)

	cookies := w.Result().Cookies()

	return cookies
}
