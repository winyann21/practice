package main

import (
	"github.com/gin-gonic/gin"
	"github.com/winyann/practice/initializers"
	"github.com/winyann/practice/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDatabase()
}

func main() {
	r := gin.Default()
	routes.LoadRoutes(r)
	r.Run()
}