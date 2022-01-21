package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	Err "pkg/error"
)

var db *gorm.DB

const (
	DefaultUsername = "root"
	DefaultPassword = "123456"
	DefaultHost = "127.0.0.1"
	DefaultPort = 3306
	DefaultDbName = "uber"
)

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DefaultUsername, DefaultPassword, DefaultHost, DefaultPort, DefaultDbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect to mysql failed, err:" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(Err.New(Err.MysqlConnectFail, err))
	}
	sqlDB.SetMaxOpenConns(200)	//设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(200)	//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
}

func GetDB() *gorm.DB {
	return db
}