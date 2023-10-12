package main

import (
	"auth-jwt/controllers"
	"auth-jwt/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.RegisterUser)
	public.POST("/login", controllers.LoginUser)

	router.Run(":8080")
}
