package member

import (
	"github.com/fighthorse/sampleBookReader/interface/api/dao/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHttp(r *gin.Engine) {
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	mb := r.Group("/member")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	mb.Use(middleware.GuestRequired)
	{
		mb.GET("/shelf", shelfList)
		mb.GET("/reader", readerList)
		mb.POST("/add_shelf", addShelf)
		mb.POST("/add_reader", addReader)
		mb.POST("/feedback", feedback)
		mb.GET("/feedback", getfeedback)
	}
}
