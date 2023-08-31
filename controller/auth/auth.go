package auth

import (
	"github.com/gin-gonic/gin"
)

func ApplyRouter(route *gin.RouterGroup){
	route.POST("/register",registerHandler)
	route.POST("/login",loginHandler)
	protectedRoute := route.Use(JwtAuthMiddleware())
	protectedRoute.POST("/changePassword",changePasswordHandler)
	protectedRoute.POST("/changeMasterKey",changeMasterKeyHandler)
	// protectedRoute.POST("/updateDetails",updateDetailsHandler)
	// protectedRoute.DELETE("/deleteUser",deleteUserHandler)
}