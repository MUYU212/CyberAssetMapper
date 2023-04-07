package cmd

import (
	"CyberAssetMapper/src/conf"
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/router"
	"CyberAssetMapper/src/utils"
	"fmt"
)

func Start() {
	var initErr error
	//读取相关的配置文件内容
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()

	//初始化数据库
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
		global.Logger.Error("初始化数据库失败")
	}

	//初始化redis
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
		global.Logger.Error("初始化redis失败")
	}
	_ = global.RedisClient.Set("test", "test1111")
	fmt.Println(global.RedisClient.Get("test"))

	//判断初始化过程中，遇到错误的最终处理
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	//初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("=============Clean===============")
}
