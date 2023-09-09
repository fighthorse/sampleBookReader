package login

import (
	"errors"
	"github.com/fighthorse/sampleBookReader/domain/component/gotoken"
	"github.com/fighthorse/sampleBookReader/domain/models/model"
	"github.com/fighthorse/sampleBookReader/domain/models/query"
	"github.com/fighthorse/sampleBookReader/interface/api/conf"
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

func (s *Service) VerifyUser(c *gin.Context, userName, pwd string) (ok bool, err error) {
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
		return true, nil
	}
	//系统自定义账户
	cfgList := conf.GConfig.LoginUser
	for _, v := range cfgList {
		if v.UserName == userName {
			if v.UserPwd == pwd {
				return true, nil
			}
			return false, errors.New("账户不存在/密码不正确")
		}
	}
	return false, errors.New("账户不存在/密码不正确")
}

func (s *Service) Check(c *gin.Context, token string) (*protos.Person, error) {
	// token 解析 jwt name
	uid, err := gotoken.ParseToken(token, gotoken.LoginSecret)
	if err != nil {
		return nil, errors.New("token无效:" + err.Error())
	}
	// uid  缓存数据
	data, ok := s.LocalCache.Get(uid)
	// 不存在
	if !ok {
		return nil, errors.New("token无效-未查询到信息")
	}
	// 存在 对比 ip
	dataInfo, _ := data.(protos.Person)
	ip := c.ClientIP()
	if dataInfo.Ip != ip {
		return nil, errors.New("ip发生变化重新登录")
	}
	return &dataInfo, nil
}
