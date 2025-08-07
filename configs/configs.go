package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	FrontendURL string
}

var configsData Configs

func BuildConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on setting environment")
	}

	configsData = Configs{
		FrontendURL: os.Getenv("FRONTEND_URL"),
	}
}

func GetConfigs() *Configs {
	return &configsData
}
