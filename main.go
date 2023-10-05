package main

import (
	"github.com/ozeer/go-api/core"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化Viper
	global.VP = core.Viper()
	initialize.OtherInit()
	// 初始化zap日志库
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)
	// gorm连接数据库
	global.DB = initialize.Gorm()
	// 初始化定时器
	initialize.Timer()
	// 初始化数据库
	initialize.DBList()
	if global.DB != nil {
		// 初始化表
		initialize.RegisterTables()
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	// 初始化雪花ID生成器
	initialize.SnowFlakeIdServer()

	// 启动Server
	core.RunWindowsServer()
}
