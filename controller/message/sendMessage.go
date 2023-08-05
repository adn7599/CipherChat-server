package message

import (
	"advait/chatServer/model"
	db "advait/chatServer/model/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendMessageHandler(c *gin.Context){
	var sentMsg model.SentMessage
	userID,_ := c.Get("userID")

	if errBind := c.BindJSON(&sentMsg); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errBind.Error(),
		})
		return 
	}

	var bufferedMsg model.BufferedMessage

	errRec := db.Db.Where("id = ?",sentMsg.Receiver).First(&model.User{}).Error

	if errRec != nil {
		log.Printf("Message Receiver not found : %v", errRec.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Receiver not found",
		})
		return 
	}
	
	bufferedMsg.Sender = userID.(string)
	bufferedMsg.Receiver = sentMsg.Receiver
	bufferedMsg.Message = sentMsg.Message
	bufferedMsg.Send_time = sentMsg.Send_time

	err := db.Db.Create(&bufferedMsg).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Message received by server",
	})

}