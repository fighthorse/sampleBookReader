package login

import (
	"encoding/json"
	"errors"
	"github.com/fighthorse/sampleBookReader/domain/component/gotoken"
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/models/model"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/conf"
	"github.com/fighthorse/sampleBookReader/interface/api/dao/redis"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
	"time"
	"xorm.io/xorm/caches"
)

// VerifyUserName 验证用户名是否存在
func (s *Service) VerifyUserName(c *gin.Context, userName string) (ok bool, err error) {
	// userName lock
	u := query.Use(s.Dao.Master).Member
	user, err1 := u.WithContext(c.Request.Context()).Where(u.MemberName.Eq(userName)).First()
	if err1 != nil {
		if err1.Error() == query.EmptyRecord {
			return false, nil
		}
		err = err1
		return
	}
	if user != nil && user.MemberName != "" {
		return true, nil
	}
	return false, nil
}

func (s *Service) LockByName(c *gin.Context, userName string) bool {
	//lock name redis
	lockOk, err0 := s.RedisCache.SetNxKey(c, redis.LockRegisterNameKey.Fmt(userName), "1", 3*time.Second)
	if err0 != nil || !lockOk {
		return false
	}
	return true
}

func (s *Service) Register(c *gin.Context, userName, pwd string) (err error) {
	// userName lock
	u := query.Use(s.Dao.Master).Member
	user, err1 := u.WithContext(c.Request.Context()).Where(u.MemberName.Eq(userName)).First()
	if err1 != nil && err1.Error() != query.EmptyRecord {
		err = err1
		return
	}
	// 验证密码
	if user != nil && user.MemberName != "" {
		err = errors.New("账户已存在")
		return
	}
	err = u.WithContext(c.Request.Context()).Create(&model.Member{
		MemberName:  userName,
		MemberPwd:   caches.Md5(pwd),
		MemberDesc:  "这个人自己注册",
		ReadBooks:   0,
		RegisterDay: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return
	}
	return
}

func (s *Service) VerifyUser(c *gin.Context, userName, pwd string) (id int32, err error) {
	// userName lock
	u := query.Use(s.Dao.Master).Member
	user, err1 := u.WithContext(c.Request.Context()).Where(u.MemberName.Eq(userName)).First()
	if err1 != nil {
		if err1.Error() == query.EmptyRecord {
			err = errors.New("用户不存在或者用户名错误")
			return
		}
		err = err1
		return
	}
	if user == nil {
		err = errors.New("用户不存在或者用户名错误")
		return
	}
	// 验证密码
	if user.MemberPwd == caches.Md5(pwd) {
		return user.ID, nil
	}
	//系统自定义账户
	cfgList := conf.GConfig.LoginUser
	for _, v := range cfgList {
		if v.UserName == userName {
			if v.UserPwd == pwd {
				return 1, nil
			}
			return 0, errors.New("账户不存在/密码不正确")
		}
	}
	return 0, errors.New("账户不存在/密码不正确")
}
func (s *Service) CacheUser(c *gin.Context, data *protos.Person) {
	str, _ := json.Marshal(data)
	_, _ = s.RedisCache.SetKey(c, redis.LoginUidKey.Fmt(data.Name), string(str), 24*time.Hour)

	s.LocalCache.Set(data.Name, data, 10*time.Minute)
}

func (s *Service) Check(c *gin.Context, token string) (*protos.Person, error) {
	res := &protos.Person{}
	// token 解析 jwt name
	uid, err := gotoken.ParseToken(token, gotoken.LoginSecret)
	if err != nil {
		return nil, errors.New("token无效:" + err.Error())
	}
	// uid  缓存数据
	data, ok := s.LocalCache.Get(uid)
	if !ok {
		str, _ := s.RedisCache.GetKey(c, redis.LoginUidKey.Fmt(uid))
		if len(str) > 0 {
			_ = json.Unmarshal([]byte(str), res)
			if res.Token != "" {
				s.LocalCache.Set(uid, res, 10*time.Minute)
				return res, nil
			}
		}
		return nil, errors.New("token无效-未查询到信息")
	}
	// 存在 对比 ip
	ip := c.ClientIP()
	dataInfo, ok := data.(*protos.Person)
	log.Warn(c, "Check", log.Fields{"uid": uid, "data": data, "dataInfo": dataInfo, "ip": ip})
	if !ok {
		return nil, errors.New("token无效-未查询到信息")
	}
	if dataInfo.Ip != ip {
		return nil, errors.New("ip发生变化重新登录")
	}
	return dataInfo, nil
}
