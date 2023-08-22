package auth

import (
	"advait/chatServer/model"
	db "advait/chatServer/model/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context){
	var user model.UserLogin
	var storedUser model.User
	
	if bindError := c.BindJSON(&user); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return 
	}

	result := db.Db.First(&storedUser,"id = ?",user.ID)

	if result.Error != nil {
		log.Printf("Error while logging in: %v",storedUser)
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "User not found",
		})
		return
	}

	if comparePassword(user.Password,storedUser.Password) {
		

		token,err1 := generateToken(user.ID)

		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
			})
			return 
		}

		c.JSON(http.StatusOK,gin.H{
			"msg":"User logged in!!",
			"token": token,
			"public_key": storedUser.PublicKey,
			"private_key": storedUser.PrivateKey,
		})
		return 
	}
	c.JSON(http.StatusBadRequest,gin.H{
		"error":"Incorrect password",
	})
}