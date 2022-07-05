package controller

import (
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/models"
)

func (h handler) Register(c *gin.Context) {
	body := UserReqBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var user models.Users

	user.Name = body.Name
	user.Password = body.Password
	user.Previlage = body.Previlage

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"value": &user,
		"status" : gin.H{
			"code" : http.StatusOK,
			"message": "registration success",
		},
	})
}

func (h handler) GetUsers(c *gin.Context) {
	var users []models.Users

	if result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &users)
}

func (h handler) GetUser(c *gin.Context) {

	id := c.Param("id")

	var user models.Users

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}
	

	c.JSON(http.StatusOK, &user)
}

func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := UserReqBody{}


	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.Users

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	user.Name = body.Name
	user.Password = body.Password
	user.Previlage = body.Previlage

	h.DB.Save(&user)

	c.JSON(http.StatusOK, &user)
}

func (h handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.Users

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"value" : gin.H{
			"code": http.StatusOK,
			"status": "user was deleted",
		},
	})
}