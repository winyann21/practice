package main

import (
	"github.com/gin-gonic/gin"
	"github.com/winyann/practice/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDatabase()
}

func main() {
	r := gin.Default()
	r.Run()
}