package middleware

import (
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func CustomRecovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(c, "panic", log.Fields{"err": err, "stack": string(debug.Stack())})
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
