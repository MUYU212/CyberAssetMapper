package router

import (
	_ "CyberAssetMapper/docs"
	"CyberAssetMapper/src/global"
	"CyberAssetMapper/src/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"strings"
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
	r.Use(middleware.Cors())
	rgPublic := r.Group("/api/v1/public") //公开接口
	rgAuth := r.Group("/api/v1/")         //需要鉴权的接口

	//注册平台基础路由
	InitBasePlatformRoutes()

	//注册自定义验证器
	registCustValidator()

	for _, fnRegistRoute := range gfnRoutes {
		//这里是循环gfnRoutes方法，得到其中的值，然后这里的值中其实是需要注册进去的方法，需要的参数也是这两个参数，真的巧妙
		fnRegistRoute(rgPublic, rgAuth)
	}

	//集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8090"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		global.Logger.Info(fmt.Sprintf("Start Server Listen:%s", stPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//TODO: 这里需要记录日志
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s\n", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Shutdown Server Error: %s\n", err.Error()))
		return
	}
	global.Logger.Info("Shutdown Server Success")
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
	InitHostRoutes()
	InitTaskRoutes()
}

func registCustValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && strings.Index(value, "a") == 0 {
					return true
				}
			}
			return false
		})
	}
}
