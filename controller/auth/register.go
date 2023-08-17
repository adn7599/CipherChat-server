package auth

import (
	"advait/chatServer/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(passwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
	return string(hash)
}

func comparePassword(passwd string, storedHash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(storedHash),[]byte(passwd))

	return err == nil
}

func registerHandler(c *gin.Context) {
	var user model.User
	
	if bindError := c.BindJSON(&user); bindError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": bindError.Error(),
		})
		return 
	}
	
	user.Password = hashPassword(user.Password)

	err := model.CreateUser(&user)


	if err != nil {
		log.Printf("Error: %v",err)
		var errMsg string = err.Error()
		if strings.Contains(err.Error(),"1062"){
			errMsg = "Username already registered!!"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})
		return
	}

	token,err1 := generateToken(user.ID)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"msg": "User Registered Successfully",	
		"token": token,
	})
}