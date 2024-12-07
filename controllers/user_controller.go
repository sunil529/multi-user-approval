package controllers

import (
	"multi-user-approval/database"
	"multi-user-approval/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var foundUser models.User
	if result := database.DB.Where("email = ?", user.Email).First(&foundUser); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}
	token := GenerateJWT(foundUser.ID) // Implement GenerateJWT in utils/jwt.go
	c.JSON(http.StatusOK, gin.H{"token": token})
}
