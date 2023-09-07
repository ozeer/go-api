package user

import "github.com/ozeer/go-api/service"

type ApiGroup struct {
	UserApi
}

var (
	userService     = service.ServiceGroupApp.UserServiceGroup.UserService
	analysisService = service.ServiceGroupApp.UserServiceGroup.AnalysisService
)
