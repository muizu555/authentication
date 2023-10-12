package middlewares

import (
	"auth-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.TokenValid(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort() //ミドルウェアを途中で打ち切るもの
			return
		}

		c.Next()
	}
}
