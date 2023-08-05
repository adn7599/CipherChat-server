package model

import "gorm.io/gorm"


type BufferedMessage struct {
	gorm.Model
	sender string 
	receiver string 
	send_time string
	message string
}