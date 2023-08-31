package auth

import (
	"advait/chatServer/model"
	db "advait/chatServer/model/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func changePasswordHandler(c *gin.Context){
	var data model.ChangePassword
	var user model.User
	userID,_  := c.Get("userID")

	
	if bindError := c.BindJSON(&data); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return 
	}

	// result := db.Db.First(&user,userID)
	result := db.Db.Where("id = ?",userID).First(&user)

	if result.Error != nil {
		log.Printf("Error while looking for: %v",userID)
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "User not found",
		})
		return
	}

	if comparePassword(data.Old_password,user.Password) {
		
		//User authenticated
		//updating password

		user.Password = hashPassword(data.New_password)

		if result := db.Db.Save(&user); result.Error != nil {

			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"Error while updating password",
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"msg":"Password changed successfully!!",
		})
		return 
	}
	c.JSON(http.StatusBadRequest,gin.H{
		"error":"Incorrect password",
	})

}

func changeMasterKeyHandler(c *gin.Context){
	var req model.ChangeMasterKey;
	var user model.User

	userID,_ := c.Get("userID")

	if bindError := c.BindJSON(&req); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return
	}

	result := db.Db.Where("id = ?",userID).First(&user)

	if result.Error != nil {
		log.Printf("Error while looking for: %v",userID)
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "User not found",
		})
		return
	}

	user.PrivateKey = req.New_encrypted_private_key;

	if result := db.Db.Save(&user); result.Error != nil {

		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error while updating password",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"Password changed successfully!!",
	})
}

// func updateDetailsHandler(c *gin.Context){
// 	var data model.UpdateUser
// 	var user model.User
// 	userID,_  := c.Get("userID")
	
// 	if bindError := c.BindJSON(&data); bindError != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": bindError.Error(),
// 		})
// 		return 
// 	}

// 	result := db.Db.First(&user,userID)

// 	if result.Error != nil {
// 		log.Printf("Error while looking for: %v",userID)
// 		c.JSON(http.StatusBadRequest,gin.H{
// 			"error": "User not found",
// 		})
// 		return
// 	}	


// 		if result := db.Db.Save(&user); result.Error != nil {

// 			c.JSON(http.StatusInternalServerError,gin.H{
// 			"error":"Error while updating details",
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK,gin.H{
// 			"msg":"Details changed successfully!!",
// 		})

// }