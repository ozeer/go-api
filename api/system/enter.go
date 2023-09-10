package system

import "github.com/ozeer/go-api/service"

type ApiGroup struct {
	JwtApi
	BaseApi
}

var (
	snowflakeService = service.ServiceGroupApp.SystemServiceGroup.SnowFlakeService
	jwtService       = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
)
