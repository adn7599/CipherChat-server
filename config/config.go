package config

type Config struct {
	HOST_IP string
	HOST_PORT string
	
	DB_NAME string
	DB_HOST string
	DB_PORT string
	DB_USERNAME string
	DB_PASSWORD string

	TOKEN_HOUR_LIFESPAN int
	API_SECRET string
}

var Conf Config

func (conf *Config) SetConfig(){
	conf.HOST_IP = "localhost"	
	conf.HOST_PORT = "8080"	


	conf.DB_NAME = "chatDB"
	conf.DB_HOST = "127.0.0.1"
	conf.DB_PORT = "3306"
	conf.DB_USERNAME = "app"
	conf.DB_PASSWORD = "password"


	conf.TOKEN_HOUR_LIFESPAN = 24
	conf.API_SECRET = "mysecret"
}
