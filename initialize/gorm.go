package initialize

import (
	"os"

	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/model/system"
	"github.com/ozeer/go-api/model/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		// 用户模块表
		user.User{},
		user.UserExtend{},
		system.SysUser{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
