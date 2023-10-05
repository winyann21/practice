package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/winyann/practice/database"
	"github.com/winyann/practice/helpers"
	"github.com/winyann/practice/models"
	inputModels "github.com/winyann/practice/models/inputs"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {

	//TODO: FIX VALIDATION OF INPUTS.
	var body inputModels.Register
	var user models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		// If there's a binding error, return an error response as an object
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//trim inputs
	trimmedEmail := strings.TrimSpace(body.Email)
	trimmedName := strings.TrimSpace(body.Name)
	trimmedPassword := strings.TrimSpace(body.Password)

	// Check email format
    if !helpers.IsValidEmail(trimmedEmail) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid email format",
        })
        return
    }

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