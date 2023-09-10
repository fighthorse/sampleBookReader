package book

import (
	"github.com/fighthorse/sampleBookReader/domain/component/self_errors"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"github.com/gin-gonic/gin"
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
	var req protos.ChapterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	var userI *protos.Person
	userInfo, ok := c.Get("user-info")
	if ok {
		userI, _ = userInfo.(*protos.Person)
	}

	data, err := service.BookService.BookInfo(c, &req, userI)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
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
	var req protos.ChapterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	data, err := service.BookService.GetBookChapter(c, &req)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

// commentList
func commentList(c *gin.Context) {
	var req protos.ChapterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	data, err := service.BookService.GetBookChapterComment(c, &req)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}
func commentAdd(c *gin.Context) {
	var req protos.ChapterCommentReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	var userI *protos.Person
	userInfo, ok := c.Get("user-info")
	if ok {
		userI, _ = userInfo.(*protos.Person)
	}

	if userI == nil || userI.Id <= 0 {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LoginErr, nil, ""))
		return
	}

	err := service.BookService.AddBookChapterComment(c, &req, userI)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": nil})
}
func chapterInfo(c *gin.Context) {
	var req protos.ChapterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	var userI *protos.Person
	userInfo, ok := c.Get("user-info")
	if ok {
		userI, _ = userInfo.(*protos.Person)
	}

	data, err := service.BookService.GetBookChapterInfo(c, &req, userI)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}

func chapterNext(c *gin.Context) {
	var req protos.ChapterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	data, err := service.BookService.GetBookChapterNext(c, &req)
	if err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.LogicErr, err, ""))
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}
