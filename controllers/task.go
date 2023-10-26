package controllers

import (
	"auth-jwt/models"
	"fmt"
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

func GetAllTasks(c *gin.Context) {
	userId := c.Query("user_id") // URLのクエリパラメータからuserIdを取得

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	var tasks []models.Task

	err := models.DB.Where("user_id = ?", userId).Find(&tasks).Error
	if err != nil {
		log.Println("Error fetching tasks: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})

}

type UpdateTaskInput struct {
	Desc   string `json:"desc" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
}

func UpdateTask(c *gin.Context) {
	var input UpdateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

type DeleteTaskInput struct {
	UserId int `json:"userId" binding:"required"`
}

func DeleteTask(c *gin.Context) {
	taskId := c.Query("task_id")

	if taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "taskId is required"})
		return
	}

	var input DeleteTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// トークンからユーザーIDを取得
	tokenUserId, tokenErr := models.ExtractTokenId(c)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token validation failed"})
		return
	}

	// ユーザーIDが一致しない場合は認証エラーを返す
	if input.UserId != int(tokenUserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID in the request does not match the token"})
		return
	}

	var task models.Task

	fmt.Println(taskId)

	err := models.DB.Where("id = ?", taskId).Delete(&task).Error

	//err := models.DB.Where("id = ? AND user_id = ?", taskId, input.UserId).Delete(&task).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
