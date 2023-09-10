package book

import (
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/component/self_errors"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"github.com/gin-gonic/gin"
	"time"
)

func categoryList(c *gin.Context) {
	var req protos.CategoryReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	data := service.BookService.GetCategory(c, &req)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func bookInfo(c *gin.Context) {
	var req protos.CategoryReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	day := time.Now().Format("2006-01-02 15:04:05")
	log.Info(c.Request.Context(), "categoryList", log.Fields{"req": req})
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": map[string]interface{}{
		"req": req, "day": day,
	}})
}

func bookList(c *gin.Context) {
	var req protos.BookListReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	resp, err := service.BookService.BookList(c, &req)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": resp})
}

func chapter(c *gin.Context) {
	var req protos.CategoryReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	day := time.Now().Format("2006-01-02 15:04:05")
	log.Info(c.Request.Context(), "categoryList", log.Fields{"req": req})
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": map[string]interface{}{
		"req": req, "day": day,
	}})
}
