package book

import (
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetCategory(c *gin.Context, req *protos.CategoryReq) interface{} {
	u := query.Use(s.Dao.Master).Category
	tt := u.WithContext(c.Request.Context()).Where(u.IsDel.Is(false))
	if req.CategoryId > 0 {
		tt = tt.Where(u.WithContext(c.Request.Context()).Or(u.ParentID.Eq(req.CategoryId)), u.WithContext(c.Request.Context()).Or(u.ID.Eq(req.CategoryId)))
	}

	res, err := tt.Debug().Find()
	log.Warn(c, "FindtSql", log.Fields{"req": req, "err": err})
	if err != nil {
		return nil
	}
	return res
}
