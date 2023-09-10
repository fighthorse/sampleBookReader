package db

//gorm文档: https://gorm.io/zh_CN/docs/index.html
//连接数据库核心代码

import (
	"context"
	"github.com/fighthorse/sampleBookReader/domain/component/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局使用DB,就需要把DB定义成公有的
var StoreInitDB map[string]*StoreDbClient

type StoreDbClient struct {
	Name   string
	Config Mysql
	Master *gorm.DB
	Slave  *gorm.DB
}

// NewDB 自动初始化数据库
func NewDB(arr []Mysql) (err error) {
	StoreInitDB = make(map[string]*StoreDbClient)
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, ip, db)
	for _, v := range arr {
		var mm, vv *gorm.DB
		if v.Master.Driver == "mysql" {
			mm, err = gorm.Open(mysql.Open(v.Master.DSN), &gorm.Config{
				SkipDefaultTransaction: true, //禁用事物
				QueryFields:            true, // 打印sql
				PrepareStmt:            true,
			})
			if v.Slave.DSN != "" {
				vv, err = gorm.Open(mysql.Open(v.Slave.DSN), &gorm.Config{
					SkipDefaultTransaction: true, //禁用事物
					QueryFields:            true, // 打印sql
					PrepareStmt:            true,
				})
			}
			if vv == nil {
				vv = mm
			}
			StoreInitDB[v.Name] = &StoreDbClient{
				Name:   v.Name,
				Config: v,
				Master: mm,
				Slave:  vv,
			}
			if err != nil {
				log.Error(context.Background(), "NewDB error", log.Fields{"config": v, "err": err})
			}
		}

	}
	return
}

func LoadDBByName(dbName string) *StoreDbClient {
	if l, ok := StoreInitDB[dbName]; ok {
		return l
	}
	return nil
}
