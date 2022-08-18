package controller

import (
	"net/http"
	"fmt"
	//"time"

	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/models"
)

func (h handler) AddGroup(c *gin.Context){
	body := GroupReqBody{}

	if err := c.BindJSON(&body); err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var group models.Group

	group.GroupSet = body.GroupSet
	group.Group = body.Group
	group.IpClass = body.IpClass
	group.Netmask = body.Netmask
	group.IpGateway = body.IpGateway

	if result := h.DB.Create(&group); result.Error != nil{
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &group)
}

func (h handler) UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	body := GroupReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var group models.Group

	if result := h.DB.First(&group, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	group.GroupSet = body.GroupSet
	group.Group = body.Group
	group.IpClass = body.IpClass
	group.Netmask = body.Netmask
	group.IpGateway = body.IpGateway

	h.DB.Save(&group)

	c.JSON(http.StatusOK, &group)
}

func (h handler) GetGroup(c *gin.Context){
	var arrIpAdd []models.IpAddress //value of struct ipAddress
	var arrIpAddress models.IpAddress
	var arrGroups []models.Groups
	var arrGroupes models.Groups
	var arrGroup []models.Group

	if result := h.DB.Find(&arrIpAdd); result.Error != nil{
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	

	if result := h.DB.Find(&arrGroup); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	fmt.Println(arrGroup)

	for _, v:= range arrGroup{
		arrGroupes.ID = v.ID
		arrGroupes.GroupSet = v.GroupSet
		arrGroupes.Group = v.Group
		arrGroupes.IpClass = v.IpClass
		arrGroupes.Netmask = v.Netmask
		arrGroupes.IpGateway = v.IpGateway

		arrGroups = append(arrGroups, arrGroupes)	

	}

	for i:= 0; i < len(arrGroups); i++{
		for j:= 0; j < len(arrIpAdd); j++{
			if arrGroups[i].Group == arrIpAdd[j].Group{
				arrGroups[i].Data = append(arrGroups[i].Data, arrIpAdd[j])
			} else {
				arrGroups[i].Data = append(arrGroups[i].Data, arrIpAddress)
			}
		}
	}

	c.JSON(http.StatusOK, &arrGroups)
}

func (h handler) DeleteGroup(c *gin.Context){
	id := c.Param("id")
	var group models.Group
	
	if result := h.DB.First(&group, id); result.Error != nil{
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&group)
	c.JSON(http.StatusOK, gin.H{
		"value" : gin.H{
			"code" : http.StatusOK,
			"status" : "group was deleted",
		},
	})
}