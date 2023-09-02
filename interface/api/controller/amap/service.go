package amap

import (
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"github.com/gin-gonic/gin"
)

func Weather(c *gin.Context) {
	ip, _ := c.GetQuery("location_ip")
	data := service.AmapService.GetIpWeather(c, ip)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}
