package main

import (
	_ "embed"
	"fmt"
)

func main() {

	//model.InitDB()
	//defer model.CloseDB()
	////model.Create_taskTable()
	////model.InsertTask("河北大学测绘项目", "hbu.cn")
	//model.GetTaskByName()
	var name string
	name = "hello world xxxx"
	if name != "" {
		fmt.Println(name)
	}
}
