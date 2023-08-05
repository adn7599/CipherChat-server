package database

import (
	conf "advait/chatServer/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const DB_NAME = "chatDB"
// const DB_HOST = "127.0.0.1"
// const DB_PORT = "3306"
// const DB_USERNAME = "app"
// const DB_PASSWORD = "password"

var Db *gorm.DB


func ConnectDatabase() (*gorm.DB){

	url := conf.Conf.DB_USERNAME +":"+ conf.Conf.DB_PASSWORD +"@tcp"+ "(" + conf.Conf.DB_HOST + ":" + conf.Conf.DB_PORT +")/" + conf.Conf.DB_NAME + "?" + "parseTime=true&loc=Local"
	log.Printf("DB url : %s",url)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error while connecting to the Database: error=%v",err)
		return nil
	}
		return db
}

