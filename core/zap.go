package core

import (
	"fmt"
	"os"

	"github.com/ozeer/go-api/core/internal"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	field := zap.Fields(zap.String("app_name", global.CONFIG.System.AppName))
	logger = zap.New(zapcore.NewTee(cores...), field)

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
