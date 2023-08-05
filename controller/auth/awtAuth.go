package auth

import (
	"advait/chatServer/config"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
)


func generateToken(user_id string) (string, error) {

	token_lifespan:= config.Conf.TOKEN_HOUR_LIFESPAN

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Conf.API_SECRET))

}

func validateUser(c *gin.Context) (string,string){
	
	bearerToken := c.Request.Header.Get("Authorization")
	var tokenStr string	
	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenStr = strings.Split(bearerToken, " ")[1]
	}else{
		log.Printf("Token not found: %v",bearerToken)
		return "","Token not found"
	}

	token,err := jwt.Parse(tokenStr, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tok.Header["alg"])
		}
		return []byte(config.Conf.API_SECRET), nil
	})

	if err != nil {
		return "",err.Error()	
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID := claims["user_id"].(string)
		return userID, ""
	}else{
		return "","Invalid token"
	}

}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID,err:= validateUser(c)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H {
				"error" : err,
			} )
			c.Abort()
			return
		}
		c.Set("userID",userID)
		c.Next()
	}
}