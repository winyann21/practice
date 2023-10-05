package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/winyann/practice/database"
	"github.com/winyann/practice/models"
	"github.com/winyann/practice/models/inputs"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var body inputs.Register
	var user models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//trim inputs
	trimmedEmail := strings.TrimSpace(body.Email)
	trimmedName := strings.TrimSpace(body.Name)
	trimmedPassword := strings.TrimSpace(body.Password)

	// Check for duplicate email
	if err := database.DB.Where("email = ?", trimmedEmail).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}
	
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(trimmedPassword), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//create user
	user = models.User{
		Email: trimmedEmail,
		Name: trimmedName,
		Password: string(hashedPassword),
	}
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, &user)
}