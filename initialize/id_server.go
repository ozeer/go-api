package initialize

import (
	"github.com/ozeer/go-api/global"
	"github.com/sony/sonyflake"
)

func SnowFlakeIdServer() {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
	global.SF = sf
}
