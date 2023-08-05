package auth

import (
	"github.com/gin-gonic/gin"
)

func ApplyRouter(route *gin.RouterGroup){
	route.POST("/register",registerHandler)
	route.POST("/login",loginHandler)
}