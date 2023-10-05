package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/winyann/practice/controllers"
)

func UsersRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	userRouter.GET("/", controllers.GetUsers)
}