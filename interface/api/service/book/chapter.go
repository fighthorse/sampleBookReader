package book

import (
	"context"
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/models/model"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
	"time"
)

// GetBookChapterNext
func (s *Service) GetBookChapterNext(c *gin.Context, req *protos.ChapterReq) (resp map[string]interface{}, err error) {
	resp = make(map[string]interface{})
	resp["chapter_id"] = 0
	//chapter
	u := query.Use(s.Dao.Master).Chapter
	chapter, _ := u.WithContext(c.Request.Context()).Where(u.BookID.Eq(req.BookId), u.ID.Eq(req.ChapterId)).First()

	rank := chapter.ChapterRank
	if req.Pn == 0 {
		chapter1, _ := u.WithContext(c.Request.Context()).Debug().Where(u.BookID.Eq(req.BookId), u.ChapterRank.Lt(rank)).Order(u.ChapterRank.Desc()).First()
		if chapter1 != nil {
			resp["chapter_id"] = chapter1.ID
		}
	} else {
		chapter1, _ := u.WithContext(c.Request.Context()).Debug().Where(u.BookID.Eq(req.BookId), u.ChapterRank.Gt(rank)).Order(u.ChapterRank).First()
		if chapter1 != nil {
			resp["chapter_id"] = chapter1.ID
		}
	}
	return
}

// GetBookChapterInfo
func (s *Service) GetBookChapterInfo(c *gin.Context, req *protos.ChapterReq, userI *protos.Person) (resp protos.BookChapterResp, err error) {
	resp = protos.BookChapterResp{}

	// book
	b := query.Use(s.Dao.Master).Book
	book, _ := b.WithContext(c.Request.Context()).Where(b.ID.Eq(req.BookId)).First()

	//chapter
	u := query.Use(s.Dao.Master).Chapter
	chapter, _ := u.WithContext(c.Request.Context()).Where(u.BookID.Eq(req.BookId), u.ID.Eq(req.ChapterId)).First()

	if userI != nil && userI.Id > 0 {
		resp.Mid = userI.Id
		resp.Name = userI.Name
		// add reader
		go s.AddReaderByBook(context.Background(), req.BookId, req.ChapterId, userI.Id)
		// change shelf
		go s.AddshelfByBook(context.Background(), req.BookId, req.ChapterId, userI.Id)
	}
	resp.Info = book
	resp.Chapter = chapter
	return
}

func (s *Service) AddReaderByBook(c context.Context, book_id, chapter_id, mid int32) {
	u := query.Use(s.Dao.Master).MemberReader
	reader, _ := u.WithContext(c).Where(u.BookID.Eq(book_id), u.MemberID.Eq(mid)).First()
	if reader != nil && reader.ID > 0 {
		data := map[string]interface{}{
			"chapter_id":  chapter_id,
			"last_update": time.Now().Format("2006-01-02 15:04:05"),
		}
		_, _ = u.WithContext(c).Where(u.ID.Eq(reader.ID), u.MemberID.Eq(mid)).Updates(data)
	} else {
		_ = u.WithContext(c).Create(&model.MemberReader{
			MemberID:   mid,
			BookID:     book_id,
			ChapterID:  chapter_id,
			LastUpdate: time.Now().Format("2006-01-02 15:04:05"),
		})
	}
}

func (s *Service) AddshelfByBook(c context.Context, book_id, chapter_id, mid int32) {
	u := query.Use(s.Dao.Master).MemberShelf
	reader, _ := u.WithContext(c).Where(u.BookID.Eq(book_id), u.MemberID.Eq(mid)).First()
	if reader != nil && reader.ID > 0 {
		data := map[string]interface{}{
			"chapter_id": chapter_id,
			"read_day":   time.Now().Format("2006-01-02 15:04:05"),
		}
		_, _ = u.WithContext(c).Where(u.ID.Eq(reader.ID), u.MemberID.Eq(mid)).Updates(data)
	}
}

func (s *Service) GetBookChapter(c *gin.Context, req *protos.ChapterReq) (resp protos.BookListResp, err error) {
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
	u := query.Use(s.Dao.Master).Chapter
	tt := u.WithContext(c.Request.Context()).Where(u.BookID.Eq(req.BookId))
	total, err := tt.Debug().Count()
	if err != nil {
		log.Error(c, "CountErr", log.Fields{"err": err, "req": req})
		return
	}
	resp.Total = total
	if total == 0 {
		return
	}
	list, err1 := tt.Order(u.ChapterRank.Desc()).Offset((req.Pn-1)*req.Ps).
		Limit(req.Ps).Select(u.ID, u.BookID, u.ChapterName, u.ChapterRank).Find()
	if err1 != nil && err1.Error() != query.EmptyRecord {
		err = err1
		log.Error(c, "FindErr", log.Fields{"err": err1, "req": req})
		return
	}

	resp.List = list
	resp.Pn = req.Pn
	return
}
