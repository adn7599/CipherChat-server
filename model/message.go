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
	Sender string	`json:"sender" binding:"required"` 
	Send_time string `json:"send_time" binding:"required"`
	Message string `json:"message" binding:"required"` 
}