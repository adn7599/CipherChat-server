package message

import (
	"advait/chatServer/model"
	"advait/chatServer/model/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteMessage(receiver string,messages []model.BufferedMessage){
	go func(){
		// log.Printf("Started timer for deletion of messages for receiver  %s",receiver)
		// time.Sleep(time.Duration(config.Conf.BUFFER_DELETE_WAIT) * time.Second)	
		// log.Printf("Deleting messages received by %s",receiver)	
		database.Db.Unscoped().Delete(&messages)
	}()
}

func receiveMessageHandler(c *gin.Context){
	userID,_ := c.Get("userID")
	var messages []model.BufferedMessage
	res := database.Db.Where("receiver = ?",userID.(string)).Find(&messages)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while retrieving messages",
		})
		return 
	}
	defer deleteMessage(userID.(string),messages)
	c.JSON(http.StatusOK,messages)
}