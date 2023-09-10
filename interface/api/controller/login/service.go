package login

import (
	"errors"
	"github.com/fighthorse/sampleBookReader/interface/api/service"
	"time"

	"github.com/fighthorse/sampleBookReader/domain/component/gotoken"
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"github.com/fighthorse/sampleBookReader/domain/component/self_errors"
	"github.com/fighthorse/sampleBookReader/interface/api/protos"
	"github.com/gin-gonic/gin"
)

func loginOutEndpoint(c *gin.Context) {
	var token protos.TokenCheck
	if err := c.ShouldBind(&token); err != nil {
		c.JSON(200, gin.H{"code": -126, "message": "token参数无效", "data": map[string]interface{}{}})
		return
	}
	// token 解析 jwt name
	uid, err := gotoken.ParseToken(token.Token, gotoken.LoginSecret)
	if err != nil {
		c.JSON(200, gin.H{"code": -126, "message": "token无效:" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	// uid  缓存数据
	data, ok := service.LoginService.LocalCache.Get(uid)
	// 不存在
	if !ok {
		c.JSON(200, gin.H{"code": -126, "message": "token无效", "data": map[string]interface{}{}})
		return
	}
	service.LoginService.LocalCache.Delete(uid)
	log.Info(c.Request.Context(), "loginOutEndpoint", log.Fields{"data": data})
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": map[string]interface{}{}})
}

func submitEndpoint(c *gin.Context) {
	var person protos.PersonLogin
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}

	if person.Name == "" || person.Pwd == "" {
		err := errors.New("用户名称/密码不能为空")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}
	//verify pwd
	id, err := service.LoginService.VerifyUser(c, person.Name, person.Pwd)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	if id == 0 {
		err := errors.New("服务异常")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
	}
	// login
	token, err := gotoken.CreateToken(person.Name, gotoken.LoginSecret)
	if err != nil {
		c.JSON(200, gin.H{"code": -126, "message": "生产token无效:" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	day := time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	ip := c.ClientIP()
	data := &protos.Person{
		Id:      id,
		Name:    person.Name,
		Ip:      ip,
		Token:   token,
		Expires: day,
	}
	service.LoginService.CacheUser(c, data)

	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": map[string]interface{}{
		"token": token, "exp": day,
	}})
}

func registerEndpoint(c *gin.Context) {
	var person protos.PersonLogin
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(200, self_errors.JsonErrExport(self_errors.JsonErr, err, ""))
		return
	}
	if person.Name == "" || person.Pwd == "" {
		err := errors.New("用户名称/密码不能为空")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}
	if person.Pwd != person.Pwd2 {
		err := errors.New("密码不一致")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}
	if len(person.Name) < 4 {
		err := errors.New("用户名长度不足4位")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}
	if len(person.Pwd) < 6 {
		err := errors.New("密码长度不足6位")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}

	//lock
	if !service.LoginService.LockByName(c, person.Name) {
		err := errors.New("操作速度过快，请稍后重试")
		c.JSON(200, self_errors.JsonErrExport(self_errors.ParamsErr, err, ""))
		return
	}

	//verify pwd
	sok, err := service.LoginService.VerifyUserName(c, person.Name)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	if sok {
		c.JSON(200, gin.H{"code": -1, "message": "用户名已存在", "data": map[string]interface{}{}})
		return
	}

	//创建数据
	err = service.LoginService.Register(c, person.Name, person.Pwd)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "message": "" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	id, err1 := service.LoginService.VerifyUser(c, person.Name, person.Pwd)
	if err1 != nil {
		c.JSON(200, gin.H{"code": -1, "message": "服务异常", "data": map[string]interface{}{}})
		return
	}
	// login
	token, err := gotoken.CreateToken(person.Name, gotoken.LoginSecret)
	if err != nil {
		c.JSON(200, gin.H{"code": -126, "message": "生产token无效:" + err.Error(), "data": map[string]interface{}{}})
		return
	}
	day := time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	ip := c.ClientIP()
	data := protos.Person{
		Id:      id,
		Name:    person.Name,
		Ip:      ip,
		Token:   token,
		Expires: day,
	}
	service.LoginService.CacheUser(c, &data)
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": map[string]interface{}{
		"token": token, "exp": day,
	}})
}
func checkEndpoint(c *gin.Context) {
	var token protos.TokenCheck
	if err := c.ShouldBind(&token); err != nil {
		c.JSON(200, gin.H{"code": -126, "message": "token参数无效", "data": map[string]interface{}{}})
		return
	}
	data, err := service.LoginService.Check(c, token.Token)
	if err != nil {
		c.JSON(200, gin.H{"code": -126, "message": err.Error(), "data": map[string]interface{}{}})
		return
	}
	c.JSON(200, gin.H{"code": 0, "message": "ok", "data": data})
}
