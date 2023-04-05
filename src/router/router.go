package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) //自定一个函数类型来实现可插拔的作用

var (
	gfnRoutes []IFnRegistRoute
)

// gfnRoutes 是全局注册方法，存放各个模块向平台注册过来的各个模块的初始化函数
func InitRouter() {
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public") //公开接口
	rgAuth := r.Group("/api/v1/")         //需要鉴权的接口

	InitBasePlatformRoutes()
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8090"
	}
	err := r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("Run server error: %s", err.Error()))
	}
}

func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		panic("fn is nil")
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
