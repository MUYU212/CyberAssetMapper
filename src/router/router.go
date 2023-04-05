package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) //自定一个函数类型来实现可插拔的作用

var (
	gfnRoutes []IFnRegistRoute
)

// gfnRoutes 是全局注册方法，存放各个模块向平台注册过来的各个模块的初始化函数
func InitRouter() {
	//新建一个可以被终止的上下文全局环境
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()
	rgPublic := r.Group("/api/v1/public") //公开接口
	rgAuth := r.Group("/api/v1/")         //需要鉴权的接口

	InitBasePlatformRoutes()
	for _, fnRegistRoute := range gfnRoutes {
		//这里是循环gfnRoutes方法，得到其中的值，然后这里的值中其实是需要注册进去的方法，需要的参数也是这两个参数，真的巧妙
		fnRegistRoute(rgPublic, rgAuth)
	}

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8090"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		fmt.Println(fmt.Sprintf("Start Server Listen:%s\n", stPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//TODO: 这里需要记录日志
			panic(fmt.Sprintf("Start Server Error: %s\n", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		//todo: 这里需要记录日志
		panic(fmt.Sprintf("Shutdown Server Error: %s\n", err.Error()))
		return
	}
	fmt.Println("Shutdown Server Success")
}

func RegistRoute(fn IFnRegistRoute) {
	//注册路由方法，将所有要注册的方法给加载到gfnRoutes中
	if fn == nil {
		panic("fn is nil")
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
