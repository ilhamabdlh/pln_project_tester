package controller

import (
	"net/http"
	"fmt"
	"strings"
	"os/exec"
	//"time"
	// /"sync"

	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/models"
)

//var wg sync.WaitGroup
const (
	rto = "Request timed out."
	dhu = "Destination host unreachable."
	online = "online"
	offline = "offline"
)

func Pinger(s string, h handler){

	var ip models.IpAddress
	if result := h.DB.Where("ip_addressed = ?", s).First(&ip); result.Error != nil{
		return
	}

	v, _:= exec.Command("ping", s, "-n", "4").Output()
	if strings.Contains(string(v), dhu){
		rateDhu := strings.Count(string(v), dhu)
		if rateDhu < 3{
			fmt.Printf("%v  is alive\n", s)
			ip.ActivityStatus = online
			h.DB.Save(&ip)
		} else{
			fmt.Printf("%v is unreachable\n", s)
			ip.ActivityStatus = offline
			h.DB.Save(&ip)
		}
	} else if strings.Contains(string(v), rto){
		rateRto := strings.Count(string(v), rto)
		if rateRto < 3{
			fmt.Printf("%v is alive\n", s)
			ip.ActivityStatus = online
			h.DB.Save(&ip)
			
		} else{
			fmt.Printf("%v is request timed out.\n", s)
			ip.ActivityStatus = offline
			h.DB.Save(&ip)
		}
	} else{
		
		fmt.Printf("%v is alive\n", s)
		ip.ActivityStatus = online
		h.DB.Save(&ip)
	}
	
}
func (h handler) PingIpAddress(c *gin.Context){
	var ips []models.IpAddress

	if result := h.DB.Find(&ips); result.Error != nil{
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var s []string
	var first, second []string

	for _, v := range ips{
		s = append(s, v.IpAddressed)
	}
	// wg.Add(1)
	// go Pinger(s, h)
	// wg.Wait()
	var temp float64
	if len(s)%2 !=0{
		temp += float64(len(s))/2
		temp += float64(1/2)
	} else{
		temp += float64(len(s))/2
	}
	for i:= 0; i < len(s); i++{
		if i < int(temp){
			first = append(first, s[i])
		} else {
			second= append(second, s[i])
		}
	}



	for i:= 0; i< len(s); i++{
		//wg.Add(1)
		go Pinger(s[i], h)
		//wg.Wait()
	}
	// wg.Add(2)
	// go Pinger(first, h)
	// go Pinger(second, h)
	// wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"value" : gin.H{
			"code" : http.StatusOK,
			"status": "Ping finished",
		},
	})
}

func (h handler) SinglePing(c *gin.Context) {
	body := PingReqBody{}
	

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("ip", body)
	v, _:= exec.Command("ping", body.IpAddresss, "-n", "4").Output()
	if strings.Contains(string(v), dhu){
		rateDhu := strings.Count(string(v), dhu)
		if rateDhu < 3{
			fmt.Printf("%v  is alive\n", body.IpAddresss)
			c.JSON(http.StatusCreated, gin.H{
				"status" : gin.H{
					"code" : http.StatusOK,
					"ip_address": body.IpAddresss,
					"status_ip": online,
					"message": "ping success",
				},
			})
		} else{
			fmt.Printf("%v is unreachable\n", body.IpAddresss)
			c.JSON(http.StatusCreated, gin.H{
				"status" : gin.H{
					"code" : http.StatusOK,
					"ip_address": body.IpAddresss,
					"status_ip": offline,
					"message": "is unreachable",
				},
			})
		}
	} else if strings.Contains(string(v), rto){
		rateRto := strings.Count(string(v), rto)
		if rateRto < 3{
			fmt.Printf("%v is alive\n", body.IpAddresss)
			c.JSON(http.StatusCreated, gin.H{
				"status" : gin.H{
					"code" : http.StatusOK,
					"ip_address": body.IpAddresss,
					"status_ip": online,
					"message": "ping success",
				},
			})
			
		} else{
			fmt.Printf("%v is request timed out.\n", body.IpAddresss)
			c.JSON(http.StatusCreated, gin.H{
				"status" : gin.H{
					"code" : http.StatusOK,
					"ip_address": body.IpAddresss,
					"status_ip": offline,
					"message": "is request timed out",
				},
			})
		}
	} else{
		
		fmt.Printf("%v is alive\n", body.IpAddresss)
		c.JSON(http.StatusCreated, gin.H{
			"status" : gin.H{
				"code" : http.StatusOK,
				"ip_address": body.IpAddresss,
				"status_ip": online,
				"message": "ping success",
			},
		})
	}

	// if result := h.DB.Create(&user); result.Error != nil {
	// 	c.AbortWithError(http.StatusNotFound, result.Error)
	// 	return
	// }
}