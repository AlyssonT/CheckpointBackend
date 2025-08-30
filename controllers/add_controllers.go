package controllers

import (
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/AlyssonT/CheckpointBackend/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controllers struct {
	UserController   *UserController
	LoginController  *LoginController
	GameController   *GameController
	ReviewController *ReviewController
}

func NewControllers(handlers *handlers.Handlers) *Controllers {
	return &Controllers{
		UserController:   NewUserControllers(handlers),
		LoginController:  NewLoginControllers(handlers),
		GameController:   NewGameControllers(handlers),
		ReviewController: NewReviewControllers(handlers),
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
		authorized.PUT("/user/profile", controllers.UserController.UpdateUserProfileDetails)
		authorized.POST("/user/games", controllers.UserController.AddGameToUser)
		authorized.PUT("/user/games/:gameId", controllers.UserController.UpdateGameToUser)
		authorized.DELETE("/user/games/:gameId", controllers.UserController.DeleteGameToUser)
	}
	server.GET("/user/games/:gameId", controllers.UserController.GetUserGameById)
	server.GET("/games", controllers.GameController.GetGames)
	server.GET("/games/:gameId", controllers.GameController.GetGameById)
	server.GET("/games/:gameId/reviews", controllers.GameController.GetGameReviews)
	server.GET("/user/:username/profile", controllers.UserController.GetUserProfile)
	server.GET("/user/:username/games", controllers.UserController.GetUserGames)
	server.GET("/reviews/latest", controllers.ReviewController.GetLatestReviews)
	server.POST("/users", controllers.UserController.RegisterUser)
	server.POST("/login", controllers.LoginController.Login)
}
