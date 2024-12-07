package controllers

import (
	"multi-user-approval/database"
	"multi-user-approval/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
var notificationService = services.NewNotificationService() // Initialize notification service

// CreateTask creates a new task that needs approval from 3 users
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming the creator's user ID is provided (from a JWT token in the future)
	// Here, we assume the `CreatedBy` field is set based on the user sending the request
	task.Status = "pending" // Set the task status to pending when created
	if result := database.DB.Create(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// For simplicity, assume that the approvers are provided in the request body
	var approvers []models.User
	if err := c.ShouldBindJSON(&approvers); err != nil || len(approvers) != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must select exactly 3 approvers"})
		return
	}

	// Assign approvers and send notifications
	notificationService.NotifyApprovers(task.ID, approvers)

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task_id": task.ID})
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
