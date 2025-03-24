package configs

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Configs struct {
	IGDBClientId               string
	IGDBSecretToken            string
	IGDBAuthorization          string
	IGDBAuthorizationURL       string
	IGDBApiURL                 string
	IGDBAuthorizationExpiresIn int
	SecretKey                  string
}

var configsData Configs

func BuildConfigs() {
	err := godotenv.Load()
	if err != nil {
		panic("Error on setting environment")
	}

	configsData = Configs{
		IGDBClientId:         os.Getenv("IGDB_CLIENT_ID"),
		IGDBSecretToken:      os.Getenv("IGDB_SECRET_TOKEN"),
		IGDBAuthorizationURL: os.Getenv("IGDB_AUTHORIZATION_URL"),
		IGDBApiURL:           os.Getenv("IGDB_API_URL"),
		SecretKey:            os.Getenv("SECRET_KEY"),
	}

	formData := url.Values{}
	formData.Set("client_id", configsData.IGDBClientId)
	formData.Set("client_secret", configsData.IGDBSecretToken)
	formData.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", configsData.IGDBAuthorizationURL, strings.NewReader(formData.Encode()))
	if err != nil {
		panic("Error on setting environment")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Error on setting environment")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error on setting environment")
	}

	var tokenData TokenResponse
	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		panic("Error on setting environment")
	}

	configsData.IGDBAuthorization = tokenData.AccessToken
	configsData.IGDBAuthorizationExpiresIn = tokenData.ExpiresIn
}

func GetConfigs() *Configs {
	return &configsData
}
