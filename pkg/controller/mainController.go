package controller

import (
	//"time"
	
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/middleware"	
)

type handler struct {
	DB *gorm.DB
}

type PingReqBody struct {
	IpAddresss	string	`json:"ip_address"`
}

type UserReqBody struct {
	Username       string `json:"username"`
	Password      string `json:"password"`
	Previlage string `json:"previlage"`
}

type GroupReqBody struct{
	GroupSet string `json:"group_set,omitempty"`
	Group string `json:"group,omitempty"`
	IpClass string `json:"ip_class,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	IpGateway string `json:"ip_gateway,omitempty"`
	Data []IpReqBody `json:"data,omitempty"`
}

type IpReqBody struct {
	ClassIp string `json:"ip_class,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	IpAddressed string `json:"ip_addressed,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	DescriptionOne string `json:"description_1,omitempty"`
	DescriptionTwo string `json:"description_2,omitempty"`
	DescriptionThree string `json:"description_3,omitempty"`
	IpGateway string `json:"ip_gateway,omitempty"`
	Location string `json:"location,omitempty"`
	ActivityStatus string `json:"activity_status,omitempty"`
	GroupSet string `json:"group_set,omitempty"`
	Group string `json:"group,omitempty"`
	IpUsageStatus string `json:"ip_usage_status,omitempty"`
	Member string `json:"member,omitempty"`
	Approve string `json:"approve,omitempty"`
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.POST("/login", GenerateToken)

	routes := r.Group("/ip").Use(middleware.Auth())
	
	routes.POST("/register", h.Register)
	routes.GET("/user", h.GetUsers)
	routes.GET("/user/:id", h.GetUser)
	routes.PUT("/user/:id", h.UpdateUser)
	routes.DELETE("/user/:id", h.DeleteUser)

	routes.POST("/address", h.AddIpAddress)
	routes.GET("/address", h.GetIps)
	routes.GET("/address/:id", h.GetIp)
	routes.GET("/address/filter", h.GetSearch)
	routes.PUT("/address/:id", h.UpdateIpByAdmin)
	routes.PUT("/addressed/:id", h.UpdateIpByUser)
	routes.DELETE("/address/:id", h.DeleteIp)

	routes.POST("/group", h.AddGroup)
	routes.GET("/group", h.GetGroup)
	routes.DELETE("/group/:id", h.DeleteGroup)

	routes.GET("/ping", h.PingIpAddress)
	routes.POST("/ping", h.SinglePing)

}
