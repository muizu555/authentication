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

func DeleteTask(c *gin.Context) {
	taskId := c.Query("task_id")

	if taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "taskId is required"})
		return
	}

	var task models.Task

	err := models.DB.Where("id = ?", taskId).Delete(&task).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
