# sampleBookReader
## 基于golang语言的小说阅读系统设计与实现
简单的书籍阅读系统 以 gin + h5 + bootstrap 框架搭建
[GITHUB](https://github.com/fighthorse/sampleBookReader)

## 环境依赖
+ golang 1.91.0 版本以上
+ mysql 5.0 
+ redis 4.0
+ 浏览器 chrome 版本 116.0.5845.180

## 系统初始化sql
+ 1.先在数据库建一个 go_admin 库
+ 2.执行命令导入初始化的sql (具体文件 sql/go_admin.sql)
> mysql -u user -p  go_admin < ./go_admin.sql

## handle book
+ 1.将本项目clone到 gopath 中 
+ 2.cd $GOPATH:/src/github.com/fighthorse/sampleBookReader

### 启动前台系统 
+ 1.cd interface/cmd
+ 2.go run main.go
+ 3.在浏览器中访问： http://127.0.0.1:8080/html

>[GIN-debug] Listening and serving HTTP on :8080
>[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
>[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
> using env:   export GIN_MODE=release
> using code:  gin.SetMode(gin.ReleaseMode)
>[GIN-debug] GET    /ping                     --> main.InnerServer.func1 (3 handlers)
>[GIN-debug] GET    /                         --> main.InnerServer.func2 (3 handlers)
>[GIN-debug] GET    /metrics                  --> main.InnerServer.func3 (3 handlers)
>[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
>Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
>[GIN-debug] Listening and serving HTTP on :8081
>[GIN-debug] redirecting request 301: /html/ --> /html/

+ 4.前台登录页面：http://127.0.0.1:8080/html/sign-in.html
+ 5.输入账目密码： test3/1234567
+ 6.其他可操作性账户：admin/admin ; test1/test1

### 启动后台系统
+ 1.cd goadmin/
+ 2.go run main.go
`
> GoAdmin 启动成功。
>目前处于 "debug" 模式。请在生产环境中切换为 "release" 模式。
>2023-09-12T09:32:39.241+0800    INFO    engine/engine.go:402    =====> 初始化数据库连接
>2023-09-12T09:32:39.811+0800    INFO    engine/engine.go:402    =====> 初始化配置
>2023-09-12T09:32:39.857+0800    INFO    engine/engine.go:402    =====> 初始化导航栏按钮
>2023-09-12T09:32:39.871+0800    INFO    engine/engine.go:402    =====> 初始化插件
>2023-09-12T09:32:39.879+0800    INFO    engine/engine.go:402    =====> 初始化成功🍺🍺
`

+ 3.在浏览器中访问： http://127.0.0.1:8090/admin/login
+ 4.输入账户: admin/admin 进行登录
+ 5.小说内容管理- 小说信息上传  http://127.0.0.1:8090/admin/info/book
+ 6.小说内容管理- 章节内容上传  http://127.0.0.1:8090/admin/info/chapter
+ 7.小说内容管理- 章节评论管理  http://127.0.0.1:8090/admin/info/comment

## interface 前台系统
+ bootstrapv3.0
+ Layer弹窗
+ Jqueryv3.5.7
+ HTML5
+ gin框架
+ gorm/gen model生产
+ redis
+ go-cache
+ log
+ Mysql

## goadmin 后台管理系统
+ go-admin
+ gorm
+ gin框架
+ mysql

## 感谢
感谢在中国互联网及开源奋斗的前辈，Jquery、Layer 前端组件目前已经停止维护，但其简单的操作模式仍然实用，对于个人开发而言，简单有效是最佳的。
后端框架gin、go-admin、go-cache、redis、mysql，gorm,gorm/gen等组件包是本系统快速搭建的助力，感谢前辈同志的付出。

## 免责声明：

本项目属于论文系统，其设计的目的是实践计算机的理论知识，所以设计上有一些欠缺，同时本系统未对小说内容、评论进行安全审核，所以请不要上传违法的信息

本 Git 仓库所包含的代码仅供参考和学习使用。作者尽力确保代码的质量和准确性，但不对代码的完整性、时效性、适用性或可靠性提供任何保证。

本 Git 仓库的代码可能包含错误或存在缺陷。使用者应仔细检查和验证代码，并自行承担使用该代码所产生的风险。

作者对因使用本 Git 仓库的代码而造成的任何损失或损害不承担任何责任，包括但不限于直接损失、间接损失、附带损失或因此产生的任何其他损失。

本 Git 仓库的代码可能受到知识产权法律的保护。使用者在使用该代码时，请遵守相关的法律和法规，并承担因此产生的任何法律责任。

本 Git 仓库可能包含其他第三方贡献的代码或链接到其他网站的资源。对于这些第三方代码和资源，作者不拥有或控制，并且不对其内容的准确性、安全性或合法性负责。

使用者在使用本 Git 仓库中的代码时，即表示已阅读、理解并同意上述免责声明中的条款和条件。如有任何疑问，请随时联系作者。

## 可视化页面参考
![首页图片](./uploads/homepage.png)

![后台图片](./uploads/img.png)
