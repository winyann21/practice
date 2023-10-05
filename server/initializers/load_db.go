package initializers

import (
	"github.com/winyann/practice/database"
	"github.com/winyann/practice/models"
)

func LoadDatabase() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Entry{})
}