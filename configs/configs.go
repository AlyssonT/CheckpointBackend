package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	FrontendURL string
	Domain      string
}

var configsData Configs

func BuildConfigs() {
	if os.Getenv("ENVIRONMENT") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error on setting environment")
		}
	}

	configsData = Configs{
		FrontendURL: os.Getenv("FRONTEND_URL"),
		Domain:      os.Getenv("DOMAIN"),
	}
}

func GetConfigs() *Configs {
	return &configsData
}
