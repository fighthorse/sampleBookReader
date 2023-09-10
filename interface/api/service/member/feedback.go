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

// GetFeedBack
func (s *Service) GetFeedBack(c *gin.Context, req *protos.GetFeedBackReq, uid int32) (resp *protos.GetFeedBackResp, err error) {
	resp = &protos.GetFeedBackResp{
		List:  []interface{}{},
		Total: 0,
	}
	// 数据组装
	if req.Ps <= 0 || req.Ps > 20 {
		req.Ps = 10
	}
	if req.Pn <= 0 || req.Pn > 1000 {
		req.Pn = 1
	}
	u := query.Use(s.Dao.Master).Feedback
	yy := u.WithContext(c.Request.Context()).Debug()
	total, err1 := yy.Count()
	if err1 != nil || total == 0 {
		return
	}
	resp.Pn = req.Pn
	resp.Total = total
	list, err2 := yy.Order(u.FeedDay.Desc()).Offset((req.Pn - 1) * req.Ps).Limit(req.Ps).Find()
	if err2 != nil && err2.Error() != query.EmptyRecord {
		err = err2
		log.Error(c, "FindErr", log.Fields{"err": err1, "req": req})
		return
	}
	resp.List = list
	return
}

// FeedBack
func (s *Service) FeedBack(c *gin.Context, req *protos.FeedBackReq, uid int32, name string) (err error) {

	// 用户uid 枷锁
	ok, err2 := s.RedisCache.SetNxKey(c, redis.FeedUidKey.Fmt(uid), "1", 3*time.Second)
	if err2 != nil {
		err = err2
		return
	}
	if !ok {
		err = errors.New("服务繁忙")
		return
	}
	// 创建记录
	var da = &model.Feedback{
		MemberID: uid,
		Feed:     "【" + name + "】:" + req.Content,
		FeedDay:  time.Now().Format("2006-01-02 15:04:05"),
		Callback: "",
		UserID:   0,
	}
	// 判断是否已经加过
	u := query.Use(s.Dao.Master).Feedback
	err = u.WithContext(c.Request.Context()).Create(da)
	return err
}
