package utils

import (
	"time"

	"github.com/ozeer/go-api/global"
)

// 获取当前日期时间: 2006-01-02 15:04:05
func Date() string {
	return time.Now().Format(global.DATE_TEMPLATE)
}
