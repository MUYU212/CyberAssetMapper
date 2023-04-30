package router

import (
	"CyberAssetMapper/src/api"
	"github.com/gin-gonic/gin"
)

func InitTaskRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		taskApi := api.NewTaskApi()
		rgPublicTask := rgPublic.Group("task").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {

			}
		}())
		{
			rgPublicTask.POST("/add", taskApi.AddTask)
		}
	})
}
