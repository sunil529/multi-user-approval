package controllers

import (
	"multi-user-approval/database"
	"multi-user-approval/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.DB.Create(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
