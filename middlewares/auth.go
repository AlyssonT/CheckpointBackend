package middlewares

import (
	"strings"

	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/services"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := services.NewJwt()
		authHeader := ctx.GetHeader("Authorization")
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		claims, err := jwtService.VerifyToken(token)

		if token == "" || err != nil {
			response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
			ctx.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		userID, ok := claims["id"].(float64)

		if !ok {
			response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
			ctx.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		ctx.Set("userID", uint(userID))
		ctx.Next()
	}
}
