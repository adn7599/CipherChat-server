package model

import "gorm.io/gorm"


type BufferedMessage struct {
	gorm.Model
	Sender string	`json:"sender" binding:"required"` 
	Receiver string `json:"receiver" binding:"required"` 
	Send_time string `json:"send_time" binding:";equired"`
	Message string `json:"message" binding:";equired"` 
}

//Message received by server from the sender
type SentMessage struct {
	Receiver string `json:"receiver" binding:"required"` 
	Send_time string `json:"send_time" binding:"required"`
	Message string `json:"message" binding:"required"` 
}

//Message received by the receiver from the server
type ReceivedMessage struct {
	Type string	`json:"type" binding:"required"` 
	Sender string	`json:"sender" binding:"required"` 
	Send_time string `json:"send_time" binding:"required"`
	Message string `json:"message" binding:"required"` 
}

func NewReceiveMessage(sender string,send_time string, message string ) ReceivedMessage {
	return ReceivedMessage{
		Type: "receive",
		Sender: sender,
		Send_time: send_time,
		Message: message,
	}
}

type SentMessageResp struct {
	Type string `json:"type"`
	StatusType string `json:"status_type"`
	StatusMsg string `json:"status"`
}

func NewSentMessageResp(status_type string, status_msg string) SentMessageResp{
	return SentMessageResp{
		Type: "sent_response",
		StatusType: status_type,
		StatusMsg: status_msg,
	}
}