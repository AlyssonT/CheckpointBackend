package configs

import (
	"encoding/json"
	"io"
	"log"
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

type ConfigsDbSync struct {
	IGDBClientId               string
	IGDBSecretToken            string
	IGDBAuthorization          string
	IGDBAuthorizationURL       string
	IGDBApiURL                 string
	IGDBAuthorizationExpiresIn int
	SecretKey                  string
	SteamStoreApiURL           string
	SteamApiURL                string
	Environment                string
	DBHost                     string
	DBUser                     string
	DBPassword                 string
	DBName                     string
	DBPort                     string
}

var configsDataDbSync ConfigsDbSync

func BuildConfigsDbSync() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on setting environment")
	}

	configsDataDbSync = ConfigsDbSync{
		IGDBClientId:         os.Getenv("IGDB_CLIENT_ID"),
		IGDBSecretToken:      os.Getenv("IGDB_SECRET_TOKEN"),
		IGDBAuthorizationURL: os.Getenv("IGDB_AUTHORIZATION_URL"),
		IGDBApiURL:           os.Getenv("IGDB_API_URL"),
		SecretKey:            os.Getenv("SECRET_KEY"),
		SteamStoreApiURL:     os.Getenv("STEAM_STORE_API_URL"),
		SteamApiURL:          os.Getenv("STEAM_API_URL"),
		Environment:          os.Getenv("ENVIRONMENT"),
		DBHost:               os.Getenv("DB_HOST"),
		DBUser:               os.Getenv("DB_USER"),
		DBPassword:           os.Getenv("DB_PASSWORD"),
		DBName:               os.Getenv("DB_NAME"),
		DBPort:               os.Getenv("DB_PORT"),
	}

	formData := url.Values{}
	formData.Set("client_id", configsDataDbSync.IGDBClientId)
	formData.Set("client_secret", configsDataDbSync.IGDBSecretToken)
	formData.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", configsDataDbSync.IGDBAuthorizationURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatal("Error on setting environment")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on setting environment")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error on setting environment")
	}

	var tokenData TokenResponse
	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		log.Fatal("Error on setting environment")
	}

	configsDataDbSync.IGDBAuthorization = tokenData.AccessToken
	configsDataDbSync.IGDBAuthorizationExpiresIn = tokenData.ExpiresIn
}

func GetConfigsDbSync() *ConfigsDbSync {
	return &configsDataDbSync
}
