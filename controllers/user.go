package controllers

import (
	"auth-jwt/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var allUsers []models.User //スライスを作る  //ここでmodels.Taskとしたら、Tasksテーブルを参照することになる

	err := models.DB.Find(&allUsers).Error

	if err != nil {
		log.Println("Error fetching users: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allUsers,
	})

}
