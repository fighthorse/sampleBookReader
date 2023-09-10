package member

import (
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/component/self_errors"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"github.com/gin-gonic/gin"
	"time"
)

func addShelf(c *gin.Context) {
	req := &protos.AddShelfReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	userInfo, ok := c.Get("user-info")
	log.Info(c.Request.Context(), "shelfList", log.Fields{"req": req, "userInfo": userInfo})
	userI, _ := userInfo.(*protos.Person)
	if !ok || userI == nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}
	data := service.MemberService.AddShelf(c, req, userI.Id)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func addReader(c *gin.Context) {
	req := &protos.AddShelfReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	userInfo, ok := c.Get("user-info")
	log.Info(c.Request.Context(), "shelfList", log.Fields{"req": req, "userInfo": userInfo})
	userI, _ := userInfo.(*protos.Person)
	if !ok || userI == nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}
	data := service.MemberService.AddReader(c, req, userI.Id)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func feedback(c *gin.Context) {
	req := &protos.FeedBackReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	userInfo, ok := c.Get("user-info")
	userI, _ := userInfo.(*protos.Person)
	if !ok || userI == nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}
	data := service.MemberService.FeedBack(c, req, userI.Id, userI.Name)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func getfeedback(c *gin.Context) {
	req := &protos.GetFeedBackReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	userInfo, ok := c.Get("user-info")
	userI, _ := userInfo.(*protos.Person)
	if !ok || userI == nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}
	data, err := service.MemberService.GetFeedBack(c, req, userI.Id)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func shelfList(c *gin.Context) {
	req := &protos.MemberReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	userInfo, ok := c.Get("user-info")
	userI, _ := userInfo.(*protos.Person)
	if !ok || userI == nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}
	data := service.MemberService.ShelfList(c, req, userI.Id)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func readerList(c *gin.Context) {
	req := &protos.MemberReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	log.Info(c.Request.Context(), "shelfList", log.Fields{"req": req})

	userInfo, ok := c.Get("user-info")
	log.Info(c.Request.Context(), "shelfList", log.Fields{"req": req, "userInfo": userInfo})
	userI, _ := userInfo.(*protos.Person)
	if !ok || userI == nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}
	data := service.MemberService.ReaderList(c, req, userI.Id)
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
