package main

import (
	"multi-user-approval/controllers"
	"multi-user-approval/database"
	"multi-user-approval/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	// Migrate schema
	database.DB.AutoMigrate(&models.User{}, &models.Task{}, &models.TaskApproval{})

	r := gin.Default()

	// Routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks/:id/approve", controllers.ApproveTask)

	r.Run(":8080")
}
