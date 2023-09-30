package main

import (
	"auth-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.RegisterUser)

	router.Run(":8080")
}
