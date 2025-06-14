package main

import (
	"github.com/AlyssonT/CheckpointBackend/configs"
	"github.com/AlyssonT/CheckpointBackend/controllers"
	"github.com/AlyssonT/CheckpointBackend/db"
	_ "github.com/AlyssonT/CheckpointBackend/docs"
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/AlyssonT/CheckpointBackend/repositories"
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

	controllers.DefineControllers(handlers, server)

	server.Run()
}
