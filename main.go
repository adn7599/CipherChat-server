package main

import (
	"advait/chatServer/config"
	"advait/chatServer/controller/auth"
	"advait/chatServer/controller/message"
	"advait/chatServer/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}

func main() {

	config.Conf.SetConfig()
	model.InitDB();

	router := gin.Default()
	router.GET("/hello",hello)

	authRouter := router.Group("/auth")
	auth.ApplyRouter(authRouter)

	msgRouter := router.Group("message")
	message.ApplyRouter(msgRouter);

	router.Run(":8080")


}
