package controller

import (
	"net/http"
	// "fmt"

	"pln/jatim/pkg/auth"
	"pln/jatim/pkg/db"
	"pln/jatim/pkg/models"
	"github.com/gin-gonic/gin"
)

type Login struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context){
	var request Login
	var user models.Users
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := db.Database.Where("username = ?", request.Username).First(&user)
	if record.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username, user.Password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"value" : gin.H{
			"code" : http.StatusOK,
			"username": user.Username,
			"password": user.Password,
			"previlage": user.Previlage,
			"token" : tokenString,
		},
	})


}