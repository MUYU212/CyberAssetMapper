package main

import (
	"CyberAssetMapper/model"
	"CyberAssetMapper/model/db"
	_ "embed"
)

func main() {

	db.InitDB()
	model.QueryData()

	//初始化数据库之后新建数据库
	//db.InitDB()
	////model.Create_taskTable()
	////model.Insert_Task("www.hbu.cn")
	//model.Create_subDomainTable()
	////model.Update_Task(2, true)
	//r := gin.Default()
	//r.GET("/task/:id", handlers.GetTaskByID)
	//r.POST("/task/create", handlers.CreateTask)
	//r.GET("/task/start/:id", handlers.StartTask)
	//r.Run(":8888")
	//result := SubDomain.GetSubDomain("baidu.com")
	//for _, domain := range result {
	//	log.Println(domain)
	//	//插入数据库的操作
	//	handlers.InsertSubdomain(domain, 1)
	//}
	//handlers.InsertSubdomain()

}
