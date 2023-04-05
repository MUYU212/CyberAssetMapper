package cmd

import (
	"CyberAssetMapper/src/conf"
	"CyberAssetMapper/src/router"
	"fmt"
)

func Start() {
	//读取相关的配置文件内容
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("=============Clean===============")
}
