package auth

import (
	"advait/chatServer/model"
	db "advait/chatServer/model/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteUserHandler(c *gin.Context){
	var user model.User
	userID,_  := c.Get("userID")

	result := db.Db.First(&user,userID)

	if result.Error != nil {
		log.Printf("Error while looking for: %v",userID)
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "User not found",
		})
		return
	}	

	if result := db.Db.Unscoped().Delete(&user); result.Error != nil {

		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error while deleting user",
		})
			return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"User deleted successfully!!",
	})
}