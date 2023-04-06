package main

import (
	"CyberAssetMapper/src/cmd"
	_ "embed"
)

// @title 资产测绘系统 API Doc
// @version 0.0.1
// @description 资产测绘系统
func main() {

	//model.InitDB()
	//defer model.CloseDB()
	////model.Create_taskTable()
	////model.InsertTask("河北大学测绘项目", "hbu.cn")
	//model.GetTaskByName()
	defer cmd.Clean()
	cmd.Start()
}
