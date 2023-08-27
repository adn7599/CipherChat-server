package websocket

import (
	"advait/chatServer/controller/auth"
	"advait/chatServer/model"
	db "advait/chatServer/model/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ApplyRouter(route *gin.RouterGroup){
	// route.GET("/connect",WebsocketConnectionHandler)	
	
	protectedRoute := route.Use(auth.JwtAuthMiddleware())
	protectedRoute.GET("/connect",WebsocketConnectionHandler)	
}

var wsConn map[string]*websocket.Conn = make(map[string]*websocket.Conn)

func WebsocketConnectionHandler(c *gin.Context){
	// c.JSON(http.StatusOK,gin.H{
	// 	"Hello":"World",
	// })
	// return
	var upgrader = websocket.Upgrader{}
	userid,_ := c.Get("userID")
	userID := userid.(string)

	fmt.Println("WS Connected: " + userID)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	wsConn[userID] = ws

	ws.SetCloseHandler(func(code int, text string) error {
		fmt.Println("Inside websocket close handler")
		if _,ok := wsConn[userID]; ok {
			delete(wsConn,userID)
		}
		return nil
	})

	defer ws.Close()

	//Checking for buffered messages and sending if any
	var buffered []model.BufferedMessage
	res := db.Db.Where("receiver = ?",userID).Find(&buffered)

	if res.Error != nil {
		resp := model.SentMessageResp{StatusType: "error", StatusMsg: "Could not retrieve buffered messages"}
		ws.WriteJSON(resp)
		return
	}

	var messages []model.ReceivedMessage = make([]model.ReceivedMessage,0)

	for _,buff := range buffered {
		messages = append(messages, model.ReceivedMessage{
			Sender: buff.Sender,
			Send_time: buff.Send_time,
			Message: buff.Message,
		})
	}

	err = ws.WriteJSON(messages)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	db.Db.Unscoped().Delete(&buffered)

	for {
		//Read Message from client
		var msg model.SentMessage 
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		
		fmt.Println(msg)
		var resp model.SentMessageResp 

		receiver,ok := wsConn[msg.Receiver] 

		if ok {
			//receiver online
			var msgRec model.ReceivedMessage = model.ReceivedMessage{Sender: userID,Send_time: msg.Send_time,Message: msg.Message}
			err := receiver.WriteJSON(msgRec)
			if err != nil {
				resp = model.SentMessageResp{StatusType: "error", StatusMsg: err.Error()}
			}else{
				resp = model.SentMessageResp{StatusType: "success", StatusMsg: "message sent"}
			}
		}else {
			//receiver offline
			//storing message
			var bufferedMsg model.BufferedMessage = model.BufferedMessage{
				Sender: userID,
				Receiver: msg.Receiver,
				Send_time: msg.Send_time,
				Message: msg.Message,
			}

			err := db.Db.Create(&bufferedMsg).Error

			if err != nil {
				resp = model.SentMessageResp{StatusType: "error", StatusMsg: err.Error()}
			}else {
				resp = model.SentMessageResp{StatusType: "success", StatusMsg: "message buffered"}
			}
		}

		err = ws.WriteJSON(resp)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	if _,ok := wsConn[userID]; ok {
		delete(wsConn,userID)
	}
}