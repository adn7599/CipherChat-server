package message

import (
	"advait/chatServer/controller/auth"
	"advait/chatServer/controller/message/websocket"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(router *gin.RouterGroup){
	websocketRouter := router.Group("/websocket")
	websocket.ApplyRouter(websocketRouter)
	
	protectedRouter := router.Use(auth.JwtAuthMiddleware())
	protectedRouter.POST("/sendMessage",sendMessageHandler);
	protectedRouter.GET("/receiveMessage",receiveMessageHandler);

	//protectedRouter.GET("/websocket/connect",websocket.WebsocketConnectionHandler)
}
