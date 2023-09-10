package book

import (
	"errors"
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/models/model"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/redis"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
	"time"
)

// AddBookChapterComment
func (s *Service) AddBookChapterComment(c *gin.Context, req *protos.ChapterCommentReq, userI *protos.Person) (err error) {
	// lock
	// 用户uid 枷锁
	ok, err2 := s.Cache.SetNxKey(c, redis.CommentUidKey.Fmt(userI.Id), "1", 3*time.Second)
	if err2 != nil {
		err = err2
		return
	}
	if !ok {
		err = errors.New("操作太频繁")
		return
	}
	// add
	data := &model.Comment{
		BookID:      req.BookId,
		ChapterID:   req.ChapterId,
		MemberID:    userI.Id,
		CommentDesc: "【" + userI.Name + "】:" + req.Content,
	}
	u := query.Use(s.Dao.Master).Comment
	err = u.WithContext(c.Request.Context()).Create(data)
	return
}

// GetBookChapterComment
func (s *Service) GetBookChapterComment(c *gin.Context, req *protos.ChapterReq) (resp protos.BookListResp, err error) {
	resp = protos.BookListResp{
		List: []interface{}{},
	}
	// 数据组装
	if req.Ps <= 0 || req.Ps > 50 {
		req.Ps = 30
	}
	if req.Pn <= 0 || req.Pn > 10000 {
		req.Pn = 1
	}
	// userName lock
	u := query.Use(s.Dao.Master).Comment
	tt := u.WithContext(c.Request.Context()).Where(u.BookID.Eq(req.BookId), u.ChapterID.Eq(req.ChapterId))
	total, err := tt.Debug().Count()
	if err != nil {
		log.Error(c, "CountErr", log.Fields{"err": err, "req": req})
		return
	}
	resp.Total = total
	if total == 0 {
		return
	}
	list, err1 := tt.Order(u.ID.Desc()).Offset((req.Pn - 1) * req.Ps).
		Limit(req.Ps).Find()
	if err1 != nil && err1.Error() != query.EmptyRecord {
		err = err1
		log.Error(c, "FindErr", log.Fields{"err": err1, "req": req})
		return
	}

	resp.List = list
	resp.Pn = req.Pn
	return
}
