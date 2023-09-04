package core

import (
	"fmt"
	"time"

	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/initialize"
	"github.com/ozeer/go-api/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 如果要启用多点登录拦截或者使用Redis，则初始化Redis
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	// 初始化路由
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	// 初始化Server服务器
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on " + address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
