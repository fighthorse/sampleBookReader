package protos

type TestHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

type PersonLogin struct {
	Name string `form:"name" json:"name"` //登录用户
	Pwd  string `form:"pwd" json:"pwd"`   // 登录密钥
	Pwd2 string `form:"pwd2" json:"pwd2"` // 登录密钥
}

type Person struct {
	Id      int32  `form:"id" json:"id"`           //登录用户id
	Name    string `form:"name" json:"name"`       //登录用户
	Ip      string `form:"ip" json:"ip"`           //登录用户
	Token   string `form:"token" json:"token"`     // token有效
	Expires string `form:"expires" json:"expires"` // 到期时间
}

type TokenCheck struct {
	Token string `form:"token" json:"token"`
}
