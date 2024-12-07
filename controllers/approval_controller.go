package controllers

import (
	"multi-user-approval/database"
	"multi-user-approval/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
var notificationService = services.NewNotificationService() // Initialize notification service

// ApproveTask handles task approval from a specific approver
func ApproveTask(c *gin.Context) 
	taskID := c.Param("task_id")
	var approvalRequest models.Approval
	if err := c.ShouldBindJSON(&approvalRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the task exists
	var task models.Task
	if result := database.DB.First(&task, taskID); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// Update approval record
	var approval models.Approval
	if result := database.DB.Where("task_id = ? AND user_id = ?", taskID, approvalRequest.UserID).First(&approval); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Approval record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	// Update the approval status and comment
	approval.Approved = approvalRequest.Approved
	approval.Comment = approvalRequest.Comment
	if result := database.DB.Save(&approval); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Check if all 3 users have approved the task
	var approvedCount int
	if result := database.DB.Model(&models.Approval{}).Where("task_id = ? AND approved = true", taskID).Count(&approvedCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if approvedCount == 3 {
		// Update task status to 'approved' when all 3 users have approved
		if result := database.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("status", "approved"); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		// Send email notifications to all users involved (implement email logic here)
		var creator models.User
		if result := database.DB.Where("id = ?", task.CreatedBy).First(&creator); result.Error == nil {
			notificationService.NotifyCreator(taskID, creator)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task approval updated"})
}
// AssignApprovers assigns 3 users to approve the task (for demonstration purposes)
func AssignApprovers(c *gin.Context) {
	taskID := c.Param("task_id")

	// Assume approvers are sent in the request body
	var approvers []int
	if err := c.ShouldBindJSON(&approvers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if exactly 3 users are selected as approvers
	if len(approvers) != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must select exactly 3 approvers"})
		return
	}

	// Update the task with approvers (dummy implementation)
	for _, approverID := range approvers {
		approval := models.Approval{
			TaskID:   taskID,
			UserID:   approverID,
			Approved: false,
			Comment:  "",
		}
		if result := database.DB.Create(&approval); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Approvers assigned successfully"})
}
