package controllers

import (
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/AlyssonT/CheckpointBackend/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controllers struct {
	UserController  *UserController
	LoginController *LoginController
	GameController  *GameController
}

func NewControllers(handlers *handlers.Handlers) *Controllers {
	return &Controllers{
		UserController:  NewUserControllers(handlers),
		LoginController: NewLoginControllers(handlers),
		GameController:  NewGameControllers(handlers),
	}
}

func DefineControllers(handlers *handlers.Handlers, server *gin.Engine) {
	controllers := NewControllers(handlers)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Static("/uploads/avatars", "./static/avatars")

	authorized := server.Group("/")
	authorized.Use(middlewares.Authenticate(handlers.LoginHandlers.JwtService))
	{
		authorized.GET("/me", controllers.UserController.Me)
		authorized.POST("/logout", controllers.LoginController.Logout)
		authorized.GET("/games", controllers.GameController.GetGames)
		authorized.GET("/games/:gameId", controllers.GameController.GetGameById)
		authorized.GET("/games/:gameId/reviews", controllers.GameController.GetGameReviews)
		authorized.GET("/user/profile", controllers.UserController.GetUserProfile)
		authorized.PUT("/user/profile", controllers.UserController.UpdateUserProfileDetails)
		authorized.POST("/user/games", controllers.UserController.AddGameToUser)
		authorized.GET("/user/games", controllers.UserController.GetUserGames)
		authorized.GET("/user/games/:gameId", controllers.UserController.GetUserGameById)
		authorized.PUT("/user/games/:gameId", controllers.UserController.UpdateGameToUser)
		authorized.DELETE("/user/games/:gameId", controllers.UserController.DeleteGameToUser)
	}
	server.POST("/users", controllers.UserController.RegisterUser)
	server.POST("/login", controllers.LoginController.Login)
}
