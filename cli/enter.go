package main

import (
	"log"
	"os"

	"github.com/ozeer/go-api/cli/user"
	"github.com/ozeer/go-api/core"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/initialize"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

type Config struct {
	Commands []*cli.Command
}

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
	// 初始化redis服务
	initialize.Redis()

	// 创建配置并添加命令
	config := &Config{
		Commands: []*cli.Command{
			// 在这里添加其他命令，根据需要
			user.CliAnalysisAgeGroup,
			user.CliDemo,
			user.Test,
		},
	}

	app := &cli.App{
		Commands: config.Commands,
		Usage:    "Go-Api的命令行工具",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
