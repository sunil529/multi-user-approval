package controllers

import (
	"multi-user-approval/database"
	"multi-user-approval/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApproveTask(c *gin.Context) {
	var approval models.TaskApproval
	if err := c.ShouldBindJSON(&approval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.DB.Create(&approval); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Check if task has 3 approvals
	var count int64
	database.DB.Model(&models.TaskApproval{}).Where("task_id = ?", approval.TaskID).Count(&count)
	if count >= 3 {
		database.DB.Model(&models.Task{}).Where("id = ?", approval.TaskID).Update("status", "Approved")
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task approved!"})
}
