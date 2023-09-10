package member

import (
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/models/model"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
)

func (s *Service) ShelfList(c *gin.Context, req *protos.MemberReq, uid int32) (resp *protos.MemberBookListResp) {
	resp = &protos.MemberBookListResp{
		Pn:   req.Pn,
		List: []interface{}{},
	}
	// 数据组装
	if req.Ps <= 0 || req.Ps > 20 {
		req.Ps = 10
	}
	if req.Pn <= 0 || req.Pn > 1000 {
		req.Pn = 1
	}

	u := query.Use(s.Dao.Master).MemberShelf
	tt := u.WithContext(c.Request.Context()).Where(u.MemberID.Eq(uid))
	total, err := tt.Debug().Count()
	if err != nil {
		log.Error(c, "CountErr", log.Fields{"err": err, "req": req})
		return
	}
	resp.Total = total
	if total == 0 {
		return
	}
	list, err1 := tt.Order(u.ReadDay.Desc()).Offset((req.Pn - 1) * req.Ps).Limit(req.Ps).Find()
	if err1 != nil && err1.Error() != query.EmptyRecord {
		err = err1
		log.Error(c, "FindErr", log.Fields{"err": err1, "req": req})
		return
	}
	var bookIds []int32
	var chapterIds []int32
	for _, kk := range list {
		bookIds = append(bookIds, kk.BookID)
		chapterIds = append(bookIds, kk.ChapterID)
	}
	// 批量查询数据信息
	yy := query.Use(s.Dao.Master).Book
	booksInfo, _ := yy.WithContext(c.Request.Context()).Where(yy.ID.In(bookIds...)).Find()
	bookMap := make(map[int32]*model.Book)
	for _, b := range booksInfo {
		bookMap[b.ID] = b
	}

	//章节
	zz := query.Use(s.Dao.Master).Chapter
	chapterInfo, _ := zz.WithContext(c.Request.Context()).Where(zz.ID.In(chapterIds...)).Select(zz.ID, zz.ChapterName, zz.ChapterRank).Find()
	chapterMap := make(map[int32]*model.Chapter)
	for _, b := range chapterInfo {
		chapterMap[b.ID] = b
	}
	var newList []interface{}
	for _, kk := range list {
		item := map[string]interface{}{
			"id":           kk.ID,
			"member_id":    kk.MemberID,
			"book_id":      kk.BookID,
			"chapter_id":   kk.ChapterID,
			"book_info":    bookMap[kk.BookID],
			"chapter_info": chapterMap[kk.ChapterID],
		}
		newList = append(newList, item)
	}
	resp.List = newList
	return
}
