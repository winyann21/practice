package initializers

import (
	"github.com/winyann/practice/database"
	"github.com/winyann/practice/models"
)

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})
}