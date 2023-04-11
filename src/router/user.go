package router

import (
	"CyberAssetMapper/src/api"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi() //返回一个UserApi的结构体，并且该结构体可以直接调用Login方法
		rgPublicUser := rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {
				//ctx.AbortWithStatusJSON(200, gin.H{
				//	"msg": "这里是一个中间件",
				//})
			}
		}())
		{
			rgPublicUser.POST("/login", userApi.Login) //这里是将userApi.Login注册到了gfnRoutes中
		}
		//这里其实应该是定义了一个权限校验的rgAuthUser，然后在这里使用
		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("/add", userApi.AddUser)
			rgAuthUser.POST("/list", userApi.GetUserList)
		}
	})
}
