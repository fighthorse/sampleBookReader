package login

import (
	"errors"

	"github.com/fighthorse/sampleBookReader/domain/component/gocache"
	"github.com/fighthorse/sampleBookReader/domain/component/gotoken"
	"github.com/fighthorse/sampleBookReader/interface/api/conf"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
)

func VerifyUser(c *gin.Context, userName, pwd string) (bool, error) {
	cfgList := conf.GConfig.LoginUser
	for _, v := range cfgList {
		if v.UserName == userName {
			if v.UserPwd == pwd {
				return true, nil
			}
			return false, errors.New("密码不正确")
		}
	}
	return false, errors.New("账户不存在")
}

func Check(c *gin.Context, token string) (*protos.Person, error) {
	// token 解析 jwt name
	uid, err := gotoken.ParseToken(token, gotoken.LoginSecret)
	if err != nil {
		return nil, errors.New("token无效:" + err.Error())
	}
	// uid  缓存数据
	data, ok := gocache.Get(uid)
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
