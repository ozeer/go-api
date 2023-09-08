package initialize

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/router"
)

// 初始化总路由
func Routers() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()
	pprof.Register(Router)

	userRouter := router.RouterGroupApp.User
	systemRouter := router.RouterGroupApp.System

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.LOG.Info("use middleware cors")

	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		// 开放API注册处

	}
	PrivateGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth())
	{
		// 用户相关路由
		userRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
	}

	global.LOG.Info("router register success")
	return Router
}
