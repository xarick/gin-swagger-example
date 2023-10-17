package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-swagger-example/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		header_auth := ctx.GetHeader("Authorization")
		if len(header_auth) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -101, "error": "login or password empty"})
			return
		}

		login, password, err := utils.BasicAuthLogPass(header_auth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -102, "error": err.Error()})
			return
		}

		if login == "admin" && password == "admin" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "login or password incorrect"})
			return
		}
	}
}
