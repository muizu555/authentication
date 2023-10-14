package controllers

import (
	"auth-jwt/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	Desc   string `json:"desc" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
}

func CreateTask(c *gin.Context) {
	var input TaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{Desc: input.Desc, UserId: input.UserId}

	newTask, err := task.Save()

	if err != nil {
		log.Println("Error saving task: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newTask,
	})

}
