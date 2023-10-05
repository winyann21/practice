package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(router *gin.Engine) {
	UsersRoutes(router)
	AuthRoutes(router)
}