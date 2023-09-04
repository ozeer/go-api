package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/utils"
)

func OtherInit() {
	// 解析JWT过期时间，默认7d
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}

	// 解析JWT缓冲时间，默认1d
	_, err = utils.ParseDuration(global.CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
