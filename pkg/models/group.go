package models

import "gorm.io/gorm"

type Groups struct{
	gorm.Model
	GroupSet string `json:"group_set,omitempty"`
	Group string `json:"group,omitempty"`
	IpClass string `"json:"ip_class,omitempty"`
	Netmask string `"json:"netmask,omitempty"`
	IpGateway string `"json:"ip_gateway,omitempty"`
	Data []IpAddress `json:"data,omitempty"`
}

type IpAddress struct{
	gorm.Model
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