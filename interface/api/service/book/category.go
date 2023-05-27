package book

import (
	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) []interface{} {

	redis.BaseReids.Get()
	return nil
}
