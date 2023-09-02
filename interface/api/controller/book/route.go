package book

import (
	"github.com/fighthorse/sampleBookReader/interface/api/dao/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHttp(r *gin.Engine) {
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	guest := r.Group("/book")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	guest.Use(middleware.GuestRequired)
	{
		guest.GET("/category", categoryList)
		guest.GET("/info", bookInfo)
		guest.GET("/list", bookList)
		guest.GET("/chapter", chapter)
	}
}
