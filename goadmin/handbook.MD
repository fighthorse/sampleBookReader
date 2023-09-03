


使用以下命令，可以打开一个web界面安装：
$ adm init web -l cn


一. 安装步骤
1. 下载并导入对应的SQL文件：
Mysql

2. 依次执行以下命令：

> GO111MODULE=on go mod init github.com/fighthorse/sampleBookReader/goadmin
> GORPOXY=https://goproxy.io GO111MODULE=on go mod tidy
> GO111MODULE=on go run .
> 
3. 登录访问

http://127.0.0.1:10080/admin/login

账号：admin  密码：admin

4. 查看文件中的readme了解更多

二. 关于GoAdmin
Github： https://github.com/GoAdminGroup/go-admin
官网：https://www.go-admin.cn/
论坛：http://discuss.go-admin.com/
文档：http://doc.go-admin.cn/zh/

使用工具快速生成数据模型文件：
1) 使用自带的命令行工具
在项目文件夹中执行：
adm generate -l cn