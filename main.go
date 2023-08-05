package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run(":8080")
}
