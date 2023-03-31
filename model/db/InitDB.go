package db

import (
	"database/sql"
	"fmt"
)

var GLOBAL_DB *sql.DB

// var GLOBAL_DB *gorm.DB
//
// func InitDB() {
//
//		db, err := gorm.Open(mysql.New(mysql.Config{
//			DSN: "root:Admin@123@tcp(127.0.0.1:3306)/CyberAssetMapper?charset=utf8&parseTime=True&loc=Local", // DSN data source name
//		}), &gorm.Config{
//			SkipDefaultTransaction: false,
//			NamingStrategy: schema.NamingStrategy{
//				TablePrefix:   "t_",
//				SingularTable: true, //使用单数表名
//			},
//			DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束，使用逻辑外键
//		})
//		fmt.Println(db, err)
//		//给全局变量赋值
//		GLOBAL_DB = db
//	}
func InitDB() {
	dsn := "root:Admin@123@tcp(127.0.0.1:3006)/CyberAssetMapper?charset=utf8&parseTime=True&loc=Local"
	// 打开数据库连接
	db1, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	GLOBAL_DB = db1
	defer db1.Close()
}
