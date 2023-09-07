package system

import "github.com/ozeer/go-api/service"

type ApiGroup struct {
	SnowFlakeApi
}

var (
	snowflakeService = service.ServiceGroupApp.SystemServiceGroup.SnowFlakeService
)
