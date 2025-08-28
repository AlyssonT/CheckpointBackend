package testutilities

import (
	"log"
	"net/http"

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

func RegisterFakeUser(server *gin.Engine) ([]*http.Cookie, communication.RegisterUserRequest) {
	user := BuildFakeUser()

	w := MakeRequest(server, "POST", "/user", user, nil)

	if w.Code != http.StatusCreated {
		log.Fatal("failed to register fake user")
	}

	login_request := communication.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}

	w = MakeRequest(server, "POST", "/login", login_request, nil)

	cookies := w.Result().Cookies()

	return cookies, user
}
