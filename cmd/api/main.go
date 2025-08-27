package main

import (
	"time"

	"github.com/AlyssonT/CheckpointBackend/configs"
	"github.com/AlyssonT/CheckpointBackend/controllers"
	"github.com/AlyssonT/CheckpointBackend/db"
	_ "github.com/AlyssonT/CheckpointBackend/docs"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/AlyssonT/CheckpointBackend/repositories"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.BuildConfigs()

	configs := configs.GetConfigs()
	if configs.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

//	@title						Checkpoint API
//	@version					1.0
//	@description				Checkpoint API Docs.
//	@securityDefinitions.apiKey	cookieAuth
//	@in							cookie
//	@name						auth_token
//	@host						localhost:8080

// @schemes	http
func main() {
	dbConnection := db.InitDb()

	repositories := repositories.NewRepositories(dbConnection)

	handlers := handlers.NewHandlers(repositories)

	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{configs.GetConfigs().FrontendURL, "http://localhost:5173"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 24 * time.Hour

	server.Use(cors.New((config)))

	controllers.DefineControllers(handlers, server)

	server.Run()
}
