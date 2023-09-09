
要在Windows上启动MySQL服务，你需要按照以下步骤进行操作：

下载MySQL安装程序：前往MySQL官方网站（https://dev.mysql.com/downloads/installer/）下载适用于Windows的MySQL安装程序。选择合适的版本（32位或64位）和安装程序类型（Web或Community）。

运行安装程序：双击下载的安装程序以运行它。根据安装程序的指示进行安装，选择适合你的需求的选项。默认情况下，MySQL服务器和其他组件将会被安装。

配置MySQL服务器：在安装过程中，你需要设置MySQL服务器的配置选项，如端口号、安全设置和管理员密码。确保选择一个安全的密码，并将其牢记。

启动MySQL服务：安装完成后，你可以在Windows服务中找到MySQL服务。按下Win + R组合键打开运行对话框，输入"services.msc"并按下回车键，将打开服务管理窗口。在窗口中，找到名为"MySQL"或"MySQLxx"（xx是MySQL版本号）的服务。

启动MySQL服务：在服务管理窗口中，找到"MySQL"或"MySQLxx"服务，右键点击它，然后选择"启动"选项。MySQL服务将会启动。

现在，MySQL服务已经成功启动了。你可以使用MySQL客户端工具连接到MySQL服务器，执行数据库操作和查询。

D:\linux\windows\mysql8.0\bin\mysqld.exe" --defaults-file="C:\ProgramData\MySQL\MySQL Server 8.1\my.ini" MySQL81


1.先在数据库建一个 go_admin 库
2.执行命令导入初始化的sql
mysql -u user -p  go_admin < ./admin.sql