package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	FrontendURL string
	Domain      string
	Environment string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
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
		Environment: os.Getenv("ENVIRONMENT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DBPort:      os.Getenv("DB_PORT"),
	}
}

func GetConfigs() *Configs {
	return &configsData
}
