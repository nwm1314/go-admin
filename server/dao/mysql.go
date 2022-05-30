package dao

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB
var SqlDB *sql.DB

func init() {
	//执行main之前 先执行init方法
	dsn := "root:123456@tcp(127.0.0.1:3306)/gva?charset=utf8mb4&parseTime=true&loc=Local"
	//db, err := sql.Open("mysql", dataSourceName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	SqlDB, err = DB.DB()
	if err != nil {
		log.Println("获取sqldb失败")
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDB.SetConnMaxLifetime(time.Hour)
}
