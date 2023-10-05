package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/winyann/practice/controllers"
)

func AuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	authRoutes.POST("/signup", controllers.RegisterUser)
}