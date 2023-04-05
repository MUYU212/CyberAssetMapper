package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success",
			})
		})
		rgAuthUser := rgAuth.Group("user")
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
	})
}
