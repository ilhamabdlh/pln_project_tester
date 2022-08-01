package controller

import (
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/models"
)

func (h handler) AddIpAddress(c *gin.Context) {
	body := IpReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ip models.IpAddress
	
	ip.ClassIp = body.ClassIp
	ip.Netmask = body.Netmask
	ip.IpAddressed = body.IpAddressed
	ip.DeviceName = body.DeviceName
	ip.DescriptionOne = body.DescriptionOne
	ip.DescriptionTwo = body.DescriptionTwo
	ip.DescriptionThree = body.DescriptionThree
	ip.IpGateway = body.IpGateway
	ip.Location = body.Location
	ip.ActivityStatus = body.ActivityStatus
	ip.GroupSet = body.GroupSet
	ip.Group = body.Group
	ip.IpUsageStatus = body.IpUsageStatus
	ip.Member = body.Member
	ip.Approve = "nil" //jika baru upload akan teridentifikasi nil, jika diapprove akan berubah jadi true, jika tdk diapprv akan menjadi false
	//nil nilai default, approve=true, not approve=false, dalam bentuk string
	if result := h.DB.Create(&ip); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	generateNotification(true, ip.Member, ip.Location) //true for "add" and false for "edit"

	c.JSON(http.StatusCreated, &ip)
}

func (h handler) GetIp(c *gin.Context) {
	id := c.Param("id")
	var ip models.IpAddress

	if result := h.DB.First(&ip, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &ip)
}

func (h handler) GetSearch(c *gin.Context){
	ipAddrs := c.Query("ip_addressed")
	deviceName := c.Query("device_name")
	location := c.Query("location")
	status := c.Query("status")
	ipUse := c.Query("ip_usage_status")

	c.JSON(http.StatusOK, gin.H{
		"ip_addressed": ipAddrs,
		"device_name": deviceName,
		"location": location,
		"status": status,
		"ip_usage_status": ipUse, 
	})
}

func (h handler) GetIps(c *gin.Context){
	
	var ips []models.IpAddress

	if result := h.DB.Find(&ips); result.Error != nil{
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &ips)
}

func (h handler) UpdateIpByAdmin(c *gin.Context) {
	id := c.Param("id")
	body := IpReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ip models.IpAddress

	if result := h.DB.First(&ip, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ip.ClassIp = body.ClassIp
	ip.Netmask = body.Netmask
	ip.IpAddressed = body.IpAddressed
	ip.DeviceName = body.DeviceName
	ip.DescriptionOne = body.DescriptionOne
	ip.DescriptionTwo = body.DescriptionTwo
	ip.DescriptionThree = body.DescriptionThree
	ip.IpGateway = body.IpGateway
	ip.Location = body.Location
	ip.ActivityStatus = body.ActivityStatus
	ip.GroupSet = body.GroupSet
	ip.Group = body.Group
	ip.IpUsageStatus = body.IpUsageStatus
	ip.Member = body.Member
	ip.Approve = body.Approve

	h.DB.Save(&ip)
	
	generateNotification(false, ip.Member, ip.Location) //true for "add" and false for "edit"
	
	c.JSON(http.StatusOK, &ip)
}

func (h handler) UpdateIpByUser(c *gin.Context) {
	id := c.Param("id")
	body := IpReqBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ip models.IpAddress

	if result := h.DB.First(&ip, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ip.ClassIp = body.ClassIp
	ip.Netmask = body.Netmask
	ip.IpAddressed = body.IpAddressed
	ip.DeviceName = body.DeviceName
	ip.DescriptionOne = body.DescriptionOne
	ip.DescriptionTwo = body.DescriptionTwo
	ip.DescriptionThree = body.DescriptionThree
	ip.IpGateway = body.IpGateway
	ip.Location = body.Location
	ip.ActivityStatus = body.ActivityStatus
	ip.GroupSet = body.GroupSet
	ip.Group = body.Group
	ip.IpUsageStatus = body.IpUsageStatus
	ip.Member = body.Member
	ip.Approve = "nil"

	h.DB.Save(&ip)
	
	generateNotification(false, ip.Member, ip.Location) //true for "add" and false for "edit"
	
	c.JSON(http.StatusOK, &ip)
}



func (h handler) DeleteIp(c *gin.Context) {
	id := c.Param("id")

	var ip models.IpAddress

	if result := h.DB.First(&ip, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&ip)
	
	c.JSON(http.StatusOK, gin.H{
		"value" : gin.H{
			"code" : http.StatusOK,
			"status": "ip was deleted",
		},
	})
}