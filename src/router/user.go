package router

import (
	"CyberAssetMapper/src/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi() //返回一个UserApi的结构体，并且该结构体可以直接调用Login方法
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login) //这里是将userApi.Login注册到了gfnRoutes中
		}
		//这里其实应该是定义了一个权限校验的rgAuthUser，然后在这里使用
		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zhangsan"},
						{"id": 2, "name": "lisi"},
						{"id": 3, "name": "wangwu"},
					},
				})
			})
			rgAuthUser.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "zhangsan",
				})
			})
		}
	})
}
