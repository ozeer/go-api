package core

import (
	"fmt"
	"time"

	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 如果要启用多点登录拦截或者使用Redis，则初始化Redis
	if global.CONFIG.System.UseMultipoint || global.CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 初始化路由
	Router := initialize.Routers()

	// 初始化Server服务器
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success!", zap.String("port", address))
	global.LOG.Error(s.ListenAndServe().Error())
}
