package book

import (
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
)

func (s *Service) BookInfo(c *gin.Context, req *protos.ChapterReq, userInfo *protos.Person) (resp *protos.BookInfoResp, err error) {

	resp = &protos.BookInfoResp{}

	u := query.Use(s.Dao.Master).Book
	tt, err := u.WithContext(c.Request.Context()).Where(u.ID.Eq(req.BookId)).First()
	if err != nil {
		return
	}
	resp.Info = tt
	if userInfo != nil && userInfo.Id > 0 {
		resp.Mid = userInfo.Id
		resp.Name = userInfo.Name
		n := query.Use(s.Dao.Master).MemberShelf
		tt2, _ := n.WithContext(c.Request.Context()).Where(n.MemberID.Eq(userInfo.Id), n.BookID.Eq(req.BookId)).First()
		if tt2 != nil && tt2.ID > 0 {
			resp.IsShelf = 1
		}
	}
	return
}

func (s *Service) BookList(c *gin.Context, req *protos.BookListReq) (resp protos.BookListResp, err error) {
	resp = protos.BookListResp{}
	// 数据组装
	if req.Ps <= 0 || req.Ps > 20 {
		req.Ps = 10
	}
	if req.Pn <= 0 || req.Pn > 1000 {
		req.Pn = 1
	}
	// userName lock
	u := query.Use(s.Dao.Master).Book
	tt := u.WithContext(c.Request.Context()).Where(u.State.Eq(5))
	if req.Name != "" {
		srr := "%" + req.Name + "%"
		tt = tt.Where(u.WithContext(c.Request.Context()).Or(u.BookTitle.Like(srr)).Or(u.BookDesc.Like(srr)).Or(u.Copyright.Like(srr)))
	}
	if req.Category > 0 {
		tt = tt.Where(u.CategoryID.Eq(req.Category))
	}
	log.Warn(c, "CountSql", log.Fields{"req": req})

	total, err := tt.Debug().Count()
	if err != nil {
		log.Error(c, "CountErr", log.Fields{"err": err, "req": req})
		return
	}
	resp.Total = total
	if total == 0 {
		return
	}
	list, err1 := tt.Offset((req.Pn - 1) * req.Ps).Limit(req.Ps).Find()
	if err1 != nil && err1.Error() != query.EmptyRecord {
		err = err1
		log.Error(c, "FindErr", log.Fields{"err": err1, "req": req})
		return
	}

	resp.List = list
	resp.Pn = req.Pn
	return
}