package main

import (
	"CyberAssetMapper/model"
	_ "embed"
)

func main() {

	model.InitDB()
	defer model.CloseDB()
	//model.Create_taskTable()
	//model.InsertTask("河北大学测绘项目", "hbu.cn")
	model.GetTaskByName()
}
