package book

import (
	"encoding/json"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetCategory(c *gin.Context, req *protos.CategoryReq) interface{} {
	var out []interface{}
	data, err := s.Cache.GetKey(c.Request.Context(), "category")
	if err == nil {
		err = json.Unmarshal([]byte(data), &out)
		if err == nil {
			return out
		}
	}
	return nil
}
