package controllers

import (
	"github.com/AlyssonT/CheckpointBackend/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controllers struct {
	UserController  *UserController
	LoginController *LoginController
}

func NewControllers(handlers *handlers.Handlers) *Controllers {
	return &Controllers{
		UserController:  NewUserControllers(handlers),
		LoginController: NewLoginControllers(handlers),
	}
}

func DefineControllers(handlers *handlers.Handlers, server *gin.Engine) {
	controllers := NewControllers(handlers)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.POST("/users", controllers.UserController.RegisterUser)
	server.POST("/login", controllers.LoginController.Login)
}
