package middleware

import (
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {

}

type TokenReq struct {
	Token string `form:"token" json:"token"` // token有效
}

func GuestRequired(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
	}
	if token == "" {
		return
	}
	// 解析token
	data, err := service.LoginService.Check(c, token)
	if err != nil {
		return
	}
	c.Set("user-info", data)
	c.Next()
}

func TokenRequired(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
	}

	if token == "" {
		c.JSON(200, gin.H{"code": -126, "message": "需要登录", "data": map[string]interface{}{}})
		c.Abort()
		return
	}
	// 解析token
	data, err := service.LoginService.Check(c, token)
	if err != nil {
		c.JSON(200, gin.H{"code": -126, "message": err.Error(), "data": map[string]interface{}{
			"token": token,
		}})
		c.Abort()
		return
	}
	c.Set("user_info", data)
	c.Next()
}
