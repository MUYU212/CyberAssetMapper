package main

import (
	"CyberAssetMapper/handlers"
	"CyberAssetMapper/model/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库之后新建数据库
	db.InitDB()
	//model.Create_taskTable()
	//model.Insert_Task("www.hbu.cn")

	//model.Update_Task(2, true)
	r := gin.Default()
	r.GET("/task/:id", handlers.GetTaskByID)
	r.POST("/task/create", handlers.CreateTask)
	r.Run(":8888")
}
