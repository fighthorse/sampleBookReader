package book

import (
	"github.com/fighthorse/sampleBookReader/interface/api/dao/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHttp(r *gin.Engine) {
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	bg := r.Group("/book")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	bg.Use(middleware.GuestRequired)
	{
		bg.GET("/list", bookList)
		bg.GET("/category", categoryList)
		bg.GET("/info", bookInfo)
		bg.GET("/chapter", chapter)
		bg.GET("/chapter_info", chapterInfo)
		bg.GET("/chapter_next", chapterNext)
		bg.GET("/comment", commentList)
		bg.POST("/add_comment", commentAdd)
	}
}
