package main

import (
	"database/sql"
	"fmt"
	"github.com/gohouse/converter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	//SqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//SqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//SqlDB.SetConnMaxLifetime(time.Hour)
}

func TableToStruct(tableName string) {
	defer SqlDB.Close()
	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		// 每个struct放入单独的文件, 默认false, 放入同一个文件(暂未提供),
		SeperatFile: false,
	})
	fmt.Println(tableName)
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(tableName).
		// 表名前缀
		Prefix("").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("main").
		// tag字段的key值,默认是orm
		//TagKey("gorm").
		// 是否添加结构体方法获取表名
		RealNameMethod("TableName").
		// 生成的结构体保存路径
		SavePath("model.go").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		// "用户名:密码@tcp(数据库地址:端口号)/数据库名?charset=utf8"
		DB(SqlDB).
		// 执行
		Run()

	fmt.Println(err)
}

func main() {
	TableToStruct("sys_apis")
}
