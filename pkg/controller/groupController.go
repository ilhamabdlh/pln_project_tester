package controller

import (
	"net/http"
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

	var group models.Groups
	
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

func (h handler) GetGroup(c *gin.Context){
	var arrIpAdd []models.IpAddress //value of struct ipAddress
	var arrGroup models.Groups
	var arrGroups []models.Groups

	if result := h.DB.Find(&arrIpAdd); result.Error != nil{
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	mapFilter := make(map[string]bool)

	for _, v := range arrIpAdd{

		if _, ok := mapFilter[v.Group]; !ok{

			mapFilter[v.Group]=true
			arrGroup.GroupSet = v.GroupSet
			arrGroup.Group = v.Group
			arrGroup.IpClass = v.ClassIp
			arrGroup.Netmask = v.Netmask
			arrGroup.IpGateway = v.IpGateway
			arrGroups = append(arrGroups, arrGroup)	
		}

	}

	for i:= 0; i < len(arrGroups); i++{
		for j:= 0; j < len(arrIpAdd); j++{
			if arrGroups[i].Group == arrIpAdd[j].Group{
				arrGroups[i].Data = append(arrGroups[i].Data, arrIpAdd[j])
			}
		}
	}

	c.JSON(http.StatusOK, &arrGroups)
}

func (h handler) DeleteGroup(c *gin.Context){
	id := c.Param("id")
	var group models.Groups
	
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