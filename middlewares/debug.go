package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Debug() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookies := ctx.Request.Cookies()
		fmt.Println("=== DEBUG: Received Cookies ===")
		for _, cookie := range cookies {
			fmt.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
		}

		fmt.Println("=== DEBUG: Headers ===")
		fmt.Printf("Origin: %s\n", ctx.GetHeader("Origin"))
		fmt.Printf("User-Agent: %s\n", ctx.GetHeader("User-Agent"))
		fmt.Printf("Cookie Header: %s\n", ctx.GetHeader("Cookie"))

		ctx.Next()
	}
}
