package main

import (
	"advait/chatServer/config"
	"advait/chatServer/controller/auth"
	"advait/chatServer/controller/message"
	"advait/chatServer/controller/user"
	"advait/chatServer/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "server up and running",
	})
}

func main() {

	config.Conf.SetConfig()
	model.InitDB();

	router := gin.Default()
	router.GET("/ping",ping)

	authRouter := router.Group("/auth")
	auth.ApplyRouter(authRouter)

	msgRouter := router.Group("/message")
	message.ApplyRouter(msgRouter);

	userRouter := router.Group("/user")
	user.ApplyRouter(userRouter)

	router.Run(":8080")
}
