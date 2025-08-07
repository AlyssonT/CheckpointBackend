package middlewares

import (
	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService interfaces.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, _ := ctx.Cookie("auth_token")
		claims, err := jwtService.VerifyToken(token)

		if token == "" || err != nil {
			response := exceptions.ErrorHandler(exceptions.ErrorInvalidCredentials)
			ctx.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		ctx.Set("userID", claims.ID)
		ctx.Set("userData", *claims)
		ctx.Next()
	}
}
