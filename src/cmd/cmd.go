package cmd

import (
	"CyberAssetMapper/src/conf"
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/router"
	"fmt"
)

func Start() {
	//读取相关的配置文件内容
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()
	//初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("=============Clean===============")
}
