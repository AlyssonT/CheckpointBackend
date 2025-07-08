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

//	@title						Checkpoint API
//	@version					1.0
//	@description				Checkpoint API Docs.
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@host						localhost:8080

// @schemes	http
func main() {
	configs.BuildConfigs()

	dbConnection := db.InitDb()

	repositories := repositories.NewRepositories(dbConnection)

	handlers := handlers.NewHandlers(repositories)

	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	server.Use(cors.New((config)))

	controllers.DefineControllers(handlers, server)

	server.Run()
}
