package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handleRoot())
	r.Run()
}

func handleRoot() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "gin")
	}
}
