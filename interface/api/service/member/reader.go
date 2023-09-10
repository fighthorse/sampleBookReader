package member

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

func (s *Service) AddShelf(c *gin.Context, req *protos.AddShelfReq, uid int32) (err error) {
	// 判断是否已经加过
	u := query.Use(s.Dao.Master).MemberShelf
	tt, _ := u.WithContext(c.Request.Context()).Where(u.MemberID.Eq(uid), u.BookID.Eq(req.BookId)).First()
	//if err1 != nil && err1.Error() != query.EmptyRecord {
	//	err = err1
	//	return
	//}
	if tt != nil && tt.ID > 0 {
		err = errors.New("记录已存在")
		return
	}
	log.Info(c, "AddShelf1", log.Fields{"req": req, "uid": uid})
	// 用户uid 枷锁
	ok, err2 := s.RedisCache.SetNxKey(c, redis.ShelfUidKey.Fmt(uid), "1", 3*time.Second)
	if err2 != nil {
		err = err2
		return
	}
	if !ok {
		err = errors.New("服务繁忙")
		return
	}
	// 创建记录
	da := &model.MemberShelf{
		BookID:    req.BookId,
		ChapterID: 0,
		ReadDay:   time.Now().Format("2006-01-02 15:04:05"),
		MemberID:  uid,
	}
	err = u.WithContext(c.Request.Context()).Create(da)
	log.Warn(c, "AddShelf", log.Fields{"da": da})
	return err
}

func (s *Service) AddReader(c *gin.Context, req *protos.AddShelfReq, uid int32) (err error) {
	// 判断是否已经加过
	u := query.Use(s.Dao.Master).MemberReader
	tt, err1 := u.WithContext(c.Request.Context()).Where(u.MemberID.Eq(uid), u.BookID.Eq(req.BookId)).First()
	if err1 != nil {
		err = err1
		return
	}
	if tt != nil && tt.ID > 0 {
		err = errors.New("记录已存在")
		return
	}
	// 用户uid 枷锁
	ok, err2 := s.RedisCache.SetNxKey(c, redis.ShelfUidKey.Fmt(uid), "1", 3*time.Second)
	if err2 != nil {
		err = err2
		return
	}
	if !ok {
		err = errors.New("服务繁忙")
		return
	}
	tda := time.Now().Format("2006-01-02 15:04:05")
	// 创建记录
	da := &model.MemberReader{
		MemberID:   uid,
		BookID:     req.BookId,
		ChapterID:  req.ChapterId,
		LastUpdate: tda,
	}
	err = u.WithContext(c.Request.Context()).Create(da)

	// 是否加入书架 更新书籍的章节信息
	u2 := query.Use(s.Dao.Master).MemberShelf
	newdata := map[string]interface{}{
		"chapter_id": req.ChapterId,
		"read_day":   tda,
	}
	_, _ = u2.WithContext(c.Request.Context()).Where(u2.MemberID.Eq(uid), u2.BookID.Eq(req.BookId)).Updates(newdata)
	return err
}

func (s *Service) ReaderList(c *gin.Context, req *protos.MemberReq, uid int32) (resp *protos.MemberBookListResp) {
	// 数据组装
	if req.Ps <= 0 || req.Ps > 20 {
		req.Ps = 10
	}
	if req.Pn <= 0 || req.Pn > 1000 {
		req.Pn = 1
	}
	resp = &protos.MemberBookListResp{
		Pn: req.Pn,
	}
	u := query.Use(s.Dao.Master).MemberReader
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
	list, err1 := tt.Order(u.LastUpdate.Desc()).Offset((req.Pn - 1) * req.Ps).Limit(req.Ps).Find()
	if err1 != nil && err1.Error() != query.EmptyRecord {
		err = err1
		log.Error(c, "FindErr", log.Fields{"err": err1, "req": req})
		return
	}
	var bookIds []int32
	var chapterIds []int32
	for _, kk := range list {
		bookIds = append(bookIds, kk.BookID)
		chapterIds = append(chapterIds, kk.ChapterID)
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
	chapterInfo, _ := zz.WithContext(c.Request.Context()).Debug().Where(zz.ID.In(chapterIds...)).Select(zz.ID, zz.ChapterName, zz.ChapterRank).Find()
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
			"last_update":  kk.LastUpdate,
			"book_info":    bookMap[kk.BookID],
			"chapter_info": chapterMap[kk.ChapterID],
		}
		newList = append(newList, item)
	}
	resp.List = newList
	return
}
