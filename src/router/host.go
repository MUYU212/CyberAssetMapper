package router

import (
	"CyberAssetMapper/src/api"
	"github.com/gin-gonic/gin"
)

func InitHostRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostApi := api.NewHostApi() //返回一个UserApi的结构体，并且该结构体可以直接调用Login方法
		//这里其实应该是定义了一个权限校验的rgAuthUser，然后在这里使用
		rgAuthUser := rgAuth.Group("host")
		{
			rgAuthUser.POST("/shutdown", hostApi.Shutdown)
		}
	})
}
