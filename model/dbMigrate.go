package model

import db "advait/chatServer/model/database"

func InitDB() {
	db.Db = db.ConnectDatabase()
	autoMigrate()
}
func autoMigrate(){
	db.Db.AutoMigrate(&User{})
}