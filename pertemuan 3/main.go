package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Hello world",
	})
}

func main() {
	webServer := gin.Default()

	webServer.GET("/", HelloWorld)
	webServer.Run(":3000")
}
