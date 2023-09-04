//go:build !windows
// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	// 设置允许读取Header头的超时时间为20s
	s.ReadHeaderTimeout = 20 * time.Second
	// 设置写超时时间为20s
	s.WriteTimeout = 20 * time.Second
	// 控制HTTP请求头的最大大小（以字节为单位）为1048576（1*2^20）
	s.MaxHeaderBytes = 1 << 20
	return s
}
