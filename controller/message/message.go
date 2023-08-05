package message

import (
	"advait/chatServer/controller/auth"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(router *gin.RouterGroup){
	protectedRouter := router.Use(auth.JwtAuthMiddleware())
	protectedRouter.POST("/sendMessage",sendMessageHandler);
	protectedRouter.GET("/receiveMessage",receiveMessageHandler);
}