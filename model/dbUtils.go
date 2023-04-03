package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入相应的数据库驱动包，这里以MySQL为例
)

var (
	db  *sql.DB
	err error // 声明 err 变量
)

func InitDB() {
	dsn := "root:Admin@123@tcp(127.0.0.1:3306)/CyberAssetMapper"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接测试失败:", err)
		return
	}

	fmt.Println("数据库连接成功")
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
