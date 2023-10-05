package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/winyann/practice/database"
	"github.com/winyann/practice/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	getUser := database.DB.Find(&users)

	if getUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No users found",
		})
		return
	}

	c.JSON(200, &users)
}