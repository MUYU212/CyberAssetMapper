package main

import (
	"CyberAssetMapper/model"
	"CyberAssetMapper/model/db"
)

func main() {
	//初始化数据库之后新建数据库
	db.InitDB()
	model.Create_taskTable()
	model.Insert_Task("www.hbu.cn")
}
