package main

import (
	"auth-jwt/controllers"
	"auth-jwt/middlewares"
	"auth-jwt/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.RegisterUser)
	public.POST("/login", controllers.LoginUser)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8080")
}
