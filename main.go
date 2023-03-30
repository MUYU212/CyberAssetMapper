package main

import (
	"CyberAssetMapper/Plugins/SubDomain"
	"log"
)

func main() {
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
	result := SubDomain.GetSubDomain("hbu.cn")
	log.Println(result)
}
