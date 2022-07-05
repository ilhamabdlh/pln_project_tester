package middleware

import (
	"strings"
	"fmt"

	"pln/jatim/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")
		jwtString := strings.Split(tokenString, "Bearer ")[1]
		fmt.Println("isi jwt: ",jwtString)
		if tokenString == ""{
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		err := auth.ValidateToken(jwtString)
		if err != nil{
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}