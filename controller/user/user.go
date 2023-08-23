package user

import (
	"advait/chatServer/model"
	db "advait/chatServer/model/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(route *gin.RouterGroup){
	route.GET("/search",searchUserHandler);
}

func searchUserHandler(c *gin.Context){
	
	username := c.DefaultQuery("username","")

	if(username == ""){
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Please provide the query string in username",
		})
		return
	}

	var res []model.UserSearchResp
	print("Query : " + username)
	db.Db.Table("users").Select("id","public_key").Where("id LIKE ?","%"+username+"%").Scan(&res)

	c.JSON(http.StatusOK,res)
}