package main

import (
	"github.com/GoAdminGroup/components/login"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                   // ui theme
	_ "github.com/GoAdminGroup/themes/adminlte/separation"
	_ "github.com/GoAdminGroup/themes/sword"
	_ "github.com/GoAdminGroup/themes/sword/separation"

	// 引入theme2登录页面主题，如不用，可以不导入
	_ "github.com/GoAdminGroup/components/login/theme2"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"github.com/fighthorse/sampleBookReader/goadmin/models"
	"github.com/fighthorse/sampleBookReader/goadmin/pages"
	"github.com/fighthorse/sampleBookReader/goadmin/tables"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)
	// 使用登录页面组件
	login.Init(login.Config{
		Theme:         "theme2",
		CaptchaDigits: 4, // 使用图片验证码，这里代表多少个验证码数字
		// 使用腾讯验证码，需提供appID与appSecret
		// TencentWaterProofWallData: login.TencentWaterProofWallData{
		//    AppID:"",
		//    AppSecret: "",
		// }
	})

	if err := eng.AddConfigFromYAML("./config.yml").
		AddPlugins(adminPlugin).
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}
	// 载入对应验证码驱动，如没使用不用载入
	adminPlugin.SetCaptcha(map[string]string{"driver": login.CaptchaDriverKeyDefault})

	r.Static("/uploads", "./../uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	models.Init(eng.MysqlConnection())

	_ = r.Run(":8090")
	log.Print("Run Server :8090 ,you can find: http://127.0.0.1:8090/admin/login")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
